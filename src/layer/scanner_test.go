package layer

import (
	"path/filepath"
	"testing"

	"github.com/dshills/layered/key"
)

const rt = "/Users/dshills/Development/projects/layered/runtime"

type tst struct {
	keys []key.Keyer
}

func newTst(isctrl bool, rs ...rune) tst {
	ts := tst{}
	for _, r := range rs {
		ts.keys = append(ts.keys, key.New(r, 0))
	}
	return ts
}

func TestScanNormal(t *testing.T) {
	var tests = []tst{
		newTst(false, ':'),
		newTst(false, '/'),
		newTst(false, 'i'),
		newTst(false, 'V'),
		newTst(false, 'v'),
		//newTst(true, 'v'),
		newTst(false, '$'),
		newTst(false, '0'),
		newTst(false, '^'),
		newTst(false, 'b'),
		newTst(false, '%'),
		newTst(false, 'e'),
		newTst(false, 'G'),
		newTst(false, 'g', '_'),
		newTst(false, 'g', 'g'),
		newTst(false, 'h'),
		newTst(false, 'j'),
		newTst(false, 'k'),
		newTst(false, 'l'),
		newTst(false, 'N'),
		newTst(false, 'n'),
		newTst(false, 'W'),
		newTst(false, 'w'),
		newTst(false, 'a'),
		newTst(false, 'A'),
		newTst(false, 'I'),
		newTst(false, 'D'),
		newTst(false, 'd', 'd'),
		newTst(false, 'x'),
		newTst(false, 'X'),
		newTst(false, 'O'),
		newTst(false, 'o'),
		newTst(false, '>', '>'),
		newTst(false, '<', '<'),
		newTst(false, 'H'),
		newTst(false, 'L'),
		newTst(false, 'M'),
		newTst(false, 'z', 'z'),
		//newTst(true, 'd'),
		//newTst(true, 'u'),
		//newTst(true, 'r'),
		newTst(false, 'u'),
		newTst(false, 'p'),
		newTst(false, 'y', 'y'),
	}
	layers := Layers{}
	fp := filepath.Join(rt, "layers")
	if err := layers.LoadDir(fp); err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		scanner, err := NewScanner(&layers, "normal")
		if err != nil {
			t.Error(err)
			continue
		}
		for _, key := range test.keys {
			acts, status, err := scanner.Scan(key)
			if err != nil {
				t.Errorf("%v %v\n", test, err)
				break
			}
			if status == NoMatch {
				t.Errorf("Expected match or partial got NoMatch %v\n", test.keys)
				break
			}
			if status == Match && len(acts) == 0 {
				t.Errorf("Expected match actions got none")
			}
		}
	}
}

/*
func TestScanSearch(t *testing.T) {
	var tests = []tst{
		//tst{keys: []key.Keyer{key.NewKey(false, false, true, 0, "<esc>")}},
		//tst{keys: []key.Keyer{key.NewKey(false, false, true, 0, "<cr>")}},
		//tst{keys: []key.Keyer{key.NewKey(false, false, true, 0, "<bs>")}},
		tst{keys: []key.Keyer{key.New(false, false, false, 'a', "<bs>"), key.NewKey(false, false, true, 0, "<bs>")}},
		tst{keys: []key.Keyer{key.New(false, false, false, 'a', "")}},
	}
	layers := Layers{}
	fp := filepath.Join(rt, "layers")
	if err := layers.LoadDir(fp); err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		scanner, err := NewScanner(&layers, "search")
		if err != nil {
			t.Error(err)
			continue
		}
		for _, key := range test.keys {
			acts, status, err := scanner.Scan(key)
			if err != nil {
				t.Errorf("%v %v\n", test, err)
				break
			}
			if status == NoMatch {
				t.Errorf("Expected match or partial got NoMatch %v\n", test.keys)
				break
			}
			if status == Match && len(acts) == 0 {
				t.Errorf("Expected match actions got none")
			}
		}
	}
}
*/
