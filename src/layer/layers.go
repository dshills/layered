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
	runtimes []string
	ls       []Layerer
}

// Load will load the layers within the runtimes
func (l *Layers) Load() error {
	errs := []string{}
	for _, rt := range l.runtimes {
		path := filepath.Join(rt, "layers")
		err := l._load(path)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("%v", strings.Join(errs, ", "))
	}
	return nil
}

func (l *Layers) _load(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
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
			lay := Layer{}
			lay.Load(f)
			l.Add(&lay)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("Layer.LoadDir: %v", strings.Join(errs, ", "))
	}
	return nil
}

// AddRuntime adds a runtime path
func (l *Layers) AddRuntime(rtpaths ...string) error {
	l.runtimes = append(l.runtimes, rtpaths...)
	return l.Load()
}

// RemoveRuntime will remove a runtime path
func (l *Layers) RemoveRuntime(path string) error {
	idx := -1
	for i := range l.runtimes {
		if l.runtimes[i] == path {
			idx = i
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("Not found")
	}
	l.runtimes = append(l.runtimes[:idx], l.runtimes[idx+1:]...)
	return l.Load()
}

// Add adds a layer
func (l *Layers) Add(a ...Layerer) {
	l.ls = append(l.ls, a...)
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

// New will return a layer manager
func New(rtpaths ...string) (Manager, error) {
	mng := &Layers{runtimes: rtpaths}
	err := mng.Load()
	return mng, err
}
