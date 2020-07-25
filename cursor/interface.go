package cursor

import "github.com/dshills/layered/textstore"

// Factory is a function that returns a new cursor
type Factory func(txt textstore.TextStorer) Cursor

// Cursor is an editor cursor
type Cursor interface {
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
	StartTrack()
	EndTrack()
	Tracked() [][]int
	SetMovePast(on bool)
	MovePast() bool
}
