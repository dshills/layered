# syntax
--
    import "."


## Usage

```go
const (
	MatchRule   = "match"
	RegionRule  = "region"
	KeywordRule = "keyword"
)
```
Rule types

#### type Factory

```go
type Factory func(rt ...string) Manager
```

Factory is a function that returns new syntax matchers

#### type Manager

```go
type Manager interface {
	LoadFileType(ft string) error
	LoadFile(path string) error
	Add(Ruler)
	Parse(textstore.TextStorer) []Resulter
}
```

Manager is a collection of syntax rules representing a set of language syntax
rules

#### func  New

```go
func New(rt ...string) Manager
```
New returns a new syntax matcher

#### type Matcher

```go
type Matcher struct {
}
```

Matcher is syntax matcher

#### func (*Matcher) Add

```go
func (m *Matcher) Add(r Ruler)
```
Add will add a rule to the matcher

#### func (*Matcher) LoadFile

```go
func (m *Matcher) LoadFile(path string) error
```
LoadFile will load a syntax file

#### func (*Matcher) LoadFileType

```go
func (m *Matcher) LoadFileType(ft string) error
```
LoadFileType will load a syntax file by file type

#### func (*Matcher) Parse

```go
func (m *Matcher) Parse(ts textstore.TextStorer) []Resulter
```
Parse will return a list of results for the text store

#### type Result

```go
type Result struct {
}
```

Result is a a syntax match result

#### func (*Result) AddRanges

```go
func (r *Result) AddRanges(rg [][]int)
```
AddRanges will append result ranges

#### func (*Result) IsEqual

```go
func (r *Result) IsEqual(Resulter) bool
```
IsEqual compares one result to another returning true if equal

#### func (*Result) Line

```go
func (r *Result) Line() int
```
Line returns the line for the result

#### func (*Result) Priority

```go
func (r *Result) Priority() int
```
Priority returns the results priority

#### func (*Result) Range

```go
func (r *Result) Range() [][]int
```
Range returns the range if matches

#### func (*Result) SetLine

```go
func (r *Result) SetLine(ln int)
```
SetLine will set the result line

#### func (*Result) SetPriority

```go
func (r *Result) SetPriority(p int)
```
SetPriority will set the priority of the result

#### func (*Result) SetRanges

```go
func (r *Result) SetRanges(rg [][]int)
```
SetRanges will set the result ranges

#### func (*Result) SetToken

```go
func (r *Result) SetToken(tok string)
```
SetToken will set the result token

#### func (*Result) Token

```go
func (r *Result) Token() string
```
Token will return the rules token type

#### type Resulter

```go
type Resulter interface {
	IsEqual(Resulter) bool
	Token() string
	Line() int
	Range() [][]int
	Priority() int
	SetToken(string)
	SetLine(int)
	SetRanges([][]int)
	SetPriority(int)
	AddRanges([][]int)
}
```

Resulter is a a syntax match result

#### type Rule

```go
type Rule struct {
}
```

Rule is a syntax matching rules

#### func (*Rule) Group

```go
func (r *Rule) Group() string
```
Group will return the rules group

#### func (*Rule) IsDependent

```go
func (r *Rule) IsDependent() bool
```
IsDependent will return true if the rule has no dependency on other rules

#### func (*Rule) Match

```go
func (r *Rule) Match(txt textstore.TextStorer) []Resulter
```
Match will return match results

#### func (*Rule) Type

```go
func (r *Rule) Type() string
```
Type returns the rule type

#### type Ruler

```go
type Ruler interface {
	Group() string
	Type() string
	Match(textstore.TextStorer) []Resulter
	IsDependent() bool
}
```

Ruler is a syntax matching rule
