# action
--
    import "."


## Usage

```go
const (
	Down           = "down"
	Move           = "move"
	MoveObj        = "moveobj"
	MoveEnd        = "moveend"
	MovePrev       = "moveprev"
	MovePrevEnd    = "moveprevend"
	Next           = "next"
	Prev           = "prev"
	ScrollDown     = "scrolldown"
	ScrollUp       = "scrollup"
	Up             = "up"
	StartSelection = "startselection"
	StopSelection  = "stopselection"
	CursorMovePast = "cursormovepast"
)
```
Movement

```go
const (
	Delete          = "delete"
	DeleteChar      = "deletechar"
	DeleteCharBack  = "deletecharback"
	DeleteLine      = "deleteline"
	DeleteObject    = "deleteobject"
	DeleteToObject  = "deletetoobject"
	Indent          = "indent"
	InsertLine      = "insertline"
	InsertLineAbove = "insertlineabove"
	InsertString    = "insertstring"
	Outdent         = "outdent"
	Content         = "content"
)
```
Edit

```go
const (
	NewBuffer    = "newbuffer"
	SaveBuffer   = "savebuffer"
	CloseBuffer  = "closebuffer"
	OpenFile     = "openfile"
	RenameFile   = "renamefile"
	SaveFileAs   = "savefileas"
	BufferList   = "bufferlist"
	SelectBuffer = "selectbuffer"
)
```
IO

```go
const (
	Search        = "search"
	SearchResults = "searchresults"
	DeleteCmdBack = "deletecmdback" // deletes a character from a partial command
)
```
Search

```go
const (
	Yank  = "yank"
	Paste = "paste"
)
```
Yank

```go
const (
	Redo           = "redo"
	Undo           = "undo"
	StartGroupUndo = "startgroupundo"
	StopGroupUndo  = "stopgroupundo"
)
```
Undo

```go
const (
	StartRecordMacro = "startrecordmacro"
	StopRecordMacro  = "stoprecmacro"
	RunMacro         = "runmacro"
)
```
Macro

```go
const (
	Syntax          = "syntax"
	RunCommand      = "runcommand"
	SetMark         = "setmark"
	ChangeLayer     = "changelayer"
	ChangePrevLayer = "changeprevlayer"
	Quit            = "quit"
)
```
Other

#### type Action

```go
type Action struct {
	Name   string `json:"name"`
	Target string `json:"target"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
	Count  int    `json:"count"`
}
```

Action is an editor action

#### func  JSONtoRequest

```go
func JSONtoRequest(js []byte) (Action, error)
```
JSONtoRequest will convert a json encoded request to a Request struct

#### func  ReaderToRequest

```go
func ReaderToRequest(r io.Reader) (Action, error)
```
ReaderToRequest will convert a json stream to a Request

#### func (Action) String

```go
func (a Action) String() string
```

#### type Def

```go
type Def struct {
	Name        string
	NoReqBuffer bool
	ReqTarget   bool
	IsMovement  bool
	ReqCount    bool
	ReqLine     bool
	ReqColumn   bool
}
```

Def is a definition for an action

#### type Definitions

```go
type Definitions struct {
}
```

Definitions is a list of action definitions

#### func  NewDefinitions

```go
func NewDefinitions() *Definitions
```
NewDefinitions will return the action definitions

#### func (*Definitions) Add

```go
func (dl *Definitions) Add(dd ...Def)
```
Add will add definitions

#### func (*Definitions) Get

```go
func (dl *Definitions) Get(n string) *Def
```
Get will return a definition by name or nil if not found

#### func (*Definitions) RequireBuffer

```go
func (dl *Definitions) RequireBuffer(n string) bool
```
RequireBuffer will return true if the action requires a buffer

#### func (*Definitions) RequireTarget

```go
func (dl *Definitions) RequireTarget(n string) bool
```
RequireTarget will return true if the action requires a target

#### func (*Definitions) StrToAction

```go
func (dl *Definitions) StrToAction(n string) (Action, error)
```
StrToAction will convert a string to an action it will return an error if the
action is not found

#### func (*Definitions) ValidAction

```go
func (dl *Definitions) ValidAction(act Action, bufid string) error
```
ValidAction will return an error for an invalid action nil otherwise

#### func (*Definitions) ValidRequest

```go
func (dl *Definitions) ValidRequest(req Request) error
```
ValidRequest will return an error for an invalid request nil otherwise

#### type Request

```go
type Request struct {
	BufferID   string   `json:"buffer_id"`
	LineOffset int      `json:"line_offset"`
	LineCount  int      `json:"line_count"`
	Actions    []Action `json:"actions"`
}
```

Request is a request for actions

#### func  NewRequest

```go
func NewRequest(bufid string, acts ...Action) Request
```
NewRequest returns a Request

#### func (*Request) Add

```go
func (r *Request) Add(act ...Action)
```
Add will add actions to a request

#### func (*Request) AsJSON

```go
func (r *Request) AsJSON() []byte
```
AsJSON will return a json encoded request

#### func (*Request) AsJSONReader

```go
func (r *Request) AsJSONReader() io.Reader
```
AsJSONReader returns a json encoded request, io.Reader
