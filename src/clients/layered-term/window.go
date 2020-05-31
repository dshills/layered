package main

import (
	"image"
	"strings"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/editor"
	"github.com/dshills/layered/logger"
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
}

func (w *window) draw() error {
	if w.hidden {
		return nil
	}
	logger.Debugf("Drawing Window line %v count %v - %v", w.startline, w.count, w.bufid)
	w.writer.TermWriter().ResetStyle()
	w.writer.ContentFgd = palette.NewColor(255, 255, 255)
	trans := action.NewTransaction(w.bufid)
	act := action.New(action.Content)
	act.SetLine(w.startline)
	act.SetCount(w.count)
	trans.Set(act)
	resp, err := w.ed.Exec(trans)
	if err != nil {
		return err
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
			con := resp.Content[ln][rng[0]:rng[1]]
			con = strings.ReplaceAll(con, "\t", " ")
			w.writer.WriteStyledStringAt(ln, rng[0], en.Fgd, empty, con)
		}
	}
	return nil
}

func newWindow(bufid string, tw *terminal.TermWriter, x, y, xx, yy int, ed editor.Editorer, pal *palette.Palette) window {
	return window{
		pal:    pal,
		ed:     ed,
		bufid:  bufid,
		writer: terminal.NewWindowWriter(tw, image.Rect(x, y, xx, yy)),
		count:  yy - y,
	}
}
