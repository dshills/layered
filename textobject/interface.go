package textobject

// Factory will return an Objectr
type Factory func(rts ...string) Objecter

// TextObjecter is a text object interface
type TextObjecter interface {
	Name() string
	FindAll(string) [][]int
	FindAfter(string, int) [][]int
	FindBefore(string, int) [][]int
	UseFirst() bool
	UseLast() bool
	MultiLine() bool
	Simple() bool
}

// Objecter is a set of text objects
type Objecter interface {
	SetRuntimes(rts ...string)
	AddRuntimes(rts ...string)
	LoadDir(path string) error
	Object(name string) (TextObjecter, error)
	Add(...TextObjecter)
	Remove(name string)
}
