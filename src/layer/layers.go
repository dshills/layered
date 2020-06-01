package layer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Layers is a set of layers
type Layers struct {
	ls  []Layerer
	def Layerer
}

// LoadDir wil load layers from a directory
func (l *Layers) LoadDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("LoadDir: %v", err)
	}

	errs := []string{}
	for _, file := range files {
		if strings.ToLower(filepath.Ext(file.Name())) == ".json" {
			fn := filepath.Join(dir, file.Name())
			f, err := os.Open(fn)
			if err != nil {
				errs = append(errs, err.Error())
				continue
			}
			defer f.Close()
			js := jsLayer{}
			if err := json.NewDecoder(f).Decode(&js); err != nil {
				errs = append(errs, err.Error())
				continue
			}
			lay := js.asLayer()
			l.Add(lay)
			if lay.IsDefault() {
				l.def = lay
			}
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("Layer.LoadDir: %v", strings.Join(errs, ", "))
	}
	if l.def == nil {
		return fmt.Errorf("Layers.LoadDir: No default layer defined")
	}
	return nil
}

// Add adds a layer
func (l *Layers) Add(a Layerer) {
	l.ls = append(l.ls, a)
}

// Remove will remove a layer
func (l *Layers) Remove(name string) {
	name = strings.ToLower(name)
	for i, lay := range l.ls {
		if strings.ToLower(lay.Name()) == name {
			// Order not perserved
			l.ls[i] = l.ls[len(l.ls)-1]
			l.ls = l.ls[:len(l.ls)-1]
			return
		}
	}
}

// Default will return the default layer
func (l *Layers) Default() Layerer {
	return l.def
}

// Layer will return a layer by name
func (l *Layers) Layer(name string) (Layerer, error) {
	name = strings.ToLower(name)
	for i := range l.ls {
		if strings.ToLower(l.ls[i].Name()) == name {
			return l.ls[i], nil
		}
	}
	return nil, fmt.Errorf("Not found")
}
