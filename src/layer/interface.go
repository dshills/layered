package layer

import (
	"io"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

// Collectioner is a collection of layers
type Collectioner interface {
	LoadDir(dir string) error
	Add(a Layerer)
	Remove(name string)
	Layer(name string) (Layerer, error)
}

// ParseStatus is the status of a parser operation
type ParseStatus int

// ParseStatus constants
const (
	NoMatch ParseStatus = iota
	PartialMatch
	Match
)

func (s ParseStatus) String() string {
	switch s {
	case NoMatch:
		return "No match"
	case PartialMatch:
		return "Partial match"
	case Match:
		return "Match"
	}
	return "Unknown status"
}

// Layerer is a layer
type Layerer interface {
	Match(keys []key.Keyer) ([]action.Action, ParseStatus)
	Name() string
	Map(name string, keys []string, actions []action.Action) error
	UnMap(name string)
	BeginActions() []action.Action
	EndActions() []action.Action
	PartialMatchActions() []action.Action
	NoMatchActions() []action.Action
	Load(io.Reader) error
}

// Parserer will parse key strokes into actions
type Parserer interface {
	Parse(keys ...key.Keyer) ([]action.Action, ParseStatus)
}
