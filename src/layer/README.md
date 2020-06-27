# layer
--
    import "."


## Usage

```go
const (
	Any       = "<any>"       // any key
	Printable = "<printable>" // Any printable character
	Control   = "<control>"   // any control character
	Digit     = "<digit>"     // 0-9
	Letter    = "<letter>"    // Any letter
	Lower     = "<lower>"     // Any lower case
	Upper     = "<upper>"     // Any upper case
	NonBlank  = "<non-blank>" // Any non space printable character
	Pattern   = "<pattern=>"  // regex pattern
)
```
Key matcher constants

#### type ALayer

```go
type ALayer struct {
}
```

ALayer is a layer

#### type Factory

```go
type Factory func(rtpaths ...string) (Manager, error)
```

Factory will create a layer manager

#### type Group

```go
type Group []ALayer
```

Group is a group of layers

#### type Layer

```go
type Layer struct {
}
```

Layer is a mapping of key strokes to actions

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

#### func (*Layer) Load

```go
func (l *Layer) Load(r io.Reader) error
```
Load will load a layer from a reader

#### func (*Layer) Map

```go
func (l *Layer) Map(name string, keys []string, actions []action.Action) error
```
Map will add a keys / actions mapping

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

#### func (*Layer) Unmap

```go
func (l *Layer) Unmap(name string)
```
Unmap will remove a mapping

#### type Layerer

```go
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
func (l *Layers) Add(a ...Layerer)
```
Add adds a layer

#### func (*Layers) AddRuntime

```go
func (l *Layers) AddRuntime(rtpaths ...string) error
```
AddRuntime adds a runtime path

#### func (*Layers) Layer

```go
func (l *Layers) Layer(name string) (Layerer, error)
```
Layer will return a layer by name

#### func (*Layers) Load

```go
func (l *Layers) Load() error
```
Load will load the layers within the runtimes

#### func (*Layers) Remove

```go
func (l *Layers) Remove(name string)
```
Remove will remove a layer

#### func (*Layers) RemoveRuntime

```go
func (l *Layers) RemoveRuntime(path string) error
```
RemoveRuntime will remove a runtime path

#### type Manager

```go
type Manager interface {
	AddRuntime(rtpaths ...string) error
	RemoveRuntime(path string) error
	Load() error
	Add(a ...Layerer)
	Remove(name string)
	Layer(name string) (Layerer, error)
}
```

Manager is a collection of managed layers

#### func  New

```go
func New(rtpaths ...string) (Manager, error)
```
New will return a layer manager

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

#### type Parser

```go
type Parser struct {
}
```

Parser is a command parser

#### func (*Parser) Parse

```go
func (p *Parser) Parse(k key.Keyer) []action.Action
```
Parse will parse a key

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
func NewScanner(layers Manager, stLayer string) (*Scanner, error)
```
NewScanner returns a layer scanner

#### func (*Scanner) Init

```go
func (s *Scanner) Init()
```
Init will initialize the scanner

#### func (*Scanner) LayerName

```go
func (s *Scanner) LayerName() string
```
LayerName will return name of the current layer

#### func (*Scanner) Partial

```go
func (s *Scanner) Partial() string
```
Partial returns the partial keys

#### func (*Scanner) Scan

```go
func (s *Scanner) Scan(key key.Keyer) ([]action.Action, ParseStatus, error)
```
Scan will match keys in the current layer
