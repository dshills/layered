package palette

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Color is a rgb color
type Color struct {
	Red         uint8
	Green       uint8
	Blue        uint8
	Transparent bool
}

// RGBA satisfies image.Color
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.Red)
	g = uint32(c.Green)
	b = uint32(c.Blue)
	a = 1
	return
}

// NewColor returns a new rgb color
func NewColor(r, g, b uint8) Color {
	return Color{Red: r, Green: g, Blue: b}
}

func parseRGB(str string) (clr Color, err error) {
	const rgbFormat = "rgb(%d,%d,%d)"
	var r, g, b uint8
	_, err = fmt.Sscanf(str, rgbFormat, &r, &g, &b)
	if err != nil {
		err = fmt.Errorf("%v %v", str, err)
		return
	}
	clr = NewColor(r, g, b)
	return
}

// ColorList is a list of colors
type ColorList struct {
	colors map[string]Color
	cm     sync.RWMutex
	terms  map[int]Color
	tm     sync.RWMutex
}

// Color will return a color by name
func (cl *ColorList) Color(name string) (Color, error) {
	cl.tm.RLock()
	defer cl.tm.RUnlock()
	clr, ok := cl.colors[strings.ToLower(name)]
	if !ok {
		return Color{}, fmt.Errorf("%v Not found", name)
	}
	return clr, nil
}

// Term will return a terminal color by id
func (cl *ColorList) Term(tid int) (Color, error) {
	cl.tm.RLock()
	defer cl.tm.RUnlock()
	clr, ok := cl.terms[tid]
	if !ok {
		return Color{}, fmt.Errorf("term %v Not found", tid)
	}
	return clr, nil
}

// Add will add a color
func (cl *ColorList) Add(name string, clr Color) {
	cl.cm.Lock()
	defer cl.cm.Unlock()
	cl.colors[strings.ToLower(name)] = clr
}

// AddTerm will add a terminal color
func (cl *ColorList) AddTerm(tid int, clr Color) {
	cl.tm.Lock()
	defer cl.tm.Unlock()
	cl.terms[tid] = clr
}

// Load will load a color file
func (cl *ColorList) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	jsc := []jsColor{}
	if err = json.NewDecoder(f).Decode(&jsc); err != nil {
		return err
	}
	errs := []string{}
	for i := range jsc {
		clr, err := parseRGB(jsc[i].Color)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		cl.Add(jsc[i].Name, clr)
		if jsc[i].TermID > 0 {
			cl.AddTerm(jsc[i].TermID, clr)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("ColorList.Load: %v", strings.Join(errs, ", "))
	}
	return nil
}

// NewColorList will return a color list
func NewColorList() ColorList {
	return ColorList{colors: make(map[string]Color), terms: make(map[int]Color)}
}

type jsColor struct {
	Name   string
	Color  string
	TermID int
}
