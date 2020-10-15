package editor

import (
	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/syntax"
)

// Response is a exec response
type Response struct {
	BufferID       string
	Action         action.Action
	Line, Column   int
	Dirty          bool
	Filename       string
	Filetype       string
	NumLines       int
	Results        []KeyValue
	Content        []string
	Syntax         []syntax.Resulter
	Search         []buffer.SearchResult
	ContentChanged bool
	CursorChanged  bool
	NewBuffer      bool
	CloseBuffer    bool
	InfoChanged    bool
	Quit           bool
	Err            error
}
