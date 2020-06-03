package editor

import (
	"testing"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/register"
	"github.com/dshills/layered/syntax"
	"github.com/dshills/layered/textobject"
	"github.com/dshills/layered/textstore"
	"github.com/dshills/layered/undo"
)

const rtpath = "/Users/dshills/Development/projects/layered/runtime"

var ed Editorer
var bufid string

func TestExect(t *testing.T) {
	var err error
	ed, err = New(undo.New, textstore.New, buffer.New, cursor.New, syntax.New, filetype.New, textobject.New, register.New, rtpath)
	if err != nil {
		t.Error(err)
	}
	resp, err := ed.Exec("", action.Action{Name: action.NewBuffer})
	if err != nil {
		t.Error(err)
	}
	if resp.Buffer == "" {
		t.Errorf("Expected buffer id got none")
	}
	bufid = resp.Buffer
}

func TestReset(t *testing.T) {
	act := action.Action{
		Name:   action.OpenFile,
		Target: "/Users/dshills/Development/projects/goed-core/testdata/scanner.go",
	}
	_, err := ed.Exec(bufid, act)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteLine(t *testing.T) {
	act := action.Action{Name: action.DeleteLine, Line: 1}
	_, err := ed.Exec(bufid, act)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUndo(t *testing.T) {
	_, err := ed.Exec(bufid, action.Action{Name: action.Undo})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBufferList(t *testing.T) {
	resp, err := ed.Exec(bufid, action.Action{Name: action.BufferList})
	if err != nil {
		t.Error(err)
	}
	if len(resp.Results) != 1 {
		t.Errorf("Expected BufferList of 1 got %v\n", len(resp.Results))
	}
}

func TestContent(t *testing.T) {
	act := action.Action{Name: action.Content, Line: 45, Count: 30}
	resp, err := ed.Exec(bufid, act)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Content) != 30 {
		t.Errorf("Expected 15 got %v", len(resp.Content))
	}
}

func TestSyntax(t *testing.T) {
	resp, err := ed.Exec(bufid, action.Action{Name: action.Syntax})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Syntax) == 0 {
		t.Errorf("Expected > 0 got 0")
	}
}

func TestSearch(t *testing.T) {
	act := action.Action{Name: action.Search, Target: "scan"}
	resp, err := ed.Exec(bufid, act)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Search) == 0 {
		t.Errorf("Expected search results got none")
	}
}
