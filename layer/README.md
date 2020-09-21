# layer
--
    import "."


## Usage

#### type Factory

```go
type Factory func() Interpriter
```

Factory will return an Interpriter

#### type Interpriter

```go
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
```

Interpriter will convert keystrokes into actions

#### func  NewInterpriter

```go
func NewInterpriter() Interpriter
```
NewInterpriter returns an interpriter

#### type KeyAction

```go
type KeyAction interface {
	Keys() []key.Keyer
	Actions() []action.Action
	Match(keys []key.Keyer) MatchStatus
}
```

KeyAction are keys that trigger actions

#### func  NewKeyAction

```go
func NewKeyAction(name string, keys []key.Keyer, acts []action.Action) KeyAction
```
NewKeyAction will return a key action structure

#### type Layer

```go
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

	Load(r io.Reader) error
}
```

Layer is a set of keyboard / action bindings

#### func  NewLayer

```go
func NewLayer(name string) Layer
```
NewLayer will return a Layer

#### type MatchStatus

```go
type MatchStatus int
```

MatchStatus is the status of a parser operation

```go
const (
	NoMatch MatchStatus = iota
	PartialMatch
	Match
)
```
ParseStatus constants

#### func (MatchStatus) String

```go
func (s MatchStatus) String() string
```
