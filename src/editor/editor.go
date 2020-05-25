package editor

import (
	"fmt"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/syntax"
	"github.com/dshills/layered/textobject"
	"github.com/dshills/layered/textstore"
)

// Editor is an editor instance
type Editor struct {
	bufs  []buffer.Bufferer
	bfac  buffer.Factory
	cfac  cursor.Factory
	tsfac textstore.Factory
	ftd   filetype.Detecter
	sfac  syntax.Factory
	objs  textobject.Objecter
}

// Buffers returns the editors currrent buffers
func (e *Editor) Buffers() []buffer.Bufferer { return e.bufs }

// Add will add a buffer to the editor
func (e *Editor) Add(buf buffer.Bufferer) {
	e.bufs = append(e.bufs, buf)
}

// Remove will remove a buffer from the editor
func (e *Editor) Remove(id string) {}

// Buffer will return a buffer by id
func (e *Editor) Buffer(id string) (buffer.Bufferer, error) {
	for i := range e.bufs {
		if e.bufs[i].ID() == id {
			return e.bufs[i], nil
		}
	}
	return nil, fmt.Errorf("Editor.Buffer: Not found")
}

// Exec will execute a transaction in the editor
func (e *Editor) Exec(tr action.Transactioner) error {
	acts := tr.Actions()
	for i := range acts {
		switch acts[i].Name() {
		case action.Move:
			obj, err := e.objs.Object(acts[i].Target())
			if err != nil {
				return err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.Move(1, obj)
		case action.MoveEnd:
			obj, err := e.objs.Object(acts[i].Target())
			if err != nil {
				return err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.MoveEnd(1, obj)
		case action.MovePrev:
			obj, err := e.objs.Object(acts[i].Target())
			if err != nil {
				return err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.MovePrev(1, obj)
		case action.MovePrevEnd:
			obj, err := e.objs.Object(acts[i].Target())
			if err != nil {
				return err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.MovePrevEnd(1, obj)
		case action.Up:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.Up(1)
		case action.Down:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.Down(1)
		case action.Prev:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.Prev(1)
		case action.Next:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.Next(1)
		}

	}
	return nil
}

// New will return a new editor
func New(bf buffer.Factory, cf cursor.Factory, tf textstore.Factory, sf syntax.Factory, ftd filetype.Detecter, objs textobject.Objecter) Editorer {
	return &Editor{bfac: bf, cfac: cf, tsfac: tf, ftd: ftd, sfac: sf, objs: objs}
}
