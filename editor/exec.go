package editor

import (
	"fmt"
	"strings"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
)

// ExecChan will listen for requests
func (e *Editor) ExecChan(reqC chan action.Request, respC chan Response, done chan struct{}) {
	go func(reqC chan action.Request, respC chan Response, done chan struct{}) {
		for {
			select {
			case req := <-reqC:
				e.exec(req, respC)
			case <-done:
				return
			}
		}
	}(reqC, respC, done)
}

// Exec will execute a transaction in the editor
func (e *Editor) exec(req action.Request, respC chan Response) {
	buf, err := e.validateRequest(req)
	if err != nil {
		respC <- Response{BufferID: req.BufferID, Err: err}
	}
	req = e.updatePos(req, buf)
	for _, act := range req.Actions {
		resp := Response{BufferID: req.BufferID, Action: act, ContentChanged: true}
		switch strings.ToLower(act.Name) {

		case action.Quit:
			if len(e.bufs) > 1 {
				e.remove(e.activeBufID)
				e.activeBufID = e.bufs[len(e.bufs)-1].ID()
				resp.BufferID = e.activeBufID
				resp.CloseBuffer = true
				respC <- resp
				return
			}
			resp.Quit = true
			resp.ContentChanged = false
			respC <- resp
			return

		case action.SelectBuffer:
			e.activeBufID = req.BufferID

		case action.CursorMovePast:
			if act.Target == "true" {
				buf.Cursor().SetMovePast(true)
			} else {
				buf.Cursor().SetMovePast(false)
			}

		// Search
		case action.Search:
			res, err := buf.Search(act.Target)
			if err != nil {
				resp.Err = err
				respC <- resp
				return
			}
			resp.Search = res
			resp.ContentChanged = false
		case action.SearchResults:
			resp.Search = buf.SearchResults()
			resp.ContentChanged = false

		// Syntax
		case action.Syntax:
			resp.Syntax = buf.SyntaxResults()
			resp.ContentChanged = false

		// File / buffer
		case action.BufferList:
			for i := range e.bufs {
				resp.Results = append(resp.Results, KeyValue{Key: e.bufs[i].ID(), Value: e.bufs[i].Filename()})
			}
			resp.ContentChanged = false
		case action.NewBuffer:
			resp.BufferID = e.newBuffer()
			resp.NewBuffer = true
		case action.SaveBuffer:
			if err := buf.SaveBuffer(act.Target); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
			resp.ContentChanged = false
		case action.CloseBuffer:
			if buf.Dirty() {
				resp.Err = fmt.Errorf("CloseBuffer: BufferID is dirty")
				respC <- resp
				return
			}
			e.remove(req.BufferID)
			resp.CloseBuffer = true
			resp.ContentChanged = false
		case action.OpenFile:
			id := req.BufferID
			if id == "" {
				resp.NewBuffer = true
				id = e.newBuffer()
				resp.BufferID = id
			}
			buf, err := e.buffer(id)
			if err != nil {
				resp.Err = err
				respC <- resp
				return
			}
			if err := buf.OpenFile(act.Target); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.RenameFile:
			if err := buf.RenameFile(act.Target); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
			resp.InfoChanged = true
			resp.ContentChanged = false
		case action.SaveFileAs:
			buf.SetFilename(act.Target)
			if err := buf.SaveBuffer(""); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
			resp.InfoChanged = true
			resp.ContentChanged = false

		// Move
		case action.Move:
			buf.MoveTo(act.Line, act.Column)
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true
			resp.ContentChanged = false
		case action.MoveObj:
			obj, err := e.objs.Object(act.Target)
			if err != nil {
				resp.Err = err
				respC <- resp
				return
			}
			buf.Move(act.Count, obj)
			resp.Line = buf.Cursor().Line()
			resp.Column = buf.Cursor().Column()
			resp.CursorChanged = true
			resp.ContentChanged = false
		case action.MoveEnd:
			obj, err := e.objs.Object(act.Target)
			if err != nil {
				resp.Err = err
				respC <- resp
				return
			}
			buf.MoveEnd(act.Count, obj)
			resp.CursorChanged = true
			resp.ContentChanged = false
		case action.MovePrev:
			obj, err := e.objs.Object(act.Target)
			if err != nil {
				resp.Err = err
				respC <- resp
				return
			}
			buf.MovePrev(act.Count, obj)
			resp.CursorChanged = true
			resp.ContentChanged = false
		case action.MovePrevEnd:
			obj, err := e.objs.Object(act.Target)
			if err != nil {
				resp.Err = err
				respC <- resp
				return
			}
			if err := buf.MovePrevEnd(act.Count, obj); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
			resp.CursorChanged = true
			resp.ContentChanged = false
		case action.Up:
			buf.Up(act.Count)
			resp.CursorChanged = true
			resp.ContentChanged = false
		case action.Down:
			buf.Down(act.Count)
			resp.CursorChanged = true
			resp.ContentChanged = false
		case action.Prev:
			buf.Prev(act.Count)
			resp.CursorChanged = true
			resp.ContentChanged = false
		case action.Next:
			buf.Next(act.Count)
			resp.CursorChanged = true
			resp.ContentChanged = false
		case action.ScrollDown:
			buf.ScrollDown()
			resp.CursorChanged = true
			resp.ContentChanged = false
		case action.ScrollUp:
			buf.ScrollUp()
			resp.CursorChanged = true
			resp.ContentChanged = false

		// Edit
		case action.DeleteChar:
			err := buf.DeleteChar(act.Line, act.Column, act.Count)
			if err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.DeleteCharBack:
			if err := buf.DeleteCharBack(act.Line, act.Column, act.Count); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.DeleteLine:
			if err := buf.DeleteLine(act.Line, act.Count); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.DeleteObject:
			obj, err := e.objs.Object(act.Target)
			if err != nil {
				resp.Err = err
				respC <- resp
				return
			}
			if err = buf.DeleteObject(act.Line, act.Column, obj, act.Count); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.InsertLine:
			if err := buf.NewLineBelow(act.Line, act.Target, act.Count); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.InsertLineAbove:
			if err := buf.NewLineAbove(act.Line, act.Target, act.Count); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.InsertString:
			if err := buf.InsertString(act.Line, act.Column, act.Target); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.Indent:
			if err := buf.Indent(act.Line, act.Count); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.Outdent:
			if err := buf.Outdent(act.Line, act.Count); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.Content:

		// Undo / redo
		case action.Undo:
			if err := buf.Undo(); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.Redo:
			if err := buf.Redo(); err != nil {
				resp.Err = err
				respC <- resp
				return
			}
		case action.StartGroupUndo:
			buf.StartGroupUndo()
			resp.ContentChanged = false
		case action.StopGroupUndo:
			buf.StopGroupUndo()
			resp.ContentChanged = false
		}
		respC <- getInfo(req, buf, resp)
	}
}

func getInfo(req action.Request, buf buffer.Bufferer, resp Response) Response {
	if buf == nil {
		return resp
	}
	resp.BufferID = buf.ID()
	resp.Dirty = buf.Dirty()
	resp.Filename = buf.Filename()
	resp.Filetype = buf.Filetype()
	resp.Line = buf.Cursor().Line()
	resp.Column = buf.Cursor().Column()
	resp.NumLines = buf.TextStore().NumLines()
	if resp.ContentChanged {
		resp.Content, _ = buf.TextStore().LineRange(req.LineOffset, req.LineCount)
		resp.Syntax = buf.SyntaxResultsRange(req.LineOffset, req.LineCount)
		resp.Search = buf.SearchResults()
	}
	return resp
}

func (e *Editor) updatePos(req action.Request, buf buffer.Bufferer) action.Request {
	if buf == nil {
		return req
	}
	for i := range req.Actions {
		if req.Actions[i].Line == -1 {
			req.Actions[i].Line = buf.Cursor().Line()
		}
		if req.Actions[i].Column == -1 {
			req.Actions[i].Column = buf.Cursor().Column()
		}
		if req.Actions[i].Count < 1 {
			req.Actions[i].Count = 1
		}
	}
	return req
}

func (e *Editor) validateRequest(req action.Request) (buffer.Bufferer, error) {
	var buf buffer.Bufferer
	var err error
	if req.BufferID != "" {
		buf, err = e.buffer(req.BufferID)
		if err != nil {
			return nil, err
		}
	}
	if err := e.actDefs.ValidRequest(req); err != nil {
		return nil, err
	}
	return buf, nil
}
