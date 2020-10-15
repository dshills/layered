package layer

import (
	"testing"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

func TestInterp(t *testing.T) {
	rt := "/Users/dshills/Development/projects/layered/runtime/layers"
	interp := NewInterpriter()
	if err := interp.LoadDirectory(action.NewDefinitions(), rt); err != nil {
		t.Fatal(err)
	}
	test := [][]string{
		{"$"},
		{"0"},
		{"^"},
		{"b"},
		{"%"},
		{"e"},
		{"G"},
		{"g", "_"},
		{"g", "g"},
		{"h"},
		{"j"},
		{"k"},
		{"l"},
		{"N"},
		{"n"},
		{"W"},
		{"w"},
		{"a"},
		{"A"},
		{"I"},
		{"D"},
		{"d", "d"},
		{"x"},
		{"X"},
		{"O"},
		{"o"},
		{">", ">"},
		{"<", "<"},
		{"H"},
		{"L"},
		{"M"},
		{"z", "z"},
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
