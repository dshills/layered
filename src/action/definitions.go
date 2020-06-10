package action

import (
	"fmt"
	"strings"
)

// StrToAction will convert a string to an action
// it will return an error if the action is not found
func StrToAction(s string) (Action, error) {
	s = strings.ToLower(s)
	for i := range Definitions {
		if Definitions[i].Name == s {
			return Action{Name: Definitions[i].Name}, nil
		}
		for _, al := range Definitions[i].Alias {
			if al == s {
				return Action{Name: Definitions[i].Name}, nil
			}
		}
	}
	return Action{}, fmt.Errorf("Not found")
}

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
	RecordMacro  = "recordmacro"
	StopRecMacro = "stoprecmacro"
)

// Other
const (
	Syntax      = "syntax"
	RunMacro    = "runmacro"
	RunCommand  = "runcommand"
	SetMark     = "setmark"
	ChangeLayer = "changelayer"
)

// Def is a definition for an action
type Def struct {
	Name      string
	Alias     []string
	ReqBuffer bool
	ReqParam  bool
	ReqTarget bool
}

// Definitions is a list of action definitions
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
	Def{Name: OpenFile, Alias: []string{"e", "edit"}, ReqParam: true},
	Def{Name: RenameFile, ReqParam: true},
	Def{Name: SaveFileAs, Alias: []string{"w", "write"}, ReqParam: true},
	Def{Name: BufferList, Alias: []string{"ls"}},
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
	Def{Name: ChangeLayer},
}
