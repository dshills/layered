package layer

import (
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
	Add(keys []key.Keyer, actions []action.Action)
	Remove(keys []key.Keyer)
	BeginActions() []action.Action
	EndActions() []action.Action
	PartialMatchActions() []action.Action
	NoMatchActions() []action.Action
}

// Parserer will parse key strokes into actions
type Parserer interface {
	Parse(keys ...key.Keyer) ([]action.Action, ParseStatus)
}
