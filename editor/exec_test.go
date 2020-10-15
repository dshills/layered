package editor

import (
	"fmt"
	"testing"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/conf"
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/register"
	"github.com/dshills/layered/syntax"
	"github.com/dshills/layered/textobject"
	"github.com/dshills/layered/textstore"
	"github.com/dshills/layered/undo"
)

const rtpath = "/Users/dshills/Development/projects/layered/runtime"
const tfile = "/Users/dshills/Development/projects/layered/testdata/scanner_test.go"

var ed Editorer
var bufid string
var reqC chan action.Request
var respC chan Response
var doneC chan struct{}

func setup() {
	if ed != nil {
		return
	}
	var err error
	config := &conf.Configuration{}
	config.AddRuntime(rtpath)
	ed, err = New(undo.New, textstore.New, buffer.New, cursor.New, syntax.New, filetype.New, textobject.New, register.New, action.New(), config)
	if err != nil {
		panic(err)
	}
	reqC = make(chan action.Request)
	respC = make(chan Response)
	doneC = make(chan struct{})
	ed.ExecChan(reqC, respC, doneC)
}

type testAct struct {
	action action.Action
	exLine int
	exCol  int
}

func newTestAct(name string, target string, exl, exc int) testAct {
	defs := action.New()
	act, err := defs.StrToAction(name)
	if err != nil {
		panic(err)
	}
	act.Target = target
	return testAct{action: act, exLine: exl, exCol: exc}
}

func TestMoveObj(t *testing.T) {
	defs := action.New()
	setup()
	// block-(),block-<>,block-[],block-{},bol-not-blankbol,eol-not-blank,eol,line,paragraph,sentence,string-double,string-single,string-tick,tag,word-ext,word
	reset := newTestAct("Move", "", 10, 0)
	reset.action.Line = 10
	reset.action.Column = 0

	tests := []testAct{}
	tests = append(tests, newTestAct("OpenFile", tfile, -1, -1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "bol-not-blank", 10, 1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "bol", 10, 0))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "eol", 10, 9))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "eol-not-blank", 10, 9))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "block-()", -1, -1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "block-<>", -1, -1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "block-[]", -1, -1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "block-{}", -1, -1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "paragraph", -1, -1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "sentence", -1, -1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "string-double", 10, 1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "string-single", -1, -1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "string-tick", -1, -1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "tag", -1, -1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "word-ext", -1, -1))
	tests = append(tests, reset)
	tests = append(tests, newTestAct("MoveObj", "word", -1, -1))
	tests = append(tests, reset)

	var buffer string
	for _, ta := range tests {
		if err := defs.ValidAction(ta.action, buffer); err != nil {
			t.Fatalf("Invalid action %v %+v", err, ta.action)
		}
		fmt.Printf("Request: %v %v %v:%v\n", ta.action.Name, ta.action.Target, ta.action.Line, ta.action.Column)
		reqC <- action.NewRequest(buffer, ta.action)
		resp := <-respC
		fmt.Printf("Response: %v %v => %v:%v\n", resp.Action.Name, resp.Action.Target, resp.Line, resp.Column)
		if resp.BufferID != "" {
			buffer = resp.BufferID
		}
		if resp.Err != nil {
			t.Error(resp.Err)
		}
		estr := ""
		if ta.exLine != -1 && resp.Line != ta.exLine {
			estr += fmt.Sprintf("Expected line %v got %v", ta.exLine, resp.Line)
		}
		if ta.exCol != -1 && resp.Column != ta.exCol {
			estr += fmt.Sprintf(" Expected column %v got %v", ta.exCol, resp.Column)
		}
		if estr != "" {
			t.Errorf("%#v => %v", resp.Action, estr)
		}
	}
}

/*
func TestExecAll(t *testing.T) {
	setup()
	wg := sync.WaitGroup{}
	go func(wg *sync.WaitGroup) {
		for {
			select {
			case <-doneC:
				return
			case resp := <-respC:
				wg.Done()
				if resp.Buffer != "" {
					bufid = resp.Buffer
				}
				if resp.Err != nil {
					t.Error(resp.Err)
				}
			}
		}
	}(&wg)
	acts := []action.Action{
		action.Action{Name: "OpenFile", Target: "testdata/scanner_test.go"},
		action.Action{Name: "BufferList"},
		action.Action{Name: "ChangeLayer"},
		action.Action{Name: "ChangePrevLayer"},
		action.Action{Name: "Content"},
		action.Action{Name: "Delete"},
		action.Action{Name: "DeleteChar"},
		action.Action{Name: "DeleteCharBack"},
		action.Action{Name: "DeleteCmdBack"},
		action.Action{Name: "DeleteLine"},
		action.Action{Name: "DeleteObject"},
		action.Action{Name: "DeleteToObject"},
		action.Action{Name: "Down"},
		action.Action{Name: "Indent"},
		action.Action{Name: "InsertLineAbove"},
		action.Action{Name: "InsertLine"},
		action.Action{Name: "InsertString"},
		action.Action{Name: "Move"},
		action.Action{Name: "MoveObj"},
		action.Action{Name: "MoveEnd"},
		action.Action{Name: "MovePrev"},
		action.Action{Name: "MovePrevEnd"},
		action.Action{Name: "CursorMovePast"},
		action.Action{Name: "NewBuffer"},
		action.Action{Name: "Next"},
		action.Action{Name: "Outdent"},
		action.Action{Name: "Paste"},
		action.Action{Name: "Prev"},
		action.Action{Name: "Quit"},
		action.Action{Name: "Redo"},
		action.Action{Name: "RenameFile"},
		action.Action{Name: "RunCommand"},
		action.Action{Name: "RunMacro"},
		action.Action{Name: "SaveBuffer"},
		action.Action{Name: "SaveFileAs"},
		action.Action{Name: "ScrollDown"},
		action.Action{Name: "ScrollUp"},
		action.Action{Name: "Search"},
		action.Action{Name: "SearchResults"},
		action.Action{Name: "SelectBuffer"},
		action.Action{Name: "SetMark"},
		action.Action{Name: "StartGroupUndo"},
		action.Action{Name: "StartRecordMacro"},
		action.Action{Name: "StartSelection"},
		action.Action{Name: "StopGroupUndo"},
		action.Action{Name: "StopRecordMacro"},
		action.Action{Name: "StopSelection"},
		action.Action{Name: "Syntax"},
		action.Action{Name: "Undo"},
		action.Action{Name: "Up"},
		action.Action{Name: "Yank"},
		action.Action{Name: "CloseBuffer"},
	}
	for _, act := range acts {
		wg.Add(1)
		for act.NeedBuffer() && bufid == "" {
			time.Sleep(10 * time.Millisecond)
		}
		fmt.Printf("Request: %#v\n", act)
		reqC <- NewRequest(bufid, act)
	}
	wg.Wait()
	doneC <- struct{}{}
}
*/
