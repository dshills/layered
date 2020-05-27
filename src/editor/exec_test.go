package editor

import (
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

func TestExect(t *testing.T) {
	act := action.New(action.NewBuffer, "", "")
	trans := action.NewTransaction("", act)
	ed, err := New(undo.New, textstore.New, buffer.New, cursor.New, syntax.New, filetype.New, textobject.New, rtpath)
	if err != nil {
		t.Error(err)
	}
	if err := ed.Exec(trans); err != nil {
		t.Error(err)
	}

}
