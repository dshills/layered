# key
--
    import "."


## Usage

```go
const (
	AltA          = "<alt-a>"
	AltB          = "<alt-b>"
	AltC          = "<alt-c>"
	AltD          = "<alt-d>"
	AltF          = "<alt-f>"
	AltG          = "<alt-g>"
	AltH          = "<alt-h>"
	AltJ          = "<alt-j>"
	AltK          = "<alt-k>"
	AltL          = "<alt-l>"
	AltM          = "<alt-m>"
	AltO          = "<alt-o>"
	AltP          = "<alt-p>"
	AltQ          = "<alt-q>"
	AltR          = "<alt-r>"
	AltS          = "<alt-s>"
	AltT          = "<alt-t>"
	AltV          = "<alt-v>"
	AltW          = "<alt-w>"
	AltX          = "<alt-x>"
	AltY          = "<alt-y>"
	AltZ          = "<alt-z>"
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
	RunAltEe         = rune(180)
	RunAltf          = rune(402)
	RunAltg          = rune(169)
	RunAlth          = rune(729)
	RunAltIi         = rune(710)
	RunAltj          = rune(8710)
	RunAltk          = rune(730)
	RunAltl          = rune(172)
	RunAltm          = rune(181)
	RunAltNn         = rune(732)
	RunAlto          = rune(248)
	RunAltp          = rune(960)
	RunAltq          = rune(339)
	RunAltr          = rune(174)
	RunAlts          = rune(223)
	RunAltt          = rune(8224)
	RunAltUu         = rune(168)
	RunAltv          = rune(8730)
	RunAltw          = rune(8721)
	RunAltx          = rune(8776)
	RunAlty          = rune(92)
	RunAltz          = rune(937)
	RunAltA          = rune(197)
	RunAltB          = rune(305)
	RunAltC          = rune(199)
	RunAltD          = rune(206)
	RunAltF          = rune(207)
	RunAltG          = rune(733)
	RunAltH          = rune(211)
	RunAltJ          = rune(212)
	RunAltK          = rune(63743)
	RunAltL          = rune(210)
	RunAltM          = rune(194)
	RunAltO          = rune(216)
	RunAltP          = rune(8719)
	RunAltQ          = rune(338)
	RunAltR          = rune(8240)
	RunAltS          = rune(205)
	RunAltT          = rune(711)
	RunAltV          = rune(9674)
	RunAltW          = rune(8222)
	RunAltX          = rune(731)
	RunAltY          = rune(193)
	RunAltZ          = rune(184)
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
	Nul       = "<nul>"
	Soh       = "<soh>"
	Stx       = "<stx>"
	Etx       = "<etx>"
	Eot       = "<eot>"
	Enq       = "<enq>"
	Ack       = "<ack>"
	Bel       = "<bel>"
	BS        = "<bs>"
	Tab       = "<tab>"
	LF        = "<lf>"
	VT        = "<vt>"
	FF        = "<ff>"
	CR        = "<cr>"
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

#### type Key

```go
type Key struct {
}
```

Key is a keyboard key press

#### func (*Key) Alt

```go
func (k *Key) Alt() bool
```
Alt returns true if an alt key press

#### func (*Key) Ctrl

```go
func (k *Key) Ctrl() bool
```
Ctrl returns true if an ctrl key press

#### func (*Key) IsEqual

```go
func (k *Key) IsEqual(keys ...Keyer) bool
```
IsEqual returns true if key(s) match

#### func (*Key) IsMatchMultiple

```go
func (k *Key) IsMatchMultiple() bool
```
IsMatchMultiple returns true if the key pattern matches multiple keys

#### func (*Key) Matches

```go
func (k *Key) Matches(keys ...Keyer) int
```
Matches returns the number of key matches from 0

#### func (*Key) Rune

```go
func (k *Key) Rune() rune
```
Rune returns the key rune

#### func (*Key) Special

```go
func (k *Key) Special() bool
```
Special returns true if a special key press

#### func (*Key) SpecialKey

```go
func (k *Key) SpecialKey() string
```
SpecialKey returns true if a special key press

#### func (*Key) String

```go
func (k *Key) String() string
```

#### type Keyer

```go
type Keyer interface {
	Special() bool
	Alt() bool
	Ctrl() bool
	Rune() rune
	SpecialKey() string
	IsEqual(keys ...Keyer) bool
	Matches(keys ...Keyer) int
	IsMatchMultiple() bool
}
```

Keyer represents a keyboard item

#### func  NewKey

```go
func NewKey(a, c, s bool, r rune, sp string) Keyer
```
NewKey will return a key

#### func  StrToKey

```go
func StrToKey(s string) Keyer
```
StrToKey converts a string representaion to a Keyer
