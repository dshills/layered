package layer

import (
	"testing"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

func TestInterpeter(t *testing.T) {
	rt := "/Users/dshills/Development/projects/layered/runtime/layers"
	interp := NewInterpreter(action.New(), "normal")
	if err := interp.LoadDirectory(rt); err != nil {
		t.Fatal(err)
	}

	acts := []action.Action{}
	var err error
	lay := interp.Active()
	for _, ka := range lay.KeyActions() {
		acts, err = interp.Match(ka.Keys()...)
		if err != nil {
			t.Error(err)
			continue
		}
		if len(acts) == 0 {
			t.Errorf("Expected actions got none for %v", ka.Keys())
		}
		k, err := key.StrToKeyer("<esc>")
		if err != nil {
			t.Error(err)
			continue
		}
		interp.Match(k)
	}
}
