# layer
--
    import "."


## Usage

#### type Collectioner

```go
type Collectioner interface {
	LoadDir(dir string) error
	Add(a Layerer)
	Remove(name string)
	Layer(name string) (Layerer, error)
}
```

Collectioner is a collection of layers

#### type Layer

```go
type Layer struct {
}
```

Layer is a mapping of key strokes to actions

#### func (*Layer) Add

```go
func (l *Layer) Add(keys []key.Keyer, actions []action.Action)
```
Add will add a keys / actions map

#### func (*Layer) BeginActions

```go
func (l *Layer) BeginActions() []action.Action
```
BeginActions will return the actions that are returned when switching to layer

#### func (*Layer) EndActions

```go
func (l *Layer) EndActions() []action.Action
```
EndActions will return the actions that are returned when leaving layer

#### func (*Layer) IsDefault

```go
func (l *Layer) IsDefault() bool
```
IsDefault returns true if it is the default layer

#### func (*Layer) Match

```go
func (l *Layer) Match(keys []key.Keyer) ([]action.Action, ParseStatus)
```
Match will attempt to map keys to actions

#### func (*Layer) Name

```go
func (l *Layer) Name() string
```
Name returns the layer name

#### func (*Layer) NoMatchActions

```go
func (l *Layer) NoMatchActions() []action.Action
```
NoMatchActions returns actions when keys do not match

#### func (*Layer) PartialMatchActions

```go
func (l *Layer) PartialMatchActions() []action.Action
```
PartialMatchActions returns the partial match actions

#### func (*Layer) Remove

```go
func (l *Layer) Remove(keys []key.Keyer)
```
Remove will remove a mapping

#### type Layerer

```go
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
```

Layerer is a layer

#### type Layers

```go
type Layers struct {
}
```

Layers is a set of layers

#### func (*Layers) Add

```go
func (l *Layers) Add(a Layerer)
```
Add adds a layer

#### func (*Layers) Layer

```go
func (l *Layers) Layer(name string) (Layerer, error)
```
Layer will return a layer by name

#### func (*Layers) LoadDir

```go
func (l *Layers) LoadDir(dir string) error
```
LoadDir wil load layers from a directory

#### func (*Layers) Remove

```go
func (l *Layers) Remove(name string)
```
Remove will remove a layer

#### type ParseStatus

```go
type ParseStatus int
```

ParseStatus is the status of a parser operation

```go
const (
	NoMatch ParseStatus = iota
	PartialMatch
	Match
)
```
ParseStatus constants

#### func (ParseStatus) String

```go
func (s ParseStatus) String() string
```

#### type Parserer

```go
type Parserer interface {
	Parse(keys ...key.Keyer) ([]action.Action, ParseStatus)
}
```

Parserer will parse key strokes into actions

#### type Scanner

```go
type Scanner struct {
}
```

Scanner evaluates keys within a layer

#### func  NewScanner

```go
func NewScanner(layers Collectioner, stLayer string) (*Scanner, error)
```
NewScanner returns a layer scanner

#### func (*Scanner) Init

```go
func (s *Scanner) Init()
```
Init will initialize the scanner

#### func (*Scanner) Scan

```go
func (s *Scanner) Scan(key key.Keyer) ([]action.Action, ParseStatus, error)
```
Scan will match keys in the current layer
