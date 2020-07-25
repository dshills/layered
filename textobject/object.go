package textobject

import (
	"encoding/json"
	"io"
	"regexp"
)

// Object is a text object
type Object struct {
	name                                 string
	start                                string
	end                                  string
	multiline                            bool
	simple                               bool
	usefirst                             bool
	uselast                              bool
	altStartRx, altEndRx, startRx, endRx *regexp.Regexp
	altStart                             string
	altEnd                               string
}

// Name returns the text object name
func (o *Object) Name() string { return o.name }

// UseFirst will return true if match expects the use of the first match
func (o *Object) UseFirst() bool { return o.usefirst }

// UseLast will return true if match expects the use of the last match
func (o *Object) UseLast() bool { return o.uselast }

// MultiLine returns true if a obj is mutliple line
func (o *Object) MultiLine() bool { return o.multiline }

// Simple returns true if the object is simple
func (o *Object) Simple() bool { return o.simple }

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
		start:     js.Start,
		end:       js.End,
		usefirst:  js.UseFirst,
		altEnd:    js.AltEnd,
		altStart:  js.AltStart,
	}
	if err := obj.compile(); err != nil {
		return nil, err
	}
	return &obj, nil
}

func (o *Object) compile() error {
	rx, err := regexp.Compile(o.start)
	if err != nil {
		return err
	}
	o.startRx = rx
	if o.end != "" {
		rx, err := regexp.Compile(o.end)
		if err != nil {
			return err
		}
		o.endRx = rx
	}
	if o.altEnd != "" {
		rx, err := regexp.Compile(o.altEnd)
		if err != nil {
			return err
		}
		o.altEndRx = rx
	}
	if o.altStart != "" {
		rx, err := regexp.Compile(o.altStart)
		if err != nil {
			return err
		}
		o.altStartRx = rx
	}
	return nil
}

type jsObject struct {
	Name      string `json:"name"`
	Simple    bool   `json:"simple"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Multiline bool   `json:"multiline"`
	AltEnd    string `json:"alt_end"`
	AltStart  string `json:"alt_st"`
	UseFirst  bool   `json:"first_match"`
	UseLast   bool   `json:"last_match"`
}
