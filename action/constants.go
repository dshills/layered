package action

// Movement
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

// Edit
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

// IO
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

// Search
const (
	Search        = "search"
	SearchResults = "searchresults"
	DeleteCmdBack = "deletecmdback" // deletes a character from a partial command
)

// Yank
const (
	Yank  = "yank"
	Paste = "paste"
)

// Undo
const (
	Redo           = "redo"
	Undo           = "undo"
	StartGroupUndo = "startgroupundo"
	StopGroupUndo  = "stopgroupundo"
)

// Macro
const (
	StartRecordMacro = "startrecordmacro"
	StopRecordMacro  = "stoprecmacro"
	RunMacro         = "runmacro"
)

// Other
const (
	Syntax          = "syntax"
	RunCommand      = "runcommand"
	SetMark         = "setmark"
	ChangeLayer     = "changelayer"
	ChangePrevLayer = "changeprevlayer"
	TypeHighlight   = "typehighlight"
	Quit            = "quit"
)
