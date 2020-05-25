package layer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Layers is a set of layers
type Layers struct {
	ls []Layerer
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
			lay := Layer{}
			if err := lay.Load(f); err != nil {
				errs = append(errs, err.Error())
				continue
			}
			f.Close()
			l.Add(&lay)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("LoadAllModes: %v", strings.Join(errs, ", "))
	}
	return nil
}

// Add adds a layer
func (l *Layers) Add(a Layerer) {
	l.ls = append(l.ls, a)
}

// Remove will remove a layer
func (l *Layers) Remove(name string) {

}

// Layer will return a layer by name
func (l *Layers) Layer(name string) (Layerer, error) {
	for i := range l.ls {
		if l.ls[i].Name() == name {
			return l.ls[i], nil
		}
	}
	return nil, fmt.Errorf("Not found")
}
