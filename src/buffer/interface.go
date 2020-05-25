package buffer

import (
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/textobject"
	"github.com/dshills/layered/textstore"
)

// Factory is a function that returns new bufferers
type Factory func(txt textstore.TextStorer, cur cursor.Cursorer) Bufferer

// Bufferer os a text buffer
type Bufferer interface {
	ID() string
	Filename() string
	SetFilename(n string)
	Filetype() string
	SetFiletype(ft string)
	TextStore() textstore.TextStorer
	Cursor() cursor.Cursorer
	Move(cnt int, obj textobject.TextObjecter) error
	MoveEnd(cnt int, obj textobject.TextObjecter) error
	MovePrev(cnt int, obj textobject.TextObjecter) error
	MovePrevEnd(cnt int, obj textobject.TextObjecter) error
	Up(int)
	Down(int)
	Prev(int)
	Next(int)
}
