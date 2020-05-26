package editor

import (
	"fmt"

	"github.com/dshills/layered/action"
)

// Exec will execute a transaction in the editor
func (e *Editor) Exec(tr action.Transactioner) error {
	acts := tr.Actions()
	for i := range acts {
		switch acts[i].Name() {
		// File / buffer
		case action.NewBuffer:
			e.newBuffer()
		case action.SaveBuffer:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			if err := buf.SaveBuffer(acts[i].Param()); err != nil {
				return err
			}
		case action.CloseBuffer:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			if buf.Dirty() {
				return fmt.Errorf("CloseBuffer: Buffer is dirty")
			}
			e.removeBuffer(tr.Buffer())
		case action.OpenFile:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			if err := buf.OpenFile(acts[i].Param()); err != nil {
				return err
			}
		case action.RenameFile:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			if err := buf.RenameFile(acts[i].Param()); err != nil {
				return err
			}
		// Move
		case action.Move:
			obj, err := e.objs.Object(acts[i].Object())
			if err != nil {
				return err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.Move(1, obj)
		case action.MoveEnd:
			obj, err := e.objs.Object(acts[i].Object())
			if err != nil {
				return err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.MoveEnd(1, obj)
		case action.MovePrev:
			obj, err := e.objs.Object(acts[i].Object())
			if err != nil {
				return err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.MovePrev(1, obj)
		case action.MovePrevEnd:
			obj, err := e.objs.Object(acts[i].Object())
			if err != nil {
				return err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			if err := buf.MovePrevEnd(1, obj); err != nil {
				return err
			}
		case action.Up:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.Up(1)
		case action.Down:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.Down(1)
		case action.Prev:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.Prev(1)
		case action.Next:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.Next(1)
		case action.ScrollDown:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.ScrollDown()
		case action.ScrollUp:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.ScrollUp()

		// Edit
		case action.DeleteChar:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteChar(acts[i].Line(), acts[i].Column(), acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return err
			}
		case action.DeleteCharBack:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteCharBack(acts[i].Line(), acts[i].Column(), acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return err
			}
		case action.DeleteLine:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteLine(acts[i].Line(), acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return err
			}
		case action.DeleteObject:
			obj, err := e.objs.Object(acts[i].Object())
			if err != nil {
				return err
			}
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.DeleteObject(acts[i].Line(), acts[i].Column(), obj, acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return err
			}
		case action.InsertLine:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.NewLineBelow(acts[i].Line(), acts[i].Param(), acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return err
			}
		case action.InsertLineAbove:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			buf.TextStore().StartGroupUndo()
			err = buf.NewLineAbove(acts[i].Line(), acts[i].Param(), acts[i].Count())
			buf.TextStore().StopGroupUndo()
			if err != nil {
				return err
			}
		case action.InsertString:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			if err := buf.InsertString(acts[i].Line(), acts[i].Column(), acts[i].Param()); err != nil {
				return err
			}
		case action.Indent:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			if err := buf.Indent(acts[i].Line(), acts[i].Count()); err != nil {
				return err
			}
		case action.Outdent:
			buf, err := e.Buffer(tr.Buffer())
			if err != nil {
				return err
			}
			if err := buf.Outdent(acts[i].Line(), acts[i].Count()); err != nil {
				return err
			}
		}

	}
	return nil
}
