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

// Other
const (
	Search        = "search"
	SearchResults = "searchresults"
	Syntax        = "syntax"
	Insert        = "insert"
	Redo          = "redo"
	Undo          = "undo"
	RunMacro      = "runmacro"
	RunCommand    = "runcommand"
	SetMark       = "setmark"
	RecordMacro   = "recordmacro"
	StopRecMacro  = "stoprecmacro"
	Yank          = "yank"
	Paste         = "paste"
)

// Def is a definition for an action
type Def struct {
	Name string
}

// Definitions is a list of action definitions
var Definitions = []Def{
	Def{Name: "delete"},
	Def{Name: "deleteto"},
	Def{Name: "indent"},
	Def{Name: "insert"},
	Def{Name: "move"},
	Def{Name: "moveend"},
	Def{Name: "moveprev"},
	Def{Name: "moveprevend"},
	Def{Name: "outdent"},
	Def{Name: "paste"},
	Def{Name: "redo"},
	Def{Name: "runmacro"},
	Def{Name: "runcommand"},
	Def{Name: "scroll"},
	Def{Name: "setmark"},
	Def{Name: "togglerecordmacro"},
	Def{Name: "undo"},
	Def{Name: "yank"},
}
