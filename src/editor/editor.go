package editor

import (
	"fmt"

	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/syntax"
	"github.com/dshills/layered/textobject"
	"github.com/dshills/layered/textstore"
	"github.com/dshills/layered/undo"
)

// Editor is an editor instance
type Editor struct {
	runtime string
	bufs    []buffer.Bufferer
	bufFac  buffer.Factory
	curFac  cursor.Factory
	txtFac  textstore.Factory
	undoFac undo.Factory
	synFac  syntax.Factory
	ftd     filetype.Detecter
	objs    textobject.Objecter
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

func (e *Editor) bufferIdx(id string) (int, error) {
	for i := range e.bufs {
		if e.bufs[i].ID() == id {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Editor.Buffer: Not found")
}

func (e *Editor) newBuffer() {
	ts := e.txtFac(e.undoFac)
	e.bufs = append(e.bufs, e.bufFac(ts, e.curFac(ts), e.synFac(e.runtime)))
}

func (e *Editor) removeBuffer(id string) error {
	i, err := e.bufferIdx(id)
	if err != nil {
		return err
	}
	// does not maintian order
	e.bufs[i] = e.bufs[len(e.bufs)-1]        // Copy last element to index i.
	e.bufs[len(e.bufs)-1] = &buffer.Buffer{} // Erase last element (write zero value).
	e.bufs = e.bufs[:len(e.bufs)-1]          // Truncate slice.
	return nil
}

// New will return a new editor
func New(
	uf undo.Factory,
	tf textstore.Factory,
	bf buffer.Factory,
	cf cursor.Factory,
	sf syntax.Factory,
	ftd filetype.Detecter,
	objs textobject.Objecter,
) Editorer {
	return &Editor{undoFac: uf, bufFac: bf, curFac: cf, txtFac: tf, ftd: ftd, synFac: sf, objs: objs}
}
