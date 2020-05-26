package syntax

import (
	"io"

	"github.com/dshills/layered/textstore"
)

// Factory is a function that returns new syntax matchers
type Factory func(rt string) Matcherer

// Matcherer is a collection of syntax rules
// representing a set of language syntax rules
type Matcherer interface {
	LoadFileType(ft string) error
	LoadFile(path string) error
	Add(Ruler)
	Parse(textstore.TextStorer) []Resulter
}

// Resulter is a a syntax match result
type Resulter interface {
	IsEqual(Resulter) bool
	Token() string
	Line() int
	Range() [][]int
	Priority() int
}

// Ruler is a syntax matching rule
type Ruler interface {
	Group() string
	Type() string
	Match(io.Reader, int) Resulter
}
