package editor

import (
	"fmt"
	"strings"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/conf"
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/register"
	"github.com/dshills/layered/syntax"
	"github.com/dshills/layered/textobject"
	"github.com/dshills/layered/textstore"
	"github.com/dshills/layered/undo"
)

// Editor is an editor instance
type Editor struct {
	bufs        []buffer.Bufferer
	bufFac      buffer.Factory
	curFac      cursor.Factory
	txtFac      textstore.Factory
	undoFac     undo.Factory
	synFac      syntax.Factory
	objs        textobject.Objecter
	ftd         filetype.Manager
	reg         register.Registerer
	actC        chan action.Request
	doneC       chan struct{}
	activeBufID string
	config      *conf.Configuration
	actDefs     *action.Definitions
}

// Buffers returns the editors currrent buffers
func (e *Editor) Buffers() []buffer.Bufferer { return e.bufs }

func (e *Editor) add(buf buffer.Bufferer) {
	e.bufs = append(e.bufs, buf)
}

func (e *Editor) remove(id string) error {
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

func (e *Editor) buffer(id string) (buffer.Bufferer, error) {
	for i := range e.bufs {
		if e.bufs[i].ID() == id {
			return e.bufs[i], nil
		}
	}
	return nil, fmt.Errorf("Editor.Buffer: %q Not found", id)
}

func (e *Editor) bufferIdx(id string) (int, error) {
	for i := range e.bufs {
		if e.bufs[i].ID() == id {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Editor.Buffer: Not found")
}

func (e *Editor) newBuffer() string {
	ts := e.txtFac(e.undoFac)
	buf := e.bufFac(ts, e.curFac(ts), e.synFac(e.config), e.ftd, e.reg)
	e.bufs = append(e.bufs, buf)
	return buf.ID()
}

// ActionChan returns the action channel
func (e *Editor) ActionChan() chan action.Request {
	return e.actC
}

// DoneChan returns the done channel
func (e *Editor) DoneChan() chan struct{} {
	return e.doneC
}

// New will return a new editor
func New(defs *action.Definitions, uf undo.Factory, tf textstore.Factory, bf buffer.Factory, cf cursor.Factory, sf syntax.Factory, ftf filetype.Factory, of textobject.Factory, rf register.Factory, config *conf.Configuration) (Editorer, error) {
	ed := &Editor{undoFac: uf, bufFac: bf, curFac: cf, txtFac: tf, synFac: sf, config: config, actDefs: defs}
	ed.reg = rf()
	ed.objs = of(config)

	var err error
	errs := []string{}
	ed.ftd, err = ftf(config)
	if err != nil {
		errs = append(errs, fmt.Sprintf("filetype: %v", err.Error()))
	}
	if len(errs) > 0 {
		return ed, fmt.Errorf("Editor: %v", strings.Join(errs, ", "))
	}
	return ed, nil
}
