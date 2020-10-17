package action

import (
	"fmt"
	"strings"
	"sync"
)

// New return the action definitions
func New() Definitions {
	dl := deflist{list: make(map[string]Def)}
	dl.Add(
		Def{Name: BufferList, NoReqBuffer: true},
		Def{Name: ChangeLayer, ReqTarget: true},
		Def{Name: ChangePrevLayer},
		Def{Name: CloseBuffer},
		Def{Name: Content},
		Def{Name: Delete},
		Def{Name: DeleteChar},
		Def{Name: DeleteCharBack},
		Def{Name: DeleteCmdBack, NoReqBuffer: true},
		Def{Name: DeleteLine},
		Def{Name: DeleteObject},
		Def{Name: DeleteToObject},
		Def{Name: Down},
		Def{Name: Indent},
		Def{Name: InsertLineAbove},
		Def{Name: InsertLine},
		Def{Name: InsertString, ReqTarget: true},
		Def{Name: Move},
		Def{Name: MoveObj, ReqTarget: true},
		Def{Name: MoveEnd},
		Def{Name: MovePrev},
		Def{Name: MovePrevEnd},
		Def{Name: CursorMovePast},
		Def{Name: NewBuffer, NoReqBuffer: true},
		Def{Name: Next},
		Def{Name: OpenFile, ReqTarget: true},
		Def{Name: Outdent},
		Def{Name: Paste},
		Def{Name: Prev},
		Def{Name: Quit, NoReqBuffer: true},
		Def{Name: Redo},
		Def{Name: RenameFile, ReqTarget: true},
		Def{Name: RunCommand},
		Def{Name: RunMacro},
		Def{Name: SaveBuffer},
		Def{Name: SaveFileAs, ReqTarget: true},
		Def{Name: ScrollDown},
		Def{Name: ScrollUp},
		Def{Name: Search, ReqTarget: true},
		Def{Name: SearchResults},
		Def{Name: SelectBuffer},
		Def{Name: SetMark},
		Def{Name: StartGroupUndo},
		Def{Name: StartRecordMacro},
		Def{Name: StartSelection},
		Def{Name: StopGroupUndo},
		Def{Name: StopRecordMacro},
		Def{Name: StopSelection},
		Def{Name: Syntax},
		Def{Name: TypeHighlight},
		Def{Name: Undo},
		Def{Name: Up},
		Def{Name: Yank},
	)
	return &dl
}

// deflist is a list of action definitions
type deflist struct {
	list map[string]Def
	m    sync.RWMutex
}

// Add will add definitions
func (dl *deflist) Add(dd ...Def) {
	dl.m.Lock()
	defer dl.m.Unlock()
	for _, d := range dd {
		dl.list[d.Name] = d
	}
}

// Get will return a definition by name or nil if not found
func (dl *deflist) Get(n string) *Def {
	dl.m.RLock()
	defer dl.m.RUnlock()
	d, ok := dl.list[n]
	if !ok {
		return nil
	}
	return &d
}

// RequireBuffer will return true if the action requires a buffer
func (dl *deflist) RequireBuffer(n string) bool {
	dl.m.RLock()
	defer dl.m.RUnlock()
	def := dl.Get(n)
	if def == nil {
		return false
	}
	return !def.NoReqBuffer
}

// RequireTarget will return true if the action requires a target
func (dl *deflist) RequireTarget(n string) bool {
	dl.m.RLock()
	defer dl.m.RUnlock()
	def := dl.Get(n)
	if def == nil {
		return false
	}
	return !def.NoReqBuffer
}

// ValidAction will return an error for an invalid action nil otherwise
func (dl *deflist) ValidAction(act Action, bufid string) error {
	def := dl.Get(act.Name)
	if def == nil {
		return fmt.Errorf("Not found")
	}
	if def.ReqTarget && act.Target == "" {
		return fmt.Errorf("Requires target")
	}
	if def.ReqCount && act.Count <= 0 {
		return fmt.Errorf("Requires count > 0")
	}
	if def.ReqLine && act.Line <= 0 {
		return fmt.Errorf("Requires line > 0")
	}
	if def.ReqColumn && act.Column <= 0 {
		return fmt.Errorf("Requires column > 0")
	}
	if !def.NoReqBuffer && bufid == "" {
		return fmt.Errorf("Requires buffer")
	}
	return nil
}

// ValidRequest will return an error for an invalid request nil otherwise
func (dl *deflist) ValidRequest(req Request) error {
	for _, act := range req.Actions {
		if err := dl.ValidAction(act, req.BufferID); err != nil {
			return err
		}
	}
	return nil
}

// StrToAction will convert a string to an action
// it will return an error if the action is not found
func (dl *deflist) StrToAction(n string) (Action, error) {
	act := Action{}
	def := dl.Get(strings.ToLower(n))
	if def == nil {
		return act, fmt.Errorf("Action %v Not found", n)
	}
	act.Name = def.Name
	return act, nil
}
