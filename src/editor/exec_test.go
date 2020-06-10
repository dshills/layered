package editor

import (
	"testing"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/layer"
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
	ed, err = New(undo.New, textstore.New, buffer.New, cursor.New, syntax.New, filetype.New, textobject.New, register.New, layer.New, rtpath)
	if err != nil {
		t.Error(err)
	}
	resp := ed.Exec("", action.Action{Name: action.NewBuffer})
	if resp.Err != nil {
		t.Error(resp.Err)
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
	resp := ed.Exec(bufid, act)
	if resp.Err != nil {
		t.Error(resp.Err)
	}
}

func TestDeleteLine(t *testing.T) {
	act := action.Action{Name: action.DeleteLine, Line: 1}
	resp := ed.Exec(bufid, act)
	if resp.Err != nil {
		t.Error(resp.Err)
	}
}

func TestUndo(t *testing.T) {
	resp := ed.Exec(bufid, action.Action{Name: action.Undo})
	if resp.Err != nil {
		t.Error(resp.Err)
	}
}

func TestBufferList(t *testing.T) {
	resp := ed.Exec(bufid, action.Action{Name: action.BufferList})
	if resp.Err != nil {
		t.Error(resp.Err)
	}
	if len(resp.Results) != 1 {
		t.Errorf("Expected BufferList of 1 got %v\n", len(resp.Results))
	}
}

func TestContent(t *testing.T) {
	act := action.Action{Name: action.Content, Line: 45, Count: 30}
	resp := ed.Exec(bufid, act)
	if resp.Err != nil {
		t.Error(resp.Err)
	}
	if len(resp.Content) != 30 {
		t.Errorf("Expected 15 got %v", len(resp.Content))
	}
}

func TestSyntax(t *testing.T) {
	resp := ed.Exec(bufid, action.Action{Name: action.Syntax})
	if resp.Err != nil {
		t.Error(resp.Err)
	}
	if len(resp.Syntax) == 0 {
		t.Errorf("Expected > 0 got 0")
	}
}

func TestSearch(t *testing.T) {
	act := action.Action{Name: action.Search, Target: "scan"}
	resp := ed.Exec(bufid, act)
	if resp.Err != nil {
		t.Error(resp.Err)
	}
	if len(resp.Search) == 0 {
		t.Errorf("Expected search results got none")
	}
}
