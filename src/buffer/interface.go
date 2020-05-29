package buffer

import (
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/register"
	"github.com/dshills/layered/syntax"
	"github.com/dshills/layered/textobject"
	"github.com/dshills/layered/textstore"
)

// Factory is a function that returns new bufferers
type Factory func(txt textstore.TextStorer, cur cursor.Cursorer, m syntax.Matcherer, ftd filetype.Detecter, reg register.Registerer) Bufferer

// Bufferer is a text buffer
type Bufferer interface {
	ID() string
	TextStore() textstore.TextStorer
	Cursor() cursor.Cursorer
	Filer
	Mover
	TextEditor
	Selector
	SyntaxResults() []syntax.Resulter
	SearchResults() []SearchResult
	Search(string) ([]SearchResult, error)
}

// Filer is file functions
type Filer interface {
	Filename() string
	SetFilename(n string)
	Filetype() string
	SetFiletype(ft string)
	SaveBuffer(path string) error
	OpenFile(path string) error
	RenameFile(path string) error
	Dirty() bool
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
	Reset(string)
	ReplaceObject(line, col int, obj textobject.TextObjecter, s string, cnt int) error
	DeleteChar(line, col, cnt int) error
	DeleteCharBack(line, col, cnt int) error
	DeleteLine(line, cnt int) error
	DeleteObject(line, col int, obj textobject.TextObjecter, cnt int) error
	InsertString(line, col int, st string) error
	NewLineAbove(line int, st string, cnt int) error
	NewLineBelow(line int, st string, cnt int) error
	Indent(ln, cnt int) error
	Outdent(ln, cnt int) error
	Undo() error
	Redo() error
	StartGroupUndo()
	StopGroupUndo()
}

// Selector is selection functions
type Selector interface {
	BeginSelect()
	EndSelect()
	Selection() [][]int
}
