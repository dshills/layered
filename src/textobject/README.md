# textobject
--
    import "."


## Usage

#### type Factory

```go
type Factory func(rts ...string) Objecter
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

#### func (*Object) Name

```go
func (o *Object) Name() string
```
Name returns the text object name

#### type Objecter

```go
type Objecter interface {
	SetRuntimes(rts ...string)
	AddRuntimes(rts ...string)
	LoadDir(path string) error
	Object(name string) (TextObjecter, error)
	Add(...TextObjecter)
	Remove(name string)
}
```

Objecter is a set of text objects

#### func  New

```go
func New(rts ...string) Objecter
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

#### func (*Objects) AddRuntimes

```go
func (o *Objects) AddRuntimes(rts ...string)
```
AddRuntimes will add to the list of runtimes

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

#### func (*Objects) SetRuntimes

```go
func (o *Objects) SetRuntimes(rts ...string)
```
SetRuntimes will set the list of runtime directories

#### type TextObjecter

```go
type TextObjecter interface {
	Name() string
	FindAll(string) [][]int
	FindAfter(string, int) [][]int
	FindBefore(string, int) [][]int
}
```

TextObjecter is a text object interface
