package textobject

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dshills/layered/conf"
)

// Objects is a collection of objects
type Objects struct {
	config *conf.Configuration
	objs   map[string]TextObjecter
	m      sync.RWMutex
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
	name = strings.ToLower(name)
	o.m.RLock()
	defer o.m.RUnlock()
	obj, ok := o.objs[name]
	if !ok || obj == nil {
		return nil, fmt.Errorf("Objects: %v Not found", name)
	}
	return obj, nil
}

// Add will add an object to the collection
func (o *Objects) Add(objs ...TextObjecter) {
	o.m.Lock()
	defer o.m.Unlock()
	for i := range objs {
		o.objs[strings.ToLower(objs[i].Name())] = objs[i]
	}
}

// Remove will remove an object from the collection
func (o *Objects) Remove(name string) {
	o.m.Lock()
	defer o.m.Unlock()
	delete(o.objs, strings.ToLower(name))
}

// New returns a text object collection
func New(config *conf.Configuration) Objecter {
	objs := &Objects{objs: make(map[string]TextObjecter)}
	objs.loadDefaults()
	return objs
}
