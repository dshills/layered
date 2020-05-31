package editor

import (
	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
)

// Editorer represents an editor
type Editorer interface {
	Buffers() []buffer.Bufferer
	Add(buffer.Bufferer)
	Remove(id string) error
	Buffer(id string) (buffer.Bufferer, error)
	Exec(bufid string, actions ...action.Action) (*Response, error)
}
