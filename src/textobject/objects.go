package textobject

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Objects is a collection of objects
type Objects struct {
	runtimes []string
	objs     map[string]TextObjecter
	m        sync.RWMutex
}

// LoadDir will load a collection of text objects
func (o *Objects) LoadDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("Objects.LoadDir: %v", err)
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
			obj, err := loadObj(f)
			if err != nil {
				errs = append(errs, err.Error())
				continue
			}
			o.Add(obj)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("Objects.LoadDir: %v", strings.Join(errs, ", "))
	}
	return nil
}

// Object returns an object by name
func (o *Objects) Object(name string) (TextObjecter, error) {
	o.m.RLock()
	defer o.m.RUnlock()
	obj, ok := o.objs[name]
	if !ok {
		return nil, fmt.Errorf("Objects.Object: Not found")
	}
	return obj, nil
}

// Add will add an object to the collection
func (o *Objects) Add(objs ...TextObjecter) {
	o.m.Lock()
	defer o.m.Unlock()
	for i := range objs {
		o.objs[objs[i].Name()] = objs[i]
	}
}

// Remove will remove an object from the collection
func (o *Objects) Remove(name string) {
	o.m.Lock()
	defer o.m.Unlock()

}

// SetRuntimes will set the list of runtime directories
func (o *Objects) SetRuntimes(rts ...string) {
	o.runtimes = rts
	for i := len(rts) - 1; i >= 0; i-- {
		p := filepath.Join(rts[i], "objects")
		o.LoadDir(p)
	}
}

// AddRuntimes will add to the list of runtimes
func (o *Objects) AddRuntimes(rts ...string) {
	o.runtimes = append(o.runtimes, rts...)
	for i := len(rts) - 1; i >= 0; i-- {
		p := filepath.Join(rts[i], "objects")
		o.LoadDir(p)
	}
}

// New returns a text object collection
func New(rts ...string) Objecter {
	objs := &Objects{}
	objs.SetRuntimes(rts...)
	return objs
}
