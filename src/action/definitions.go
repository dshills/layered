package action

// Movement
const (
	Move        = "move"
	MoveEnd     = "moveend"
	MovePrev    = "moveprev"
	MovePrevEnd = "moveprevend"
	Up          = "up"
	Down        = "down"
	Prev        = "prev"
	Next        = "next"
	ScrollDown  = "scrolldown"
	ScrollUp    = "scrollup"
)

// Action constants
const (
	Delete       = "delete"
	Deleteto     = "deleteto"
	Indent       = "indent"
	Insert       = "insert"
	Outdent      = "outdent"
	Paste        = "paste"
	Redo         = "redo"
	RunMacro     = "runmacro"
	RunCommand   = "runcommand"
	SetMark      = "setmark"
	RecordMacro  = "recordmacro"
	StopRecMacro = "stoprecmacro"
	Undo         = "undo"
	Yank         = "yank"
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
