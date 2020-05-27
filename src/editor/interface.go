package editor

import (
	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
)

// Editorer represents an editor
type Editorer interface {
	Buffers() []buffer.Bufferer
	Add(buffer.Bufferer)
	Remove(id string)
	Buffer(id string) (buffer.Bufferer, error)
	Exec(action.Transactioner) (*Response, error)
}
