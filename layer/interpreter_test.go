package layer

import (
	"testing"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

func TestInterp(t *testing.T) {
	rt := "/Users/dshills/Development/projects/layered/runtime/layers"
	interp := NewInterpriter()
	if err := interp.LoadDirectory(rt); err != nil {
		t.Fatal(err)
	}
	test := [][]string{
		[]string{"$"},
		[]string{"0"},
		[]string{"^"},
		[]string{"b"},
		[]string{"%"},
		[]string{"e"},
		[]string{"G"},
		[]string{"g", "_"},
		[]string{"g", "g"},
		[]string{"h"},
		[]string{"j"},
		[]string{"k"},
		[]string{"l"},
		[]string{"N"},
		[]string{"n"},
		[]string{"W"},
		[]string{"w"},
		[]string{"a"},
		[]string{"A"},
		[]string{"I"},
		[]string{"D"},
		[]string{"d", "d"},
		[]string{"x"},
		[]string{"X"},
		[]string{"O"},
		[]string{"o"},
		[]string{">", ">"},
		[]string{"<", "<"},
		[]string{"H"},
		[]string{"L"},
		[]string{"M"},
		[]string{"z", "z"},
	}

	acts := []action.Action{}
	for _, kk := range test {
		for _, k := range kk {
			ak, err := key.StrToKeyer(k)
			if err != nil {
				t.Error(err)
				continue
			}
			acts = interp.Match(ak)
		}
		if len(acts) == 0 {
			t.Errorf("Expected actions got none for %v", kk)
		}
	}

}
