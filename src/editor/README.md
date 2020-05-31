# editor
--
    import "."


## Usage

#### type Editor

```go
type Editor struct {
}
```

Editor is an editor instance

#### func (*Editor) Add

```go
func (e *Editor) Add(buf buffer.Bufferer)
```
Add will add a buffer to the editor

#### func (*Editor) Buffer

```go
func (e *Editor) Buffer(id string) (buffer.Bufferer, error)
```
Buffer will return a buffer by id

#### func (*Editor) Buffers

```go
func (e *Editor) Buffers() []buffer.Bufferer
```
Buffers returns the editors currrent buffers

#### func (*Editor) Exec

```go
func (e *Editor) Exec(bufid string, actions ...action.Action) (resp *Response, err error)
```
Exec will execute a transaction in the editor

#### func (*Editor) Remove

```go
func (e *Editor) Remove(id string) error
```
Remove will remove a buffer from the editor

#### type Editorer

```go
type Editorer interface {
	Buffers() []buffer.Bufferer
	Add(buffer.Bufferer)
	Remove(id string) error
	Buffer(id string) (buffer.Bufferer, error)
	Exec(bufid string, actions ...action.Action) (*Response, error)
}
```

Editorer represents an editor

#### func  New

```go
func New(uf undo.Factory, tf textstore.Factory, bf buffer.Factory, cf cursor.Factory, sf syntax.Factory, ftf filetype.Factory, of textobject.Factory, rf register.Factory, rt ...string) (Editorer, error)
```
New will return a new editor

#### type KeyValue

```go
type KeyValue struct {
	Key   string
	Value string
}
```

KeyValue is key/value data

#### type Response

```go
type Response struct {
	Buffer       string
	Action       string
	Line, Column int
	Results      []KeyValue
	Content      []string
	Syntax       []syntax.Resulter
	Search       []buffer.SearchResult
}
```

Response is a exec response
