package terminal

import (
	"fmt"
	"image"
	"strings"

	"github.com/dshills/layered/palette"
)

// WindowWriter is a terminal writer
// supporting windowing
type WindowWriter struct {
	tw              *TermWriter
	region          image.Rectangle
	contentRegion   image.Rectangle
	bordered        bool
	paddingTop      int
	paddingBottom   int
	paddingLeft     int
	paddingRight    int
	horizChar       rune
	verticalChar    rune
	topLeftChar     rune
	topRightChar    rune
	bottomLeftChar  rune
	bottomRightChar rune
	BorderFgd       palette.Color
	BorderBck       palette.Color
	ContentFgd      palette.Color
	ContentBck      palette.Color
}

// NewWindowWriter returns a writer
func NewWindowWriter(w *TermWriter, reg image.Rectangle) *WindowWriter {
	return &WindowWriter{
		tw:              w,
		region:          reg,
		contentRegion:   image.Rect(reg.Min.X+1, reg.Min.Y+1, reg.Max.X-1, reg.Max.Y-1),
		bordered:        true,
		horizChar:       '─',
		verticalChar:    '│',
		topRightChar:    '┐',
		topLeftChar:     '┌',
		bottomLeftChar:  '└',
		bottomRightChar: '┘',
		BorderFgd:       palette.Color{Red: 122, Green: 122, Blue: 122},
		BorderBck:       palette.Color{Transparent: true},
		ContentFgd:      palette.Color{Transparent: true},
		ContentBck:      palette.Color{Transparent: true},
	}
}

// TermWriter will return the term writer
func (ww *WindowWriter) TermWriter() *TermWriter { return ww.tw }

// SetBorderChars will set the border characters
func (ww *WindowWriter) SetBorderChars(vert, horz, tl, tr, bl, br rune) {
	ww.horizChar = horz
	ww.verticalChar = vert
	ww.topLeftChar = tl
	ww.topRightChar = tr
	ww.bottomLeftChar = bl
	ww.bottomRightChar = br
}

// SetPadding will set the window padding
func (ww *WindowWriter) SetPadding(top, right, bottom, left int) {
	ww.paddingTop = top
	ww.paddingBottom = bottom
	ww.paddingLeft = left
	ww.paddingRight = right
}

// Resize will move the window to reg
func (ww *WindowWriter) Resize(reg image.Rectangle) {
	ww.region = reg
	if !ww.bordered {
		ww.contentRegion = reg
	}
	ww.contentRegion = image.Rect(reg.Min.X+1, reg.Min.Y+1, reg.Max.X-1, reg.Max.Y-1)
}

// SetBordered turns the border on / off
func (ww *WindowWriter) SetBordered(on bool) {
	ww.bordered = on
	ww.Resize(ww.region)
}

// WriteStyledStringAt will write a styled string at a location
// a negative value for any r,g,b will not set fgd, bkd where it appears
func (ww WindowWriter) WriteStyledStringAt(line, col int, fgd, bck palette.Color, s string) (int, error) {
	if !fgd.Transparent {
		ww.tw.SetForeground(fgd)
	}
	if !bck.Transparent {
		ww.tw.SetBackground(bck)
	}
	l, err := ww.writeAt(line, col, s)
	ww.tw.ResetStyle()
	return l, err
}

// WriteStringAt will write a string relative to the window
func (ww WindowWriter) WriteStringAt(line, col int, s string) (int, error) {
	ww.tw.SetForeground(ww.ContentFgd)
	ww.tw.SetBackground(ww.ContentBck)
	return ww.writeAt(line, col, s)
}

func (ww WindowWriter) writeAt(line, col int, s string) (int, error) {
	if line > ww.region.Dy()-ww.paddingTop-ww.paddingBottom {
		return 0, fmt.Errorf("Line out of range %v", line)
	}
	if col > ww.region.Dx()-ww.paddingLeft-ww.paddingRight {
		return 0, fmt.Errorf("Column out of range %v", col)
	}

	ls := ww.contentRegion.Min.Y + ww.paddingTop
	cs := ww.contentRegion.Min.X + ww.paddingLeft
	ce := cs + ww.contentRegion.Dx() - ww.paddingRight
	ww.tw.To(ls+line, cs+col)
	if len(s)+col >= ce-cs {
		s = string([]rune(s)[:ce-cs])
	}
	l, err := ww.tw.WriteString(s)
	return l, err
}

// MoveTo will move the cursor to position
func (ww WindowWriter) MoveTo(line, col int) {
	lmax := ww.region.Dy() - ww.paddingTop - ww.paddingBottom
	cmax := ww.region.Dx() - ww.paddingLeft - ww.paddingRight
	if line < 0 {
		line = 0
	}
	if line > lmax {
		line = lmax
	}
	if col < 0 {
		col = 0
	}
	if col > cmax {
		col = cmax
	}
	ls := ww.contentRegion.Min.Y + ww.paddingTop
	cs := ww.contentRegion.Min.X + ww.paddingLeft
	ww.tw.To(ls+line, cs+col)
}

// Clear will clear the content area
func (ww WindowWriter) Clear() {
	ww.Fill(' ')
}

// Fill will fill the content area, including padding areas
// with character r
func (ww WindowWriter) Fill(r rune) {
	ww.tw.SavePos()
	ww.tw.SetForeground(ww.ContentFgd)
	ww.tw.SetBackground(ww.ContentBck)
	ls := ww.contentRegion.Min.Y
	le := ls + ww.contentRegion.Dy()
	cs := ww.contentRegion.Min.X
	ce := cs + ww.contentRegion.Dx()
	for i := ls; i <= le; i++ {
		ww.tw.To(i, cs)
		ww.tw.WriteString(strings.Repeat(string(r), ce-cs+1))
	}
	ww.tw.ResetStyle()
	ww.tw.RestorePos()
}

// Draw will draw the window
func (ww WindowWriter) Draw() {
	ww.Clear()
	if !ww.bordered {
		return
	}
	ww.tw.SavePos()
	ww.tw.SetForeground(ww.BorderFgd)
	ww.tw.SetBackground(ww.BorderBck)
	ww.tw.To(ww.region.Min.Y, ww.region.Min.X)
	for i := 0; i < ww.region.Dx(); i++ {
		ww.tw.WriteRune(ww.horizChar)
	}
	ww.tw.To(ww.region.Max.Y, ww.region.Min.X)
	for i := 0; i < ww.region.Dx(); i++ {
		ww.tw.WriteRune(ww.horizChar)
	}
	for i := 1; i < ww.region.Dy(); i++ {
		ww.tw.To(ww.region.Min.Y+i, ww.region.Min.X)
		ww.tw.WriteRune(ww.verticalChar)
	}
	for i := 1; i < ww.region.Dy(); i++ {
		ww.tw.To(ww.region.Min.Y+i, ww.region.Max.X)
		ww.tw.WriteRune(ww.verticalChar)
	}
	ww.tw.To(ww.region.Min.Y, ww.region.Min.X)
	ww.tw.WriteRune(ww.topLeftChar)
	ww.tw.To(ww.region.Min.Y, ww.region.Max.X)
	ww.tw.WriteRune(ww.topRightChar)
	ww.tw.To(ww.region.Max.Y, ww.region.Min.X)
	ww.tw.WriteRune(ww.bottomLeftChar)
	ww.tw.To(ww.region.Max.Y, ww.region.Max.X)
	ww.tw.WriteRune(ww.bottomRightChar)
	ww.tw.ResetStyle()
	ww.tw.RestorePos()
}
