package layer

import (
	"io"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

// Factory will create a layer manager
type Factory func(rtpaths ...string) (Manager, error)

// Manager is a collection of managed layers
type Manager interface {
	AddRuntime(rtpaths ...string) error
	RemoveRuntime(path string) error
	Load() error
	Add(a ...Layerer)
	Remove(name string)
	Layer(name string) (Layerer, error)
}

// Layerer is a layer
type Layerer interface {
	Match(keys []key.Keyer) ([]action.Action, ParseStatus)
	Name() string
	Map(name string, keys []string, actions []action.Action) error
	Unmap(name string)
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
