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
	Default() Layerer
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

#### func (*Layer) PartialAsParam

```go
func (l *Layer) PartialAsParam() bool
```
PartialAsParam returns true if the keys should be used as an action parameter
requires a trigger key

#### func (*Layer) PartialIncludeTrigger

```go
func (l *Layer) PartialIncludeTrigger() bool
```
PartialIncludeTrigger will add the trigger to the param

#### func (*Layer) PartialMatchActions

```go
func (l *Layer) PartialMatchActions() []action.Action
```
PartialMatchActions returns the partial match actions

#### func (*Layer) PartialTrigger

```go
func (l *Layer) PartialTrigger() key.Keyer
```
PartialTrigger will trigger a match using previous partial keys as a parameter

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
	IsDefault() bool
	PartialAsParam() bool
	PartialIncludeTrigger() bool
	PartialTrigger() key.Keyer
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

#### func (*Layers) Default

```go
func (l *Layers) Default() Layerer
```
Default will return the default layer

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

#### type OldLayer

```go
type OldLayer struct {
}
```

OldLayer is a keyboard action map

#### func (*OldLayer) Add

```go
func (l *OldLayer) Add(keys []key.Keyer, actions []action.Action)
```
Add will map keys to actions

#### func (*OldLayer) BeginActions

```go
func (l *OldLayer) BeginActions() []action.Action
```
BeginActions returns actions that occur when switching to the layer

#### func (*OldLayer) EndActions

```go
func (l *OldLayer) EndActions() []action.Action
```
EndActions returns action that occur when switching away from layer

#### func (*OldLayer) IsDefault

```go
func (l *OldLayer) IsDefault() bool
```
IsDefault returns true if this is the default layer

#### func (*OldLayer) MatchActions

```go
func (l *OldLayer) MatchActions() []action.Action
```
MatchActions returns actions that occur when a match is made they are in
addition to key mapped actions

#### func (*OldLayer) Name

```go
func (l *OldLayer) Name() string
```
Name returns the layer's name

#### func (*OldLayer) NewParser

```go
func (l *OldLayer) NewParser() Parserer
```
NewParser returns a new key parser

#### func (*OldLayer) NoMatchActions

```go
func (l *OldLayer) NoMatchActions() []action.Action
```
NoMatchActions returns actions the occur when a match is not made

#### func (*OldLayer) PartialMatchActions

```go
func (l *OldLayer) PartialMatchActions() []action.Action
```
PartialMatchActions returns actions that occur when a partial match is made

#### func (*OldLayer) Remove

```go
func (l *OldLayer) Remove(keys []key.Keyer)
```
Remove will remove a key mapping

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

#### func (ParseStatus) String

```go
func (s ParseStatus) String() string
```

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

#### type Scanner

```go
type Scanner struct {
}
```

Scanner evaluates keys within a layer

#### func  NewScanner

```go
func NewScanner(layers Collectioner) (*Scanner, error)
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
