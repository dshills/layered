package layer

import (
	"io"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

// Collectioner is a collection of layers
type Collectioner interface {
	LoadLayers(dir string) error
	Add(Layerer)
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

// Layerer is a layer
type Layerer interface {
	Name() string
	Add(keys []key.Keyer, actions []action.Action)
	Remove(keys []key.Keyer)
	NewParser() Parserer
	BeginActions() []action.Action
	EndActions() []action.Action
	PartialMatchActions() []action.Action
	NoMatchActions() []action.Action
	MatchActions() []action.Action
	Load(io.Reader) error
}

// Parserer will parse key strokes into actions
type Parserer interface {
	Parse(keys ...key.Keyer) ([]action.Action, ParseStatus)
}
