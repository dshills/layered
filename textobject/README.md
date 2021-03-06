# textobject
--
    import "."


## Usage

```go
const (
	BlockParen   = "blockparen"
	BlockAngle   = "blockangle"
	BlockSquare  = "blocksquare"
	BlockCurly   = "blockcurly"
	BolNotBlank  = "bolnotblank"
	BOL          = "bol"
	EOLNotBlank  = "eolnotblank"
	EOL          = "eol"
	Line         = "line"
	Paragraph    = "paragraph"
	Sentence     = "sentence"
	StringDouble = "stringdouble"
	StringSingle = "stringsingle"
	StringTick   = "stringtick"
	Tag          = "tag"
	WordExt      = "wordext"
	Word         = "word"
)
```
Default objects

#### type Factory

```go
type Factory func(*conf.Configuration) Objecter
```

Factory will return an Objectr

#### type Object

```go
type Object struct {
}
```

Object is a text object

#### func (*Object) FindAfter

```go
func (o *Object) FindAfter(s string, col int) [][]int
```
FindAfter will return results after col

#### func (*Object) FindAll

```go
func (o *Object) FindAll(s string) [][]int
```
FindAll returns text object matches

#### func (*Object) FindBefore

```go
func (o *Object) FindBefore(s string, col int) [][]int
```
FindBefore will return results before col

#### func (*Object) MultiLine

```go
func (o *Object) MultiLine() bool
```
MultiLine returns true if a obj is mutliple line

#### func (*Object) Name

```go
func (o *Object) Name() string
```
Name returns the text object name

#### func (*Object) Simple

```go
func (o *Object) Simple() bool
```
Simple returns true if the object is simple

#### func (*Object) UseFirst

```go
func (o *Object) UseFirst() bool
```
UseFirst will return true if match expects the use of the first match

#### func (*Object) UseLast

```go
func (o *Object) UseLast() bool
```
UseLast will return true if match expects the use of the last match

#### type Objecter

```go
type Objecter interface {
	LoadDir(path string) error
	Object(name string) (TextObjecter, error)
	Add(...TextObjecter)
	Remove(name string)
}
```

Objecter is a set of text objects

#### func  New

```go
func New(config *conf.Configuration) Objecter
```
New returns a text object collection

#### type Objects

```go
type Objects struct {
}
```

Objects is a collection of objects

#### func (*Objects) Add

```go
func (o *Objects) Add(objs ...TextObjecter)
```
Add will add an object to the collection

#### func (*Objects) LoadDir

```go
func (o *Objects) LoadDir(dir string) error
```
LoadDir will load a collection of text objects

#### func (*Objects) Object

```go
func (o *Objects) Object(name string) (TextObjecter, error)
```
Object returns an object by name

#### func (*Objects) Remove

```go
func (o *Objects) Remove(name string)
```
Remove will remove an object from the collection

#### func (*Objects) Search

```go
func (o *Objects) Search(txt textstore.TextStorer, cur cursor.Cursor, oname string, cnt int) ([]int, error)
```
Search will search a text store

#### type TextObjecter

```go
type TextObjecter interface {
	Name() string
	FindAll(string) [][]int
	FindAfter(string, int) [][]int
	FindBefore(string, int) [][]int
	UseFirst() bool
	UseLast() bool
	MultiLine() bool
	Simple() bool
}
```

TextObjecter is a text object interface
