package action

import (
	"fmt"
)

// Action is an editor action
type Action struct {
	Name   string `json:"name"`
	Target string `json:"target"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
	Count  int    `json:"count"`
}

func (a Action) String() string {
	return fmt.Sprintf("{ Name: %q, Target: %q, Count: %v, Line: %v, Column: %v}", a.Name, a.Target, a.Count, a.Line, a.Column)
}
