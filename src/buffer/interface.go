package buffer

import (
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/textobject"
	"github.com/dshills/layered/textstore"
)

// Factory is a function that returns new bufferers
type Factory func(txt textstore.TextStorer, cur cursor.Cursorer) Bufferer

// Bufferer is a text buffer
type Bufferer interface {
	ID() string
	TextStore() textstore.TextStorer
	Cursor() cursor.Cursorer
	Filer
	Mover
	TextEditor
	Selector
}

// Filer is file functions
type Filer interface {
	Filename() string
	SetFilename(n string)
	Filetype() string
	SetFiletype(ft string)
}

// Mover is cursor movement functions
type Mover interface {
	Move(cnt int, obj textobject.TextObjecter) error
	MoveEnd(cnt int, obj textobject.TextObjecter) error
	MovePrev(cnt int, obj textobject.TextObjecter) error
	MovePrevEnd(cnt int, obj textobject.TextObjecter) error
	MoveTo(line, col int) error
	Up(int)
	Down(int)
	Prev(int)
	Next(int)
	ScrollDown()
	ScrollUp()
	Position() []int
}

// TextEditor is text editing functions
type TextEditor interface {
	DeleteObject(line, col int, obj textobject.TextObjecter) error
	NewLineAbove(line int, st string) error
	NewLineBelow(line int, st string) error
	DeleteChar(line, col int) error
	DeleteCharBack(line, col int) error
	InsertString(line, col int, st string) error
	DeleteLine(line int) error
	Indent(ln, cnt int) error
	Outdent(ln, cnt int) error
}

// Selector is selection functions
type Selector interface {
	BeginSelect()
	EndSelect()
	Selection() [][]int
}
