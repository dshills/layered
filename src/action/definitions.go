package action

import (
	"fmt"
	"strings"
)

// Definitions is a list of action definitions
var Definitions = []Def{
	Def{Name: BufferList, Alias: []string{"ls"}},
	Def{Name: ChangeLayer},
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
	Def{Name: Insert, ReqBuffer: true},
	Def{Name: InsertLineAbove},
	Def{Name: InsertLine},
	Def{Name: InsertString, ReqParam: true},
	Def{Name: Move, ReqBuffer: true},
	Def{Name: MoveEnd, ReqBuffer: true},
	Def{Name: MovePrev, ReqBuffer: true},
	Def{Name: MovePrevEnd, ReqBuffer: true},
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
	return Action{}, fmt.Errorf("Action %v Not found", s)
}

// Def is a definition for an action
type Def struct {
	Name       string
	Alias      []string
	ReqBuffer  bool
	ReqParam   bool
	ReqTarget  bool
	Targets    []string
	IsMovement bool
}
