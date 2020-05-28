package editor

import (
	"fmt"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/syntax"
)

// KeyValue is key/value data
type KeyValue struct {
	Key   string
	Value string
}

// Response is a exec response
type Response struct {
	Buffer       string
	Action       string
	Line, Column int
	Results      []KeyValue
	Content      []string
	Syntax       []syntax.Resulter
	Search       []buffer.SearchResult
}

// Exec will execute a transaction in the editor
func (e *Editor) Exec(tr action.Transactioner) (resp *Response, err error) {
	acts := tr.Actions()
	resp = &Response{Buffer: tr.Buffer()}
	if err = tr.Valid(); err != nil {
		return
	}
	var buf buffer.Bufferer
	if tr.NeedBuffer() {
		buf, err = e.Buffer(tr.Buffer())
		if err != nil {
			return
		}
	}

	for _, act := range acts {
		switch act.Name() {

		// Search
		case action.Search:
			res, err := buf.Search(act.Param())
			if err != nil {
				return nil, err
			}
			resp.Search = res
		case action.SearchResults:
			resp.Search = buf.SearchResults()

		// Syntax
		case action.Syntax:
			resp.Syntax = buf.SyntaxResults()

		// File / buffer
		case action.BufferList:
			for i := range e.bufs {
				resp.Results = append(resp.Results, KeyValue{Key: e.bufs[i].ID(), Value: e.bufs[i].Filename()})
			}
			return
		case action.NewBuffer:
			resp.Buffer = e.newBuffer()
		case action.SaveBuffer:
			if err := buf.SaveBuffer(act.Param()); err != nil {
				return nil, err
			}
		case action.CloseBuffer:
			if buf.Dirty() {
				return nil, fmt.Errorf("CloseBuffer: Buffer is dirty")
			}
			e.Remove(tr.Buffer())
		case action.OpenFile:
			id := tr.Buffer()
			if id == "" {
				id = e.newBuffer()
				resp.Buffer = id
			}
			buf, err := e.Buffer(id)
			if err != nil {
				return nil, err
			}
			if err := buf.OpenFile(act.Param()); err != nil {
				return nil, err
			}
		case action.RenameFile:
			if err := buf.RenameFile(act.Param()); err != nil {
				return nil, err
			}
		case action.SaveFileAs:
			buf.SetFilename(act.Param())
			if err := buf.SaveBuffer(""); err != nil {
				return nil, err
			}
		// Move
		case action.Move:
			obj, err := e.objs.Object(act.Object())
			if err != nil {
				return nil, err
			}
			buf.Move(1, obj)
		case action.MoveEnd:
			obj, err := e.objs.Object(act.Object())
			if err != nil {
				return nil, err
			}
			buf.MoveEnd(1, obj)
		case action.MovePrev:
			obj, err := e.objs.Object(act.Object())
			if err != nil {
				return nil, err
			}
			buf.MovePrev(1, obj)
		case action.MovePrevEnd:
			obj, err := e.objs.Object(act.Object())
			if err != nil {
				return nil, err
			}
			if err := buf.MovePrevEnd(1, obj); err != nil {
				return nil, err
			}
		case action.Up:
			buf.Up(1)
		case action.Down:
			buf.Down(1)
		case action.Prev:
			buf.Prev(1)
		case action.Next:
			buf.Next(1)
		case action.ScrollDown:
			buf.ScrollDown()
		case action.ScrollUp:
			buf.ScrollUp()

		// Edit
		case action.DeleteChar:
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteChar(act.Line(), act.Column(), act.Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.DeleteCharBack:
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteCharBack(act.Line(), act.Column(), act.Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.DeleteLine:
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteLine(act.Line(), act.Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.DeleteObject:
			obj, err := e.objs.Object(act.Object())
			if err != nil {
				return nil, err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteObject(act.Line(), act.Column(), obj, act.Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.InsertLine:
			buf.TextStore().StartGroupUndo()
			err = buf.NewLineBelow(act.Line(), act.Param(), act.Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.InsertLineAbove:
			buf.TextStore().StartGroupUndo()
			err = buf.NewLineAbove(act.Line(), act.Param(), act.Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.InsertString:
			if err := buf.InsertString(act.Line(), act.Column(), act.Param()); err != nil {
				return nil, err
			}
		case action.Indent:
			if err := buf.Indent(act.Line(), act.Count()); err != nil {
				return nil, err
			}
		case action.Outdent:
			if err := buf.Outdent(act.Line(), act.Count()); err != nil {
				return nil, err
			}
		case action.Content:
			resp.Content, err = buf.TextStore().LineRangeString(act.Line(), act.Count())
			return
		}

	}
	return
}
