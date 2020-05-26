package undo

// ChangeType is a the type of change made
type ChangeType int

// Change types
const (
	DeleteLine ChangeType = iota
	AddLine
	ChangeLine
)

func (c ChangeType) String() string {
	switch c {
	case DeleteLine:
		return "DeleteLine"
	case AddLine:
		return "AddLine"
	case ChangeLine:
		return "ChangeLine"
	}
	return "Unknown"
}

// Factory is a function that returns a new change set
type Factory func() ChangeSetter

// ChangeSetter is a set of changes
type ChangeSetter interface {
	AddChanges(...Changer)
	Changes() []Changer
	RemoveLine(ln int)
	AddLine(ln int)
	ChangeLine(ln int, before, after string)
}

// Changer is a change to a text store
type Changer interface {
	Cursor() []int
	Dirty() bool
	Type() ChangeType
	Line() int
	Undo(after string) string
}
