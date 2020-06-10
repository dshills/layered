package textobject

import (
	"encoding/json"
	"io"
	"regexp"
)

// Object is a text object
type Object struct {
	name              string
	st, end           string
	multiline, simple bool
	startRx, endRx    *regexp.Regexp
}

// Name returns the text object name
func (o *Object) Name() string { return o.name }

// FindAll returns text object matches
func (o *Object) FindAll(s string) [][]int {
	if o.simple {
		return o.startRx.FindAllStringIndex(s, -1)
	}
	return nil
}

// FindAfter will return results after col
func (o *Object) FindAfter(s string, col int) [][]int {
	res := [][]int{}
	for _, r := range o.FindAll(s) {
		if r[0] > col {
			res = append(res, r)
		}
	}
	return res
}

// FindBefore will return results before col
func (o *Object) FindBefore(s string, col int) [][]int {
	res := [][]int{}
	for _, r := range o.FindAll(s) {
		if r[1] < col {
			res = append(res, r)
		}
	}
	return res
}

func loadObj(r io.Reader) (TextObjecter, error) {
	js := jsObject{}
	if err := json.NewDecoder(r).Decode(&js); err != nil {
		return nil, err
	}
	obj := Object{
		name:      js.Name,
		simple:    js.Simple,
		multiline: js.Multiline,
		st:        js.Start,
		end:       js.End,
	}
	rx, err := regexp.Compile(js.Start)
	if err != nil {
		return nil, err
	}
	obj.startRx = rx
	if js.End != "" {
		rx, err := regexp.Compile(js.End)
		if err != nil {
			return nil, err
		}
		obj.endRx = rx
	}
	return &obj, nil
}

type jsObject struct {
	Name      string `json:"name"`
	Simple    bool   `json:"simple"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Multiline bool   `json:"multiline"`
}
