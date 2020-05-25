package textobject

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Objects is a collection of objects
type Objects struct {
	objs []TextObjecter
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
	for i := range o.objs {
		if o.objs[i].Name() == name {
			return o.objs[i], nil
		}
	}
	return nil, fmt.Errorf("Objects.Object: Not found")
}

// Add will add an object to the collection
func (o *Objects) Add(objs ...TextObjecter) {
	o.objs = append(o.objs, objs...)
}

// Remove will remove an object from the collection
func (o *Objects) Remove(name string) {

}

// NewObjects returns a text object collection
func NewObjects() Objecter {
	return &Objects{}
}
