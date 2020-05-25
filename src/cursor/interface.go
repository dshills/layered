package cursor

import "github.com/dshills/layered/textstore"

// Factory is a function that returns a new cursor
type Factory func(txt textstore.TextStorer) Cursorer

// Cursorer is an editor cursor
type Cursorer interface {
	AsRange() []int
	Line() int
	Column() int
	GotoLine(ln int) bool
	Top() bool
	Bottom() bool
	Get() (int, int)
	Down(cnt int) bool
	Up(cnt int) bool
	Prev(cnt int) bool
	Next(cnt int) bool
	MoveValid(line, col int) bool
}
