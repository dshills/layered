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

```go
var Definitions = []Def{
	Def{Name: BufferList, Alias: []string{"ls"}},
	Def{Name: ChangeLayer},
	Def{Name: ChangePrevLayer},
	Def{Name: CloseBuffer, ReqBuffer: true},
	Def{Name: Content, ReqBuffer: true},
	Def{Name: Delete, ReqBuffer: true},
	Def{Name: DeleteChar, ReqBuffer: true},
	Def{Name: DeleteCharBack, ReqBuffer: true},
	Def{Name: DeleteCmdBack},
	Def{Name: DeleteLine, ReqBuffer: true},
	Def{Name: DeleteObject, ReqBuffer: true},
	Def{Name: DeleteToObject, ReqBuffer: true},
	Def{Name: Down, ReqBuffer: true},
	Def{Name: Indent, ReqBuffer: true},
	Def{Name: InsertLineAbove},
	Def{Name: InsertLine},
	Def{Name: InsertString, ReqBuffer: true, ReqParam: true},
	Def{Name: Move, ReqBuffer: true},
	Def{Name: MoveObj, ReqBuffer: true},
	Def{Name: MoveEnd, ReqBuffer: true},
	Def{Name: MovePrev, ReqBuffer: true},
	Def{Name: MovePrevEnd, ReqBuffer: true},
	Def{Name: CursorMovePast, ReqBuffer: true},
	Def{Name: NewBuffer},
	Def{Name: Next, ReqBuffer: true},
	Def{Name: OpenFile, Alias: []string{"e", "edit"}, ReqParam: true},
	Def{Name: Outdent, ReqBuffer: true},
	Def{Name: Paste, ReqBuffer: true},
	Def{Name: Prev, ReqBuffer: true},
	Def{Name: Quit, Alias: []string{"q"}},
	Def{Name: Redo, ReqBuffer: true},
	Def{Name: RenameFile, ReqParam: true},
	Def{Name: RunCommand, ReqBuffer: true},
	Def{Name: RunMacro, ReqBuffer: true},
	Def{Name: SaveBuffer},
	Def{Name: SaveFileAs, Alias: []string{"w", "write"}, ReqParam: true},
	Def{Name: ScrollDown, ReqBuffer: true},
	Def{Name: ScrollUp, ReqBuffer: true},
	Def{Name: Search, ReqBuffer: true, ReqParam: true},
	Def{Name: SearchResults, ReqBuffer: true},
	Def{Name: SelectBuffer, ReqBuffer: true},
	Def{Name: SetMark, ReqBuffer: true},
	Def{Name: StartGroupUndo, ReqBuffer: true},
	Def{Name: StartRecordMacro, ReqBuffer: true},
	Def{Name: StartSelection, ReqBuffer: true},
	Def{Name: StopGroupUndo, ReqBuffer: true},
	Def{Name: StopRecordMacro, ReqBuffer: true},
	Def{Name: StopSelection, ReqBuffer: true},
	Def{Name: Syntax, ReqBuffer: true},
	Def{Name: Undo, ReqBuffer: true},
	Def{Name: Up, ReqBuffer: true},
	Def{Name: Yank, ReqBuffer: true},
}
```
Definitions is a list of action definitions

#### type Action

```go
type Action struct {
	Name, Target string
	Line, Column int
	Count        int
}
```

Action is an editor action

#### func  StrToAction

```go
func StrToAction(s string) (Action, error)
```
StrToAction will convert a string to an action it will return an error if the
action is not found

#### func (*Action) NeedBuffer

```go
func (a *Action) NeedBuffer() bool
```
NeedBuffer will return true if the action requires a buffer

#### func (Action) String

```go
func (a Action) String() string
```

#### func (*Action) Valid

```go
func (a *Action) Valid(bufid string) error
```
Valid will return true if it is a valid action

#### type Def

```go
type Def struct {
	Name       string
	Alias      []string
	ReqBuffer  bool
	ReqParam   bool
	ReqTarget  bool
	Targets    []string
	IsMovement bool
}
```

Def is a definition for an action