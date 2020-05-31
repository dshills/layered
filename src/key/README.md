# key
--
    import "."


## Usage

```go
const (
	Alta          = rune(229)
	Altb          = rune(8747)
	Altc          = rune(231)
	Altd          = rune(8706)
	AltEe         = rune(180)
	Altf          = rune(402)
	Altg          = rune(169)
	Alth          = rune(729)
	AltIi         = rune(710)
	Altj          = rune(8710)
	Altk          = rune(730)
	Altl          = rune(172)
	Altm          = rune(181)
	AltNn         = rune(732)
	Alto          = rune(248)
	Altp          = rune(960)
	Altq          = rune(339)
	Altr          = rune(174)
	Alts          = rune(223)
	Altt          = rune(8224)
	AltUu         = rune(168)
	Altv          = rune(8730)
	Altw          = rune(8721)
	Altx          = rune(8776)
	Alty          = rune(92)
	Altz          = rune(937)
	AltA          = rune(197)
	AltB          = rune(305)
	AltC          = rune(199)
	AltD          = rune(206)
	AltF          = rune(207)
	AltG          = rune(733)
	AltH          = rune(211)
	AltJ          = rune(212)
	AltK          = rune(63743)
	AltL          = rune(210)
	AltM          = rune(194)
	AltO          = rune(216)
	AltP          = rune(8719)
	AltQ          = rune(338)
	AltR          = rune(8240)
	AltS          = rune(205)
	AltT          = rune(711)
	AltV          = rune(9674)
	AltW          = rune(8222)
	AltX          = rune(731)
	AltY          = rune(193)
	AltZ          = rune(184)
	Alt0          = rune(186)
	Alt1          = rune(161)
	Alt2          = rune(8482)
	Alt3          = rune(163)
	Alt4          = rune(162)
	Alt5          = rune(8734)
	Alt6          = rune(167)
	Alt7          = rune(182)
	Alt8          = rune(8226)
	Alt9          = rune(170)
	AltBang       = rune(8260)
	AltAt         = rune(8364)
	AltPound      = rune(8249)
	AltDollar     = rune(8250)
	AltPercent    = rune(64257)
	AltCarrot     = rune(64258)
	AltAnd        = rune(8225)
	AltStar       = rune(176)
	AltLeftParan  = rune(183)
	AltRightParan = rune(8218)
)
```
Alt rune codes

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
func (k *Key) IsEqual(o Keyer) bool
```
IsEqual returns true if same key

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

#### type Keyer

```go
type Keyer interface {
	Special() bool
	Alt() bool
	Ctrl() bool
	Rune() rune
	SpecialKey() string
	IsEqual(o Keyer) bool
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
