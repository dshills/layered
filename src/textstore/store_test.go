package textstore

import (
	"testing"

	"github.com/dshills/layered/undo"
)

var store = New(undo.New)

func TestReset(t *testing.T) {
	str := "This is a test\nThis is only a test\nOf the emergency broadcast system"
	store.Reset(str)
	if store.NumLines() != 3 {
		t.Errorf("Expected 3 lines got %v\n", store.NumLines())
	}
}

func TestNewLine(t *testing.T) {
	store.NewLine("Extra line", 1)
	if store.NumLines() != 4 {
		t.Errorf("Expected 4 lines got %v\n", store.NumLines())
	}
	str, err := store.LineString(1)
	if err != nil {
		t.Error(err)
	}
	if str != "Extra line" {
		t.Errorf("Expected 'Extra line' got %v\n", str)
	}
}

func TestDeleteLine(t *testing.T) {
	store.DeleteLine(1)
	if store.NumLines() != 3 {
		t.Errorf("Expected 3 lines got %v\n", store.NumLines())
	}
}

func TestResetLine(t *testing.T) {
	str := "This is a reset line"
	store.ResetLine(str, 1)
	ls, err := store.LineString(1)
	if err != nil {
		t.Error(err)
	}
	if ls != str {
		t.Errorf("Expected %v got %v\n", str, ls)
	}
}

func TestUndo(t *testing.T) {
	err := store.Undo()
	if err != nil {
		t.Error(err)
	}
	err = store.Undo()
	if err != nil {
		t.Error(err)
	}
}
