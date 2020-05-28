package action

// Movement
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

// Edit
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

// IO
const (
	NewBuffer   = "newbuffer"
	SaveBuffer  = "savebuffer"
	CloseBuffer = "closebuffer"
	OpenFile    = "openfile"
	RenameFile  = "renamefile"
	SaveFileAs  = "savefileas"
	BufferList  = "bufferlist"
)

// Search
const (
	Search        = "search"
	SearchResults = "searchresults"
)

// Yank
const (
	Yank  = "yank"
	Paste = "paste"
)

// Undo
const (
	Redo = "redo"
	Undo = "undo"
)

// Macro
const (
	RecordMacro  = "recordmacro"
	StopRecMacro = "stoprecmacro"
)

// Other
const (
	Syntax     = "syntax"
	RunMacro   = "runmacro"
	RunCommand = "runcommand"
	SetMark    = "setmark"
)

// Def is a definition for an action
type Def struct {
	Name     string
	NoBuffer bool
	NoParam  bool
}

// Definitions is a list of action definitions
var Definitions = []Def{
	Def{Name: Down, NoParam: true},
	Def{Name: Move, NoParam: true},
	Def{Name: MoveEnd, NoParam: true},
	Def{Name: MovePrev, NoParam: true},
	Def{Name: MovePrevEnd, NoParam: true},
	Def{Name: Next, NoParam: true},
	Def{Name: Prev, NoParam: true},
	Def{Name: ScrollDown, NoParam: true},
	Def{Name: ScrollUp, NoParam: true},
	Def{Name: Up, NoParam: true},
	Def{Name: DeleteChar, NoParam: true},
	Def{Name: DeleteCharBack, NoParam: true},
	Def{Name: DeleteLine, NoParam: true},
	Def{Name: DeleteObject, NoParam: true},
	Def{Name: Indent, NoParam: true},
	Def{Name: InsertLine},
	Def{Name: InsertLineAbove},
	Def{Name: InsertString},
	Def{Name: Outdent, NoParam: true},
	Def{Name: Content, NoParam: true},
	Def{Name: Insert, NoParam: true},
	Def{Name: NewBuffer, NoBuffer: true, NoParam: true},
	Def{Name: SaveBuffer},
	Def{Name: CloseBuffer, NoParam: true},
	Def{Name: OpenFile, NoBuffer: true},
	Def{Name: RenameFile},
	Def{Name: SaveFileAs},
	Def{Name: BufferList, NoBuffer: true, NoParam: true},
	Def{Name: Search},
	Def{Name: SearchResults, NoParam: true},
	Def{Name: Yank, NoParam: true},
	Def{Name: Paste, NoParam: true},
	Def{Name: Redo, NoParam: true},
	Def{Name: Undo, NoParam: true},
	Def{Name: RecordMacro, NoParam: true},
	Def{Name: StopRecMacro, NoParam: true},
	Def{Name: Syntax, NoParam: true},
	Def{Name: RunMacro, NoParam: true},
	Def{Name: RunCommand, NoParam: true},
	Def{Name: SetMark, NoParam: true},
}
