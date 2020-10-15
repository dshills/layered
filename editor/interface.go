package editor

import "github.com/dshills/layered/action"

// Editorer is an editor interface
type Editorer interface {
	ExecChan(reqC chan action.Request, respC chan Response, done chan struct{})
}

// KeyValue is key/value data
type KeyValue struct {
	Key   string
	Value string
}
