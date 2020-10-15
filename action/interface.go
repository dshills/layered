package action

// Factory is function that returns new action definitions
type Factory func() Definitions

// Definitions is a validator and string conversion tool
// for Actions and Requests
type Definitions interface {
	Add(dd ...Def)
	Get(n string) *Def
	RequireBuffer(n string) bool
	RequireTarget(n string) bool
	ValidAction(act Action, bufid string) error
	ValidRequest(req Request) error
	StrToAction(n string) (Action, error)
}

// Def is a definition for an action
type Def struct {
	Name        string
	NoReqBuffer bool
	ReqTarget   bool
	IsMovement  bool
	ReqCount    bool
	ReqLine     bool
	ReqColumn   bool
}
