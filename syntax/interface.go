package syntax

import (
	"github.com/dshills/layered/conf"
	"github.com/dshills/layered/textstore"
)

// Factory is a function that returns new syntax matchers
type Factory func(*conf.Configuration) Manager

// Manager is a collection of syntax rules
// representing a set of language syntax rules
type Manager interface {
	LoadFileType(ft string) error
	LoadFile(path string) error
	Add(Ruler)
	Parse(ts textstore.TextStorer, groups ...string) []Resulter
	FilterResults(results []Resulter, groups ...string) []Resulter
}

// Resulter is a a syntax match result
type Resulter interface {
	IsEqual(Resulter) bool
	Token() string
	Line() int
	Range() [][]int
	Priority() int
	SetToken(string)
	SetLine(int)
	SetRanges([][]int)
	SetPriority(int)
	AddRanges([][]int)
}

// Ruler is a syntax matching rule
type Ruler interface {
	Group() string
	Type() string
	Match(textstore.TextStorer) []Resulter
	IsDependent() bool
}
