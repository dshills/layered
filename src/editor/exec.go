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
}

// Exec will execute a transaction in the editor
func (e *Editor) Exec(tr action.Transactioner) (resp *Response, err error) {
	acts := tr.Actions()
	resp = &Response{Buffer: tr.Buffer()}
	for i := range acts {
		switch acts[i].Name() {

		case action.Syntax:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
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
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			if err := buf.SaveBuffer(acts[i].Param()); err != nil {
				return nil, err
			}
		case action.CloseBuffer:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			if buf.Dirty() {
				return nil, fmt.Errorf("CloseBuffer: Buffer is dirty")
			}
			e.removeBuffer(tr.Buffer())
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
			if err := buf.OpenFile(acts[i].Param()); err != nil {
				return nil, err
			}
		case action.RenameFile:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			if err := buf.RenameFile(acts[i].Param()); err != nil {
				return nil, err
			}
		case action.SaveFileAs:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.SetFilename(acts[i].Param())
			if err := buf.SaveBuffer(""); err != nil {
				return nil, err
			}
		// Move
		case action.Move:
			obj, err := e.objs.Object(acts[i].Object())
			if err != nil {
				return nil, err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.Move(1, obj)
		case action.MoveEnd:
			obj, err := e.objs.Object(acts[i].Object())
			if err != nil {
				return nil, err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.MoveEnd(1, obj)
		case action.MovePrev:
			obj, err := e.objs.Object(acts[i].Object())
			if err != nil {
				return nil, err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.MovePrev(1, obj)
		case action.MovePrevEnd:
			obj, err := e.objs.Object(acts[i].Object())
			if err != nil {
				return nil, err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			if err := buf.MovePrevEnd(1, obj); err != nil {
				return nil, err
			}
		case action.Up:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.Up(1)
		case action.Down:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.Down(1)
		case action.Prev:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.Prev(1)
		case action.Next:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.Next(1)
		case action.ScrollDown:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.ScrollDown()
		case action.ScrollUp:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.ScrollUp()

		// Edit
		case action.DeleteChar:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteChar(acts[i].Line(), acts[i].Column(), acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.DeleteCharBack:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteCharBack(acts[i].Line(), acts[i].Column(), acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.DeleteLine:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteLine(acts[i].Line(), acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.DeleteObject:
			obj, err := e.objs.Object(acts[i].Object())
			if err != nil {
				return nil, err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteObject(acts[i].Line(), acts[i].Column(), obj, acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.InsertLine:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.NewLineBelow(acts[i].Line(), acts[i].Param(), acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.InsertLineAbove:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.NewLineAbove(acts[i].Line(), acts[i].Param(), acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return nil, err
			}
		case action.InsertString:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			if err := buf.InsertString(acts[i].Line(), acts[i].Column(), acts[i].Param()); err != nil {
				return nil, err
			}
		case action.Indent:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			if err := buf.Indent(acts[i].Line(), acts[i].Count()); err != nil {
				return nil, err
			}
		case action.Outdent:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return nil, err
			}
			if err := buf.Outdent(acts[i].Line(), acts[i].Count()); err != nil {
				return nil, err
			}
		case action.Content:
			var buf buffer.Bufferer
			buf, err = e.Buffer(tr.Buffer())
			if err != nil {
				return
			}
			resp.Content, err = buf.TextStore().LineRangeString(acts[i].Line(), acts[i].Count())
			return
		}

	}
	return
}
