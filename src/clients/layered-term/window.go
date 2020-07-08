package main

import (
	"image"
	"strings"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/editor"
	"github.com/dshills/layered/palette"
	"github.com/dshills/layered/terminal"
)

const tabsize = 2

type window struct {
	pal       *palette.Palette
	hidden    bool
	bufid     string
	writer    *terminal.WindowWriter
	startline int
	count     int
	ed        editor.Editorer
	reqC      chan editor.Request
	respC     chan editor.Response
}

func (w *window) setChan(reqC chan editor.Request, respC chan editor.Response) {
	w.reqC = reqC
	w.respC = respC
}

func (w *window) makeVisible(ln, col int) {
	if ln >= w.startline && ln <= w.startline+w.count {
		return
	}
	if ln < w.startline {
		w.startline = ln
	} else {
		w.startline = ln - w.count + 1
	}
	if w.startline < 0 {
		w.startline = 0
	}

	w.draw()
}

func (w *window) drawCursor(ln, col int) {
	w.makeVisible(ln, col)
	ln -= w.startline
	w.writer.MoveTo(ln, col)
}

func (w *window) draw() error {
	if w.hidden || w.reqC == nil || w.respC == nil {
		return nil
	}
	w.writer.TermWriter().ResetStyle()
	w.writer.ContentFgd = palette.NewColor(255, 255, 255)
	act := action.Action{
		Name:  action.Content,
		Line:  w.startline,
		Count: w.count,
	}
	w.writer.Clear()
	w.reqC <- editor.NewRequest(w.bufid, act)
	resp := <-w.respC
	if resp.Err != nil {
		return resp.Err
	}
	for i := range resp.Content {
		con := resp.Content[i]
		con = strings.ReplaceAll(con, "\t", " ")
		w.writer.WriteStringAt(i, 0, con)
	}
	cl := resp.Line - w.startline
	if cl >= 0 && cl < w.count {
		w.writer.TermWriter().To(cl, resp.Column)
	}

	empty := palette.Color{}
	empty.Transparent = true
	errs := []string{}
	var bad, good int
	for _, res := range resp.Syntax {
		ln := res.Line() - w.startline
		if ln < 0 || ln >= w.count {
			continue
		}
		en, err := w.pal.Entry(res.Token())
		if err != nil {
			errs = append(errs, err.Error())
		}
		w.writer.TermWriter().ResetStyle()
		for _, rng := range res.Range() {
			if rng[0] > rng[1] {
				bad++
				continue
			}
			good++
			con := resp.Content[ln][rng[0]:rng[1]]
			con = strings.ReplaceAll(con, "\t", " ")
			w.writer.WriteStyledStringAt(ln, rng[0], en.Fgd, empty, con)
		}
	}
	w.makeVisible(resp.Line, resp.Column)
	ln := resp.Line - w.startline
	w.writer.MoveTo(ln, resp.Column)
	return nil
}

func newWindow(bufid string, tw *terminal.TermWriter, x, y, xx, yy int, ed editor.Editorer, pal *palette.Palette) window {
	return window{
		pal:    pal,
		ed:     ed,
		bufid:  bufid,
		writer: terminal.NewWindowWriter(tw, image.Rect(x, y, xx, yy)),
		count:  yy - y - 1,
	}
}
