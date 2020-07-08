package editor

import (
	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/syntax"
)

// Editorer is an editor interface
type Editorer interface {
	Exec(Request) Response
	ExecChan(reqC chan Request, respC chan Response, done chan struct{})
}

// Request is a request for actions
type Request struct {
	BufferID string
	Actions  []action.Action
}

// Add will add actions to a request
func (r *Request) Add(act ...action.Action) {
	r.Actions = append(r.Actions, act...)
}

// NewRequest returns a Request
func NewRequest(bufid string, acts ...action.Action) Request {
	r := Request{BufferID: bufid}
	r.Add(acts...)
	return r
}

// KeyValue is key/value data
type KeyValue struct {
	Key   string
	Value string
}

// Response is a exec response
type Response struct {
	Buffer         string
	Action         string
	Line, Column   int
	Results        []KeyValue
	Content        []string
	Syntax         []syntax.Resulter
	Search         []buffer.SearchResult
	Layer          string
	Partial        string
	ContentChanged bool
	CursorChanged  bool
	NewBuffer      bool
	CloseBuffer    bool
	InfoChanged    bool
	Quit           bool
	Err            error
}
