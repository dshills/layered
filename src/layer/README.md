# layer
--
    import "."


## Usage

#### type Collectioner

```go
type Collectioner interface {
	LoadLayers(dir string) error
	Add(Layerer)
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

Layer is a keyboard action map

#### func (*Layer) Add

```go
func (l *Layer) Add(keys []key.Keyer, actions []action.Action)
```
Add will map keys to actions

#### func (*Layer) BeginActions

```go
func (l *Layer) BeginActions() []action.Action
```
BeginActions returns actions that occur when switching to the layer

#### func (*Layer) EndActions

```go
func (l *Layer) EndActions() []action.Action
```
EndActions returns action that occur when switching away from layer

#### func (*Layer) Load

```go
func (l *Layer) Load(r io.Reader) error
```
Load will load a layer from a reader

#### func (*Layer) MatchActions

```go
func (l *Layer) MatchActions() []action.Action
```
MatchActions returns actions that occur when a match is made they are in
addition to key mapped actions

#### func (*Layer) Name

```go
func (l *Layer) Name() string
```
Name returns the layer's name

#### func (*Layer) NewParser

```go
func (l *Layer) NewParser() Parserer
```
NewParser returns a new key parser

#### func (*Layer) NoMatchActions

```go
func (l *Layer) NoMatchActions() []action.Action
```
NoMatchActions returns actions the occur when a match is not made

#### func (*Layer) PartialMatchActions

```go
func (l *Layer) PartialMatchActions() []action.Action
```
PartialMatchActions returns actions that occur when a partial match is made

#### func (*Layer) Remove

```go
func (l *Layer) Remove(keys []key.Keyer)
```
Remove will remove a key mapping

#### type Layerer

```go
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

#### func  SameKeys

```go
func SameKeys(a, b []key.Keyer) ParseStatus
```
SameKeys compares two key lists

#### type Parser

```go
type Parser struct {
}
```

Parser is a key stroke parser specific to a layer

#### func (*Parser) Parse

```go
func (p *Parser) Parse(keys ...key.Keyer) (actions []action.Action, status ParseStatus)
```
Parse will take key strokes and will return actions when matches

#### type Parserer

```go
type Parserer interface {
	Parse(keys ...key.Keyer) ([]action.Action, ParseStatus)
}
```

Parserer will parse key strokes into actions
