package action

import (
	"bytes"
	"encoding/json"
	"io"
)

// Request is a request for actions
type Request struct {
	BufferID   string   `json:"buffer_id"`
	LineOffset int      `json:"line_offset"`
	LineCount  int      `json:"line_count"`
	Actions    []Action `json:"actions"`
}

// Add will add actions to a request
func (r *Request) Add(act ...Action) {
	r.Actions = append(r.Actions, act...)
}

// AsJSON will return a json encoded request
func (r *Request) AsJSON() []byte {
	js, _ := json.Marshal(r)
	return js
}

// AsJSONReader returns a json encoded request, io.Reader
func (r *Request) AsJSONReader() io.Reader {
	return bytes.NewReader(r.AsJSON())
}

// NewRequest returns a Request
func NewRequest(bufid string, acts ...Action) Request {
	r := Request{BufferID: bufid}
	r.Add(acts...)
	return r
}

// JSONtoRequest will convert a json encoded request to a Request struct
func JSONtoRequest(js []byte) (Action, error) {
	act := Action{}
	err := json.Unmarshal(js, &act)
	return act, err
}

// ReaderToRequest will convert a json stream to a Request
func ReaderToRequest(r io.Reader) (Action, error) {
	act := Action{}
	err := json.NewDecoder(r).Decode(&act)
	return act, err
}
