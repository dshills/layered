# action
--
    import "."


## Usage

```go
const (
	Down        = "down"
	Move        = "move"
	MoveEnd     = "moveend"
	MovePrev    = "moveprev"
	MovePrevEnd = "moveprevend"
	Next        = "next"
	Prev        = "prev"
	ScrollDown  = "scrolldown"
	ScrollUp    = "scrollup"
	Up          = "up"
)
```
Movement

```go
const (
	DeleteChar      = "deletechar"
	DeleteCharBack  = "deletecharback"
	DeleteLine      = "deleteline"
	DeleteObject    = "deleteobject"
	Indent          = "indent"
	InsertLine      = "insertline"
	InsertLineAbove = "insertlineabove"
	InsertString    = "insertstring"
	Outdent         = "outdent"
	Content         = "content"
	Insert          = "insert"
)
```
Edit

```go
const (
	NewBuffer   = "newbuffer"
	SaveBuffer  = "savebuffer"
	CloseBuffer = "closebuffer"
	OpenFile    = "openfile"
	RenameFile  = "renamefile"
	SaveFileAs  = "savefileas"
	BufferList  = "bufferlist"
)
```
IO

```go
const (
	Search        = "search"
	SearchResults = "searchresults"
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
	RecordMacro  = "recordmacro"
	StopRecMacro = "stoprecmacro"
)
```
Macro

```go
const (
	Syntax     = "syntax"
	RunMacro   = "runmacro"
	RunCommand = "runcommand"
	SetMark    = "setmark"
)
```
Other

```go
var Definitions = []Def{
	Def{Name: Down, ReqBuffer: true},
	Def{Name: Move, ReqBuffer: true},
	Def{Name: MoveEnd, ReqBuffer: true},
	Def{Name: MovePrev, ReqBuffer: true},
	Def{Name: MovePrevEnd, ReqBuffer: true},
	Def{Name: Next, ReqBuffer: true},
	Def{Name: Prev, ReqBuffer: true},
	Def{Name: ScrollDown, ReqBuffer: true},
	Def{Name: ScrollUp, ReqBuffer: true},
	Def{Name: Up, ReqBuffer: true},
	Def{Name: DeleteChar, ReqBuffer: true},
	Def{Name: DeleteCharBack, ReqBuffer: true},
	Def{Name: DeleteLine, ReqBuffer: true},
	Def{Name: DeleteObject, ReqBuffer: true},
	Def{Name: Indent, ReqBuffer: true},
	Def{Name: InsertLine},
	Def{Name: InsertLineAbove},
	Def{Name: InsertString, ReqParam: true},
	Def{Name: Outdent, ReqBuffer: true},
	Def{Name: Content, ReqBuffer: true},
	Def{Name: Insert, ReqBuffer: true},
	Def{Name: NewBuffer},
	Def{Name: SaveBuffer},
	Def{Name: CloseBuffer, ReqBuffer: true},
	Def{Name: OpenFile, ReqParam: true},
	Def{Name: RenameFile, ReqParam: true},
	Def{Name: SaveFileAs, ReqParam: true},
	Def{Name: BufferList},
	Def{Name: Search, ReqBuffer: true, ReqParam: true},
	Def{Name: SearchResults, ReqBuffer: true},
	Def{Name: Yank, ReqBuffer: true},
	Def{Name: Paste, ReqBuffer: true},
	Def{Name: Redo, ReqBuffer: true},
	Def{Name: Undo, ReqBuffer: true},
	Def{Name: StopGroupUndo, ReqBuffer: true},
	Def{Name: StartGroupUndo, ReqBuffer: true},
	Def{Name: RecordMacro, ReqBuffer: true},
	Def{Name: StopRecMacro, ReqBuffer: true},
	Def{Name: Syntax, ReqBuffer: true},
	Def{Name: RunMacro, ReqBuffer: true},
	Def{Name: RunCommand, ReqBuffer: true},
	Def{Name: SetMark, ReqBuffer: true},
}
```
Definitions is a list of action definitions

#### type Action

```go
type Action struct {
}
```

Action is an editor action

#### func (*Action) Column

```go
func (a *Action) Column() int
```
Column returns the column -1 represents not set

#### func (*Action) Count

```go
func (a *Action) Count() int
```
Count will return the action count

#### func (*Action) Line

```go
func (a *Action) Line() int
```
Line will return the line -1 represents not set

#### func (*Action) Name

```go
func (a *Action) Name() string
```
Name returns the action name

#### func (*Action) NeedBuffer

```go
func (a *Action) NeedBuffer() bool
```
NeedBuffer will return true if the action requires a buffer

#### func (*Action) Object

```go
func (a *Action) Object() string
```
Object returns the object associated with the action

#### func (*Action) Param

```go
func (a *Action) Param() string
```
Param is a required parameter

#### func (*Action) SetColumn

```go
func (a *Action) SetColumn(c int)
```
SetColumn will set the column associated with the action

#### func (*Action) SetCount

```go
func (a *Action) SetCount(c int)
```
SetCount will set the action count

#### func (*Action) SetLine

```go
func (a *Action) SetLine(l int)
```
SetLine will set the line associated with the action

#### func (*Action) SetObject

```go
func (a *Action) SetObject(obj string)
```
SetObject will set the object for the action

#### func (*Action) SetParam

```go
func (a *Action) SetParam(p string)
```
SetParam will set the param

#### func (*Action) SetTarget

```go
func (a *Action) SetTarget(t string)
```
SetTarget will set the target

#### func (*Action) Target

```go
func (a *Action) Target() string
```
Target is the target of the action

#### func (*Action) Valid

```go
func (a *Action) Valid(bufid string) error
```
Valid will return true if it is a valid action

#### type Actioner

```go
type Actioner interface {
	Name() string
	Target() string
	SetTarget(string)
	Param() string
	SetParam(string)
	Column() int
	SetColumn(int)
	Line() int
	SetLine(int)
	Object() string
	SetObject(string)
	Count() int
	SetCount(int)
	Valid(bufid string) error
	NeedBuffer() bool
}
```

Actioner represents an editor action

#### func  New

```go
func New(act string) Actioner
```
New will return a new Actioner

#### type Def

```go
type Def struct {
	Name      string
	ReqBuffer bool
	ReqParam  bool
	ReqTarget bool
}
```

Def is a definition for an action

#### type Transaction

```go
type Transaction struct {
}
```

Transaction is a group of actions with a buffer identifier

#### func (*Transaction) Actions

```go
func (t *Transaction) Actions() []Actioner
```
Actions returns the set of actions

#### func (*Transaction) Add

```go
func (t *Transaction) Add(acts ...Actioner)
```
Add will add actions to the transactions

#### func (*Transaction) Buffer

```go
func (t *Transaction) Buffer() string
```
Buffer will return the buffer id

#### func (*Transaction) NeedBuffer

```go
func (t *Transaction) NeedBuffer() bool
```
NeedBuffer will return true if the action list requires a buffer

#### func (*Transaction) Set

```go
func (t *Transaction) Set(acts ...Actioner)
```
Set will set actions to the transactions

#### func (*Transaction) SetBuffer

```go
func (t *Transaction) SetBuffer(id string)
```
SetBuffer will set the transaction buffer

#### func (*Transaction) Valid

```go
func (t *Transaction) Valid() error
```
Valid will return true if a valid transaction

#### type Transactioner

```go
type Transactioner interface {
	Buffer() string
	SetBuffer(string)
	Actions() []Actioner
	Add(acts ...Actioner)
	Set(acts ...Actioner)
	Valid() error
	NeedBuffer() bool
}
```

Transactioner is group of actions for with a a buffer id

#### func  NewTransaction

```go
func NewTransaction(id string, actions ...Actioner) Transactioner
```
NewTransaction will return an we transactioner
