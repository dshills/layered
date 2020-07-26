package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/editor"
	"github.com/dshills/layered/logger"
	"github.com/dshills/layered/palette"
	"github.com/dshills/layered/syntax"
	"github.com/gdamore/tcell"
)

// Window is an editor window
type Window struct {
	region                image.Rectangle
	screen                tcell.Screen
	hidden                bool
	style                 tcell.Style
	offset                int
	id                    string
	bnum                  int
	hideNumber            bool
	content               []string
	syntax                []syntax.Resulter
	pal                   *palette.Palette
	noGutter              bool
	noNumber              bool
	tabsize               int
	preStyled             [][]tcell.Style
	cursorLine, cursorCol int
	numLines              int
	reqC                  chan editor.Request
}

// SetResponse will interprit an editor response
func (w *Window) SetResponse(resp editor.Response) {
	if w.id == "" && resp.BufferID != "" {
		w.id = resp.BufferID
	}
	needShow := false
	if resp.Line != w.cursorLine || resp.Column != w.cursorCol {
		w.cursorLine = resp.Line
		w.cursorCol = resp.Column
		if !w.hidden {
			if w.moveScreenVisible() {
				w.getContent()
				return
			}
			needShow = true
		}
	}
	if resp.ContentChanged {
		w.content = resp.Content
		w.syntax = resp.Syntax
		w.numLines = resp.NumLines
		if !w.hidden {
			w.preDraw()
			w.draw()
			needShow = true
		}
	}
	if needShow {
		w.drawCursor()
		w.screen.Show()
	}
}

func (w *Window) moveScreenVisible() bool {
	off := w.offset
	if w.cursorLine >= w.offset && w.cursorLine <= w.offset+w.region.Dy() {
		return false
	}
	if w.cursorLine < w.offset {
		w.offset = w.cursorLine
	} else {
		w.offset = w.cursorLine - w.region.Dy() + 1
	}
	if w.offset < 0 {
		w.offset = 0
	}
	if off != w.offset {
		return true
	}
	return false
}

func (w *Window) moveCursorVisible() {
}

func (w *Window) getContent() {
	req := editor.NewRequest(w.id, action.Action{Name: action.Content})
	req.LineOffset = w.offset
	req.LineCount = w.region.Dy()
	w.reqC <- req
}

func (w *Window) hasPoint(x, y int) bool {
	if x < w.region.Min.X || x > w.region.Max.X {
		return false
	}
	if y < w.region.Min.Y || y > w.region.Max.Y {
		return false
	}
	return true
}

func (w *Window) pointToPos(x, y int) (int, int) {
	ln := w.offset + y - w.region.Min.Y
	off := len(fmt.Sprintf("%v", w.offset+w.region.Dy()))
	if !w.noGutter {
		off++
	}
	off++ // left padding
	col := x - w.region.Min.X - off
	return ln, col
}

func (w *Window) draw() {
	w.clear()

	// Content
	off := len(fmt.Sprintf("%v", w.offset+w.region.Dy()))
	off += w.region.Min.X
	if !w.noGutter {
		off++
	}
	off++ // left padding
	yoff := w.region.Min.Y

	for i := range w.content {
		y := yoff + i
		if i < w.region.Min.Y || i > w.region.Max.Y {
			logger.Errorf("draw: Line out of range %v %v-%v", i, w.region.Min.Y, w.region.Max.Y)
			continue
		}
		coff := 0
		for ii, r := range w.content[i] {
			if r == '\t' {
				str := strings.Repeat(" ", w.tabsize)
				for _, r := range str {
					x := off + ii + coff
					if x < w.region.Min.X || x > w.region.Max.X {
						continue
					}
					w.screen.SetContent(x, y, r, nil, w.preStyled[i][ii])
				}
				coff += w.tabsize - 1
				continue
			}
			x := off + ii + coff
			if x < w.region.Min.X || x > w.region.Max.X {
				continue
			}
			w.screen.SetContent(x, y, r, nil, w.preStyled[i][ii])
		}
	}
	w.drawNumbers()
}

func (w *Window) drawCursor() {
	off := len(fmt.Sprintf("%v", w.offset+w.region.Dy()))
	if !w.noGutter {
		off++
	}
	off++ // left padding
	w.screen.ShowCursor(w.cursorCol+off, w.offset+w.cursorLine)
}

func (w *Window) preDrawClear() {
	st := styleEntry(w.pal, "default", tcell.StyleDefault)
	for i := range w.preStyled {
		for ii := range w.preStyled[i] {
			w.preStyled[i][ii] = st
		}
	}
}

func (w *Window) preDrawGen() {
	lncnt := w.region.Dy()
	colcnt := w.region.Dx()
	w.preStyled = make([][]tcell.Style, lncnt, lncnt)
	for i := range w.preStyled {
		w.preStyled[i] = make([]tcell.Style, colcnt, colcnt)
	}
	w.preDrawClear()
}

func (w *Window) preDraw() {
	w.preDrawClear()

	minl := w.offset
	maxl := w.offset + w.region.Dy()
	lps := len(w.preStyled)
	cps := len(w.preStyled[0])

	for _, sy := range w.syntax {
		ln := sy.Line()
		if ln < minl || ln > maxl || ln-w.offset < 0 || ln-w.offset >= lps {
			continue
		}
		for _, rg := range sy.Range() {
			for col := rg[0]; col < rg[1]; col++ {
				if col >= cps {
					continue
				}
				w.preStyled[ln-w.offset][col] = styleEntry(w.pal, sy.Token(), tcell.StyleDefault)
			}
		}
	}
}

func (w *Window) drawTab(x, y int, st tcell.Style) int {
	str := strings.Repeat(" ", w.tabsize)
	for _, r := range str {
		w.screen.SetContent(x, y, r, nil, st)
		x++
	}
	return x
}

func (w *Window) drawNumbers() {
	ss := len(fmt.Sprintf("%v", w.offset+w.region.Dy()))
	ff := fmt.Sprintf("%%%dd", ss)
	colOff := 0
	if !w.noGutter {
		colOff++
	}
	for i := 0; i < w.region.Dy(); i++ {
		str := fmt.Sprintf(ff, w.offset+i+1)
		for ii, r := range str {
			w.screen.SetContent(colOff+ii, i, r, nil, w.style)
		}
	}
}

// Clear will clear the status bar
func (w *Window) clear() {
	for y := w.region.Min.Y; y <= w.region.Max.Y; y++ {
		for x := w.region.Min.X; x < w.region.Max.X; x++ {
			w.screen.SetContent(x, y, ' ', nil, w.style)
		}
	}
}

// NewWindow returns a new window
func NewWindow(sc tcell.Screen, rgn image.Rectangle, pal *palette.Palette, reqC chan editor.Request) *Window {
	win := &Window{screen: sc, region: rgn, style: tcell.StyleDefault, pal: pal, tabsize: 4, reqC: reqC}
	win.preDrawGen()
	return win
}
