# layer
--
    import "."


## Usage

#### type Factory

```go
type Factory func(action.Definitions) Interpriter
```

Factory will return an Interpriter

#### type Interpriter

```go
type Interpriter interface {
	Layers() []Layer
	Active() Layer
	Match(k ...key.Keyer) ([]action.Action, error)
	Partial() string
	Status() MatchStatus
	Add(...Layer)
	Remove(name string)
	LoadDirectory(dir string) error
}
```

Interpriter will convert keystrokes into actions

#### func  NewInterpreter

```go
func NewInterpreter(ad action.Definitions, deflayer string) Interpriter
```
NewInterpreter returns an interpreter

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

	Match(keys ...key.Keyer) MatchInfo
	MatchSpecial(k key.Keyer) ([]action.Action, bool)

	Load(dl action.Definitions, r io.Reader) error
}
```

Layer is a set of keyboard / action bindings

#### func  NewLayer

```go
func NewLayer(name string) Layer
```
NewLayer will return a Layer

#### type MatchInfo

```go
type MatchInfo struct {
	Actions    []action.Action
	Status     MatchStatus
	MatchValue []key.Keyer
	Remaining  []key.Keyer
}
```

MatchInfo is information about the match

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
	ErrorMatch
	CancelMatch
)
```
ParseStatus constants

#### func (MatchStatus) String

```go
func (s MatchStatus) String() string
```
