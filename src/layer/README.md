# layer
--
    import "."


## Usage

```go
const (
	Alta          = "<alt-a>"
	Altb          = "<alt-b>"
	Altc          = "<alt-c>"
	Altd          = "<alt-d>"
	Alte          = "<alt-e>"
	Altf          = "<alt-f>"
	Altg          = "<alt-g>"
	Alth          = "<alt-h>"
	Alti          = "<alt-i>"
	Altj          = "<alt-j>"
	Altk          = "<alt-k>"
	Altl          = "<alt-l>"
	Altm          = "<alt-m>"
	Altn          = "<alt-m>"
	Alto          = "<alt-o>"
	Altp          = "<alt-p>"
	Altq          = "<alt-q>"
	Altr          = "<alt-r>"
	Alts          = "<alt-s>"
	Altt          = "<alt-t>"
	Altu          = "<alt-u>"
	Altv          = "<alt-v>"
	Altw          = "<alt-w>"
	Altx          = "<alt-x>"
	Alty          = "<alt-y>"
	Altz          = "<alt-z>"
	Alt0          = "<alt-0>"
	Alt1          = "<alt-1>"
	Alt2          = "<alt-2>"
	Alt3          = "<alt-3>"
	Alt4          = "<alt-4>"
	Alt5          = "<alt-5>"
	Alt6          = "<alt-6>"
	Alt7          = "<alt-7>"
	Alt8          = "<alt-8>"
	Alt9          = "<alt-9>"
	AltBang       = "<alt-!>"
	AltAt         = "<alt-@>"
	AltPound      = "<alt-#>"
	AltDollar     = "<alt-$>"
	AltPercent    = "<alt-%>"
	AltCarrot     = "<alt-^>"
	AltAnd        = "<alt-&>"
	AltStar       = "<alt-*>"
	AltLeftParan  = "<alt-(>"
	AltRightParan = "<alt-)>"
)
```
Alt key constants

```go
const (
	RunAlta          = rune(229)
	RunAltb          = rune(8747)
	RunAltc          = rune(231)
	RunAltd          = rune(8706)
	RunAlte          = rune(180)
	RunAltf          = rune(402)
	RunAltg          = rune(169)
	RunAlth          = rune(729)
	RunAlti          = rune(710)
	RunAltj          = rune(8710)
	RunAltk          = rune(730)
	RunAltl          = rune(172)
	RunAltm          = rune(181)
	RunAltn          = rune(732)
	RunAlto          = rune(248)
	RunAltp          = rune(960)
	RunAltq          = rune(339)
	RunAltr          = rune(174)
	RunAlts          = rune(223)
	RunAltt          = rune(8224)
	RunAltu          = rune(168)
	RunAltv          = rune(8730)
	RunAltw          = rune(8721)
	RunAltx          = rune(8776)
	RunAlty          = rune(92)
	RunAltz          = rune(937)
	RunAlt0          = rune(186)
	RunAlt1          = rune(161)
	RunAlt2          = rune(8482)
	RunAlt3          = rune(163)
	RunAlt4          = rune(162)
	RunAlt5          = rune(8734)
	RunAlt6          = rune(167)
	RunAlt7          = rune(182)
	RunAlt8          = rune(8226)
	RunAlt9          = rune(170)
	RunAltBang       = rune(8260)
	RunAltAt         = rune(8364)
	RunAltPound      = rune(8249)
	RunAltDollar     = rune(8250)
	RunAltPercent    = rune(64257)
	RunAltCarrot     = rune(64258)
	RunAltAnd        = rune(8225)
	RunAltStar       = rune(176)
	RunAltLeftParan  = rune(183)
	RunAltRightParan = rune(8218)
)
```
Alt rune codes

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
Special special keys these are artificial keys used for layer matching

```go
const (
	KeyF1 = 0xFFFF - iota
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyInsert
	KeyDelete
	KeyHome
	KeyEnd
	KeyPgup
	KeyPgdn
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyEnter     = 0x0D
	KeyEsc       = 0x1B
	KeySpace     = 0x20
	KeyTab       = 0x09
	KeyBackspace = 0x08
)
```
ctrl key constants

```go
const (
	CtrlA            = "<ctrl-a>"
	CtrlB            = "<ctrl-b>"
	CtrlC            = "<ctrl-c>"
	CtrlD            = "<ctrl-d>"
	CtrlE            = "<ctrl-e>"
	CtrlF            = "<ctrl-f>"
	CtrlG            = "<ctrl-g>"
	CtrlH            = "<ctrl-h>"
	CtrlI            = "<ctrl-i>"
	CtrlJ            = "<ctrl-j>"
	CtrlK            = "<ctrl-k>"
	CtrlL            = "<ctrl-l>"
	CtrlM            = "<ctrl-m>"
	CtrlN            = "<ctrl-n>"
	CtrlO            = "<ctrl-o>"
	CtrlP            = "<ctrl-p>"
	CtrlQ            = "<ctrl-q>"
	CtrlR            = "<ctrl-r>"
	CtrlS            = "<ctrl-s>"
	CtrlT            = "<ctrl-t>"
	CtrlU            = "<ctrl-u>"
	CtrlV            = "<ctrl-v>"
	CtrlW            = "<ctrl-w>"
	CtrlX            = "<ctrl-x>"
	CtrlY            = "<ctrl-y>"
	CtrlZ            = "<ctrl-z>"
	Ctrl0            = "<ctrl-0>"
	Ctrl1            = "<ctrl-1>"
	Ctrl2            = "<ctrl-2>"
	Ctrl3            = "<ctrl-3>"
	Ctrl4            = "<ctrl-4>"
	Ctrl5            = "<ctrl-5>"
	Ctrl6            = "<ctrl-6>"
	Ctrl7            = "<ctrl-7>"
	Ctrl8            = "<ctrl-8>"
	Ctrl9            = "<ctrl-9>"
	CtrlTilde        = "<ctrl-~"
	CtrlSlash        = "<ctrl-/>"
	CtrlSpace        = "<ctrl- >"
	CtrlLeftBracket  = "<ctrl-[>"
	CtrlBackslash    = "<ctrl-\\>"
	CtrlRightBracket = "<ctrl-]>"
	CtrlCarrot       = "<ctrl-^>"
	CtrlUnderscore   = "<ctrl-_>"
)
```
Control constants

```go
const (
	Space     = "<space>"
	Nul       = "<nul>"
	Soh       = "<soh>"
	Stx       = "<stx>"
	Etx       = "<etx>"
	Eot       = "<eot>"
	Enq       = "<enq>"
	Ack       = "<ack>"
	Bel       = "<bel>"
	Backspace = "<bs>"
	Tab       = "<tab>"
	LF        = "<lf>"
	VT        = "<vt>"
	FF        = "<ff>"
	Enter     = "<cr>"
	SO        = "<so>"
	SI        = "<si>"
	Dle       = "<dle>"
	Dc1       = "<dc1>"
	Dc2       = "<dc2>"
	Dc3       = "<dc3>"
	Dc4       = "<dc4>"
	Nak       = "<nak>"
	Syn       = "<syn>"
	Etb       = "<etb>"
	Can       = "<can>"
	Em        = "<em>"
	Sub       = "<sub>"
	Esc       = "<esc>"
	Fs        = "<fs>"
	Gs        = "<gs>"
	Rs        = "<rs>"
	Us        = "<us>"
	Up        = "<up>"
	Down      = "<down>"
	Right     = "<right>"
	Left      = "<left>"
	Upleft    = "<upleft>"
	Upright   = "<upright>"
	Downleft  = "<downleft>"
	Downright = "<downright>"
	Center    = "<center>"
	Pgup      = "<pgup>"
	Pgdn      = "<pgdn>"
	Home      = "<home>"
	End       = "<end>"
	Insert    = "<insert>"
	Delete    = "<delete>"
	Help      = "<help>"
	Exit      = "<exit>"
	Clear     = "<clear>"
	Cancel    = "<cancel>"
	Print     = "<print>"
	Pause     = "<pause>"
	Backtab   = "<backtab>"
	F1        = "<f1>"
	F2        = "<f2>"
	F3        = "<f3>"
	F4        = "<f4>"
	F5        = "<f5>"
	F6        = "<f6>"
	F7        = "<f7>"
	F8        = "<f8>"
	F9        = "<f9>"
	F10       = "<f10>"
	F11       = "<f11>"
	F12       = "<f12>"
	F13       = "<f13>"
	F14       = "<f14>"
	F15       = "<f15>"
	F16       = "<f16>"
	F17       = "<f17>"
	F18       = "<f18>"
	F19       = "<f19>"
	F20       = "<f20>"
	F21       = "<f21>"
	F22       = "<f22>"
	F23       = "<f23>"
	F24       = "<f24>"
	F25       = "<f25>"
	F26       = "<f26>"
	F27       = "<f27>"
	F28       = "<f28>"
	F29       = "<f29>"
	F30       = "<f30>"
	F31       = "<f31>"
	F32       = "<f32>"
	F33       = "<f33>"
	F34       = "<f34>"
	F35       = "<f35>"
	F36       = "<f36>"
	F37       = "<f37>"
	F38       = "<f38>"
	F39       = "<f39>"
	F40       = "<f40>"
	F41       = "<f41>"
	F42       = "<f42>"
	F43       = "<f43>"
	F44       = "<f44>"
	F45       = "<f45>"
	F46       = "<f46>"
	F47       = "<f47>"
	F48       = "<f48>"
	F49       = "<f49>"
	F50       = "<f50>"
	F51       = "<f51>"
	F52       = "<f52>"
	F53       = "<f53>"
	F54       = "<f54>"
	F55       = "<f55>"
	F56       = "<f56>"
	F57       = "<f57>"
	F58       = "<f58>"
	F59       = "<f59>"
	F60       = "<f60>"
	F61       = "<f61>"
	F62       = "<f62>"
	F63       = "<f63>"
	F64       = "<f64>"
	Del       = "<del>"
)
```
Special key constants

#### func  StrToKey

```go
func StrToKey(s string) (r rune, k int, err error)
```
StrToKey converts a string representaion to a rune, key

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
func (l *Layer) Add(keys []string, actions []action.Action) error
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
func (l *Layer) Remove(keys []string)
```
Remove will remove a mapping

#### type Layerer

```go
type Layerer interface {
	Match(keys []key.Keyer) ([]action.Action, ParseStatus)
	Name() string
	Add(keys []string, actions []action.Action) error
	Remove(keys []string)
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
