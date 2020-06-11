package action

import (
	"fmt"
	"strings"
)

// Action is an editor action
type Action struct {
	Name, Target string
	Line, Column int
	Object       string
	Count        int
}

func (a Action) String() string {
	return fmt.Sprintf("{ Name: %q, Target: %q, Count: %v, Line: %v, Column: %v}", a.Name, a.Target, a.Count, a.Line, a.Column)
}

// Valid will return true if it is a valid action
func (a *Action) Valid(bufid string) error {
	errs := []string{}
	found := false
	an := strings.ToLower(a.Name)
	for _, d := range Definitions {
		if d.Name == an {
			found = true
			if a.Target == "" && d.ReqTarget {
				errs = append(errs, "Missing Target")
			}
			if bufid == "" && d.ReqBuffer {
				errs = append(errs, "Missing buffer id")
			}
		}
	}
	if !found {
		errs = append(errs, "Action %v not found", a.Name)
	}
	if len(errs) > 0 {
		return fmt.Errorf("Action: %v Invalid %v", a.Name, strings.Join(errs, ", "))
	}
	return nil
}

// NeedBuffer will return true if the action requires a buffer
func (a *Action) NeedBuffer() bool {
	for _, d := range Definitions {
		if d.Name == a.Name {
			if d.ReqBuffer {
				return true
			}
			return false
		}
	}
	return false
}
