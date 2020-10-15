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

#### func (*Editor) ActionChan

```go
func (e *Editor) ActionChan() chan action.Request
```
ActionChan returns the action channel

#### func (*Editor) ActionDefs

```go
func (e *Editor) ActionDefs() action.Definitions
```
ActionDefs returns the action definitions

#### func (*Editor) Buffers

```go
func (e *Editor) Buffers() []buffer.Bufferer
```
Buffers returns the editors currrent buffers

#### func (*Editor) DoneChan

```go
func (e *Editor) DoneChan() chan struct{}
```
DoneChan returns the done channel

#### func (*Editor) ExecChan

```go
func (e *Editor) ExecChan(reqC chan action.Request, respC chan Response, done chan struct{})
```
ExecChan will listen for requests

#### type Editorer

```go
type Editorer interface {
	ExecChan(reqC chan action.Request, respC chan Response, done chan struct{})
}
```

Editorer is an editor interface

#### func  New

```go
func New(uf undo.Factory, tf textstore.Factory, bf buffer.Factory, cf cursor.Factory, sf syntax.Factory, ftf filetype.Factory, of textobject.Factory, rf register.Factory, ad action.Definitions, config *conf.Configuration) (Editorer, error)
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
	BufferID       string
	Action         action.Action
	Line, Column   int
	Dirty          bool
	Filename       string
	Filetype       string
	NumLines       int
	Results        []KeyValue
	Content        []string
	Syntax         []syntax.Resulter
	Search         []buffer.SearchResult
	ContentChanged bool
	CursorChanged  bool
	NewBuffer      bool
	CloseBuffer    bool
	InfoChanged    bool
	Quit           bool
	Err            error
}
```

Response is a exec response
