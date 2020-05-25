package textobject

// TextObjecter is a text object interface
type TextObjecter interface {
	Name() string
	FindAll(string) [][]int
	FindAfter(string, int) [][]int
	FindBefore(string, int) [][]int
}

// Objecter is a set of text objects
type Objecter interface {
	LoadDir(path string) error
	Object(name string) (TextObjecter, error)
	Add(...TextObjecter)
	Remove(name string)
}
