package undo

// Change types
const (
	DeleteLine = iota
	AddLine
	ChangeLine
)

// Factory is a function that returns a new change set
type Factory func() ChangeSetter

// ChangeSetter is a set of changes
type ChangeSetter interface {
	Changes() []Changer
	RemoveLine(ln int)
	AddLine(ln int)
	ChangeLine(ln int, before, after string)
}

// Changer is a change to a text store
type Changer interface {
	Cursor() []int
	Dirty() bool
	Type() int
	Line() int
	Undo(after string) string
}
