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
func (e *Editor) ActionChan() chan []action.Action
```
ActionChan returns the action channel

#### func (*Editor) Buffers

```go
func (e *Editor) Buffers() []buffer.Bufferer
```
Buffers returns the editors currrent buffers

#### func (*Editor) Exec

```go
func (e *Editor) Exec(bufid string, actions ...action.Action) Response
```
Exec will execute a transaction in the editor

#### func (*Editor) KeyChan

```go
func (e *Editor) KeyChan() chan key.Keyer
```
KeyChan returns the key channel

#### func (*Editor) SetRespChan

```go
func (e *Editor) SetRespChan(rc chan Response)
```
SetRespChan will set the channel for sending responses

#### type Editorer

```go
type Editorer interface {
	Exec(bufid string, actions ...action.Action) Response
	KeyChan() chan key.Keyer
	ActionChan() chan []action.Action
	SetRespChan(chan Response)
}
```

Editorer is an editor interface

#### func  New

```go
func New(
	uf undo.Factory,
	tf textstore.Factory,
	bf buffer.Factory,
	cf cursor.Factory,
	sf syntax.Factory,
	ftf filetype.Factory,
	of textobject.Factory,
	rf register.Factory,
	lf layer.Factory,
	rt ...string,
) (Editorer, error)
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
	Buffer         string
	Action         string
	Line, Column   int
	Results        []KeyValue
	Content        []string
	Syntax         []syntax.Resulter
	Search         []buffer.SearchResult
	Layer          string
	Status         layer.ParseStatus
	Partial        string
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
