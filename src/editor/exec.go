package editor

import (
	"fmt"
	"strings"

	"github.com/dshills/layered/action"
)

// Exec will execute a transaction in the editor
func (e *Editor) Exec(req Request) Response {
	resp := Response{Buffer: req.BufferID}
	need := false
	for _, act := range req.Actions {
		if err := act.Valid(req.BufferID); err != nil {
			resp.Err = err
			return resp
		}
		if act.NeedBuffer() {
			need = true
		}
	}

	buf, err := e.buffer(req.BufferID)
	if err != nil && need {
		resp.Err = err
		return resp
	}

	for _, act := range req.Actions {
		//logger.Debugf("Editor.Exec: %v %v", bufid, act.Name)
		switch strings.ToLower(act.Name) {

		case action.Quit:
			if len(e.bufs) > 1 {
				e.remove(e.activeBufID)
				e.activeBufID = e.bufs[len(e.bufs)-1].ID()
				resp.Buffer = e.activeBufID
				return resp
			}
			resp.Quit = true
			return resp

		case action.SelectBuffer:
			e.activeBufID = req.BufferID

		// Search
		case action.Search:
			res, err := buf.Search(act.Target)
			if err != nil {
				resp.Err = err
				return resp
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
			return resp
		case action.NewBuffer:
			resp.Buffer = e.newBuffer()
			resp.NewBuffer = true
		case action.SaveBuffer:
			if err := buf.SaveBuffer(act.Target); err != nil {
				resp.Err = err
				return resp
			}
		case action.CloseBuffer:
			if buf.Dirty() {
				resp.Err = fmt.Errorf("CloseBuffer: Buffer is dirty")
				return resp
			}
			e.remove(req.BufferID)
			resp.CloseBuffer = true
		case action.OpenFile:
			id := req.BufferID
			if id == "" {
				resp.NewBuffer = true
				id = e.newBuffer()
				resp.Buffer = id
			}
			buf, err := e.buffer(id)
			if err != nil {
				resp.Err = err
				return resp
			}
			if err := buf.OpenFile(act.Target); err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.RenameFile:
			if err := buf.RenameFile(act.Target); err != nil {
				resp.Err = err
				return resp
			}
			resp.InfoChanged = true
		case action.SaveFileAs:
			buf.SetFilename(act.Target)
			if err := buf.SaveBuffer(""); err != nil {
				resp.Err = err
				return resp
			}
			resp.InfoChanged = true

		// Move
		case action.Move:
			obj, err := e.objs.Object(act.Object)
			if err != nil {
				resp.Err = err
				return resp
			}
			buf.Move(act.Count, obj)
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true
		case action.MoveEnd:
			obj, err := e.objs.Object(act.Object)
			if err != nil {
				resp.Err = err
				return resp
			}
			buf.MoveEnd(act.Count, obj)
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true
		case action.MovePrev:
			obj, err := e.objs.Object(act.Object)
			if err != nil {
				resp.Err = err
				return resp
			}
			buf.MovePrev(act.Count, obj)
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true
		case action.MovePrevEnd:
			obj, err := e.objs.Object(act.Object)
			if err != nil {
				resp.Err = err
				return resp
			}
			if err := buf.MovePrevEnd(act.Count, obj); err != nil {
				resp.Err = err
				return resp
			}
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true
		case action.Up:
			buf.Up(act.Count)
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true
		case action.Down:
			buf.Down(act.Count)
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true
		case action.Prev:
			buf.Prev(act.Count)
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true
		case action.Next:
			buf.Next(act.Count)
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true
		case action.ScrollDown:
			buf.ScrollDown()
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true
		case action.ScrollUp:
			buf.ScrollUp()
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true

		// Edit
		case action.DeleteChar:
			err := buf.DeleteChar(act.Line, act.Column, act.Count)
			if err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.DeleteCharBack:
			if err := buf.DeleteCharBack(act.Line, act.Column, act.Count); err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.DeleteLine:
			if err := buf.DeleteLine(act.Line, act.Count); err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.DeleteObject:
			obj, err := e.objs.Object(act.Object)
			if err != nil {
				resp.Err = err
				return resp
			}
			if err = buf.DeleteObject(act.Line, act.Column, obj, act.Count); err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.InsertLine:
			if err := buf.NewLineBelow(act.Line, act.Target, act.Count); err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.InsertLineAbove:
			if err := buf.NewLineAbove(act.Line, act.Target, act.Count); err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.InsertString:
			if err := buf.InsertString(act.Line, act.Column, act.Target); err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.Indent:
			if err := buf.Indent(act.Line, act.Count); err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.Outdent:
			if err := buf.Outdent(act.Line, act.Count); err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.Content:
			var err error
			resp.Content, err = buf.TextStore().LineRangeString(act.Line, act.Count)
			if err != nil {
				resp.Err = err
				return resp
			}
			resp.Syntax = buf.SyntaxResults()
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()

		// Undo / redo
		case action.Undo:
			if err := buf.Undo(); err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.Redo:
			if err := buf.Redo(); err != nil {
				resp.Err = err
				return resp
			}
			resp.ContentChanged = true
		case action.StartGroupUndo:
			buf.StartGroupUndo()
		case action.StopGroupUndo:
			buf.StopGroupUndo()
		}

	}
	return resp
}
