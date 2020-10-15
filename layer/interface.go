package layer

import (
	"io"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

// Factory will return an Interpriter
type Factory func(action.Definitions) Interpriter

// KeyAction are keys that trigger actions
type KeyAction interface {
	Keys() []key.Keyer
	Actions() []action.Action
	Match(keys []key.Keyer) MatchStatus
}

// Layer is a set of keyboard / action bindings
type Layer interface {
	Name() string
	Map(name string, keys []key.Keyer, actions []action.Action)
	UnMap(name string)

	Editable() bool
	AllowCursorPastEnd() bool
	WaitForComplete() bool
	NotStacked() bool

	CancelKey() key.Keyer
	PrevLayerKey() key.Keyer
	CompleteKey() key.Keyer

	OnAnyKey() []action.Action
	OnPrintableKey() []action.Action
	OnNonPritableKey() []action.Action
	OnEnterLayer() []action.Action
	OnExitLayer() []action.Action
	OnComplete() []action.Action
	OnMatch() []action.Action
	OnNoMatch() []action.Action
	OnPartialMatch() []action.Action

	KeyActions() []KeyAction

	Match(keys []key.Keyer) ([]action.Action, MatchStatus)
	MatchSpecial(k key.Keyer) ([]action.Action, bool)

	Load(dl action.Definitions, r io.Reader) error
}

// Interpriter will convert keystrokes into actions
type Interpriter interface {
	Layers() []Layer
	Active() Layer
	Match(k ...key.Keyer) []action.Action
	Partial() string
	Status() MatchStatus
	Add(...Layer)
	Remove(name string)
	LoadDirectory(dir string) error
}
