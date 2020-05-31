package action

import (
	"fmt"
	"strings"
)

// Action is an editor action
type Action struct {
	act, target, param string
	ln, col            int
	obj                string
	cnt                int
}

// Name returns the action name
func (a *Action) Name() string { return a.act }

// Target is the target of the action
func (a *Action) Target() string { return a.target }

// SetTarget will set the target
func (a *Action) SetTarget(t string) { a.target = t }

// Param is a required parameter
func (a *Action) Param() string { return a.param }

// SetParam will set the param
func (a *Action) SetParam(p string) { a.param = p }

// Line will return the line -1 represents not set
func (a *Action) Line() int { return a.ln }

// Column returns the column -1 represents not set
func (a *Action) Column() int { return a.col }

// SetLine will set the line associated with the action
func (a *Action) SetLine(l int) { a.ln = l }

// SetColumn will set the column associated with the action
func (a *Action) SetColumn(c int) { a.col = c }

// Object returns the object associated with the action
func (a *Action) Object() string { return a.obj }

// SetObject will set the object for the action
func (a *Action) SetObject(obj string) { a.obj = obj }

// Count will return the action count
func (a *Action) Count() int { return a.cnt }

// SetCount will set the action count
func (a *Action) SetCount(c int) { a.cnt = c }

// Valid will return true if it is a valid action
func (a *Action) Valid(bufid string) error {
	errs := []string{}
	found := false
	for _, d := range Definitions {
		if d.Name == a.Name() {
			found = true
			if a.Param() == "" && d.ReqParam {
				errs = append(errs, "Missing Param")
			}
			if a.Target() == "" && d.ReqTarget {
				errs = append(errs, "Missing Target")
			}
			if bufid == "" && d.ReqBuffer {
				errs = append(errs, "Missing buffer id")
			}
		}
	}
	if !found {
		errs = append(errs, "Action not found")
	}
	if len(errs) > 0 {
		return fmt.Errorf("Action: %v Invalid %v", a.Name(), strings.Join(errs, ", "))
	}
	return nil
}

// NeedBuffer will return true if the action requires a buffer
func (a *Action) NeedBuffer() bool {
	for _, d := range Definitions {
		if d.Name == a.Name() {
			if d.ReqBuffer {
				return true
			}
			return false
		}
	}
	return false
}

// New will return a new Actioner
func New(act string) Actioner {
	return &Action{act: act, ln: -1, col: -1, cnt: 1}
}

// Transaction is a group of actions with a buffer identifier
type Transaction struct {
	acts []Actioner
	id   string
}

// Buffer will return the buffer id
func (t *Transaction) Buffer() string { return t.id }

// SetBuffer will set the transaction buffer
func (t *Transaction) SetBuffer(id string) {
	t.id = id
}

// Actions returns the set of actions
func (t *Transaction) Actions() []Actioner { return t.acts }

// Add will add actions to the transactions
func (t *Transaction) Add(acts ...Actioner) {
	t.acts = append(t.acts, acts...)
}

// Set will set actions to the transactions
func (t *Transaction) Set(acts ...Actioner) {
	t.acts = acts
}

// Valid will return true if a valid transaction
func (t *Transaction) Valid() error {
	errs := []string{}
	for i := range t.acts {
		if err := t.acts[i].Valid(t.id); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("Invalid transaction %v", strings.Join(errs, ", "))
	}
	return nil
}

// NeedBuffer will return true if the action list requires a buffer
func (t *Transaction) NeedBuffer() bool {
	for i := range t.acts {
		if t.acts[i].NeedBuffer() {
			return true
		}
	}
	return false
}

// NewTransaction will return an we transactioner
func NewTransaction(id string, actions ...Actioner) Transactioner {
	return &Transaction{id: id, acts: actions}
}
