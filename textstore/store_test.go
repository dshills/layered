package textstore

import (
	"os"
	"testing"

	"github.com/dshills/layered/undo"
)

var store = New(undo.New)

func TestReset(t *testing.T) {
	str := "This is a test"
	store.Reset(str)
}

func TestReadFrom(t *testing.T) {
	f, err := os.Open("testdata/test1.txt")
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}
	store.ReadFrom(f)
}

func TestNewLine(t *testing.T) {
	txt := "0 Added line"
	store.NewLine(0, txt)
	str, err := store.LineString(0)
	if err != nil {
		t.Error(err)
	}
	if str != txt {
		t.Errorf("Expected %q got %q\n", txt, str)
	}
}

func TestDeleteLine(t *testing.T) {
	txt := "1 This is a test of the emergency broadcast system"
	store.DeleteLine(0)
	str, err := store.LineString(0)
	if err != nil {
		t.Error(err)
	}
	if str != txt {
		t.Errorf("Expected %q got %q\n", txt, str)
	}
}

func TestResetLine(t *testing.T) {
	txt := "0 Ta-Da"
	store.ResetLine(0, txt)
	str, err := store.LineString(0)
	if err != nil {
		t.Error(err)
	}
	if str != txt {
		t.Errorf("Expected %q got %q\n", txt, str)
	}
}

func TestLineRange(t *testing.T) {
	strs, err := store.LineRange(0, 10)
	if err != nil {
		t.Error(err)
	}
	if len(strs) != 10 {
		t.Errorf("Expected 10 got %v\n", len(strs))
	}
}

func TestNumLines(t *testing.T) {
	exp := 29
	ln := store.NumLines()
	if ln != exp {
		t.Errorf("Expected %v got %v", exp, ln)
	}
}

func TestLineLen(t *testing.T) {
	exp := 7
	ln := store.LineLen(0)
	if ln != exp {
		t.Errorf("Expected %v got %v", exp, ln)
	}
}

func TestLen(t *testing.T) {
	exp := 6945
	ln := store.Len()
	if ln != exp {
		t.Errorf("Expected %v got %v", exp, ln)
	}
}

func TestReadRuneAt(t *testing.T) {
	exp := '?'
	r, err := store.RuneAt(9, 4)
	if err != nil {
		t.Error(err)
	}
	if r != exp {
		t.Errorf("Expected %v got %v\n", exp, r)
	}
}
func TestLineAt(t *testing.T) {
	exp := "4 Line 4"
	str, err := store.LineString(3)
	if err != nil {
		t.Error(err)
	}
	if str != exp {
		t.Errorf("Expected %q got %q\n", exp, str)
	}
}

func TestSetLineDelim(t *testing.T) {
	store.SetLineDelim("\n")
}

func TestLineDelim(t *testing.T) {
	exp := "\n"
	str := store.LineDelim()
	if str != exp {
		t.Errorf("Expected %q got %q\n", exp, str)
	}
}

func TestUndo(t *testing.T) {
	txt := "1 This is a test of the emergency broadcast system"
	store.Undo()
	str, err := store.LineString(0)
	if err != nil {
		t.Error(err)
	}
	if str != txt {
		t.Errorf("Expected %q got %q\n", txt, str)
	}
}

func TestRedo(t *testing.T) {
	txt := "0 Ta-Da"
	store.Redo()
	str, err := store.LineString(0)
	if err != nil {
		t.Error(err)
	}
	if str != txt {
		t.Errorf("Expected %q got %q\n", txt, str)
	}
}
