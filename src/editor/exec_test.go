package editor

import (
	"fmt"
	"testing"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/filetype"
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
	ed, err = New(undo.New, textstore.New, buffer.New, cursor.New, syntax.New, filetype.New, textobject.New, rtpath)
	if err != nil {
		t.Error(err)
	}
	trans := action.NewTransaction("")
	trans.Set(action.New(action.NewBuffer, "", ""))
	resp, err := ed.Exec(trans)
	if err != nil {
		t.Error(err)
	}
	if resp.Buffer == "" {
		t.Errorf("Expected buffer id got none")
	}
	bufid = resp.Buffer
}

func TestReset(t *testing.T) {
	trans := action.NewTransaction(bufid)
	act := action.New(action.OpenFile, bufid, "/Users/dshills/Development/projects/goed-core/testdata/scanner.go")
	trans.Add(act)
	_, err := ed.Exec(trans)
	if err != nil {
		t.Error(err)
	}
}

/*
func TestSaveFileAs(t *testing.T) {
	ed, err := New(undo.New, textstore.New, buffer.New, cursor.New, syntax.New, filetype.New, textobject.New, rtpath)
	if err != nil {
		t.Error(err)
	}
	trans := action.NewTransaction("")
	trans.Set(action.New(action.OpenFile, "", "/Users/dshills/Development/projects/goed-core/testdata/scanner.go"))
	resp, err := ed.Exec(trans)
	if err != nil {
		t.Fatal(err)
	}
	trans.SetBuffer(resp.Buffer)
	trans.Set(action.New(action.SaveFileAs, "", "./scan.go"))
	_, err = ed.Exec(trans)
	if err != nil {
		t.Error(err)
	}
}
*/

func TestBufferList(t *testing.T) {
	trans := action.NewTransaction("")
	trans.Set(action.New(action.BufferList, "", ""))
	resp, err := ed.Exec(trans)
	if err != nil {
		t.Error(err)
	}
	if len(resp.Results) != 1 {
		t.Errorf("Expected BufferList of 1 got %v\n", len(resp.Results))
	}
}

func TestContent(t *testing.T) {
	trans := action.NewTransaction(bufid)
	act := action.New(action.Content, "", "")
	act.SetLine(45)
	act.SetCount(30)
	trans.Set(act)
	resp, err := ed.Exec(trans)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Content) != 30 {
		t.Errorf("Expected 15 got %v", len(resp.Content))
	}
}

func TestSyntax(t *testing.T) {
	trans := action.NewTransaction(bufid)
	act := action.New(action.Syntax, "", "")
	trans.Set(act)
	resp, err := ed.Exec(trans)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Syntax) == 0 {
		t.Errorf("Expected > 0 got 0")
	}
	for _, sr := range resp.Syntax {
		fmt.Println(sr.Token())
	}
}
