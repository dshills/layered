package buffer

import (
	"fmt"
	"strings"

	"github.com/dshills/layered/textobject"
)

// Reset will reset the buffer content
func (b *Buffer) Reset(st string) {
	b.txthash = b.txt.Reset(st)
}

// ReplaceObject will replace an object with s
func (b *Buffer) ReplaceObject(line, col int, obj textobject.TextObjecter, s string, cnt int) error {
	for i := line; i < b.txt.NumLines(); i++ {
		st, err := b.txt.LineString(i)
		if err != nil {
			return fmt.Errorf("Buffer.DeleteObject: %v", err)
		}
		matches := obj.FindAfter(st, col)
		col = -1
		if len(matches) == 0 {
			continue
		}
		b.txt.Replace(i, col, matches[0][1], s)
		if cnt > 1 {
			b.ReplaceObject(i, col, obj, s, cnt-1)
		}
		return nil
	}
	return nil

}

// DeleteObject will remove the next occurence of an object
func (b *Buffer) DeleteObject(line, col int, obj textobject.TextObjecter, cnt int) error {
	for i := line; i < b.txt.NumLines(); i++ {
		st, err := b.txt.LineString(i)
		if err != nil {
			return fmt.Errorf("Buffer.DeleteObject: %v", err)
		}
		matches := obj.FindAfter(st, col)
		col = -1
		if len(matches) == 0 {
			continue
		}
		b.txt.Delete(i, col, matches[0][1])
		if cnt > 1 {
			b.DeleteObject(i, col, obj, cnt-1)
		}
		return nil
	}
	return nil
}

// NewLineBelow will add a line below line with string st
func (b *Buffer) NewLineBelow(line int, st string, cnt int) error {
	b.txt.NewLine(line+1, st)
	b.cur.MoveValid(b.cur.Line()+1, 0)
	if cnt > 1 {
		b.NewLineBelow(line, st, cnt-1)
	}
	return nil
}

// NewLineAbove will add a line above line with string st
func (b *Buffer) NewLineAbove(line int, st string, cnt int) error {
	b.txt.NewLine(line, st)
	b.cur.MoveValid(b.cur.Line(), 0)
	if cnt > 1 {
		b.NewLineAbove(line, st, cnt-1)
	}
	return nil
}

// DeleteChar will delete the next char
func (b *Buffer) DeleteChar(line, col, cnt int) error {
	b.txt.Delete(line, col, cnt)
	return nil
}

// DeleteCharBack will delete the prev char
func (b *Buffer) DeleteCharBack(line, col, cnt int) error {
	b.cur.Prev(1)
	b.txt.Delete(line, col, 1)
	if cnt > 1 {
		b.DeleteCharBack(line, col, cnt-1)
	}
	return nil
}

// InsertString will insert a string at line, col
func (b *Buffer) InsertString(line, col int, st string) error {
	b.txt.Insert(line, col, st)
	b.cur.Next(len(st))
	return nil
}

// DeleteLine will remove a line
func (b *Buffer) DeleteLine(line, cnt int) error {
	_, err := b.txt.DeleteLine(line)
	if err != nil {
		return fmt.Errorf("Buffer.DeleteLine: %v", err)
	}
	if cnt > 1 {
		b.DeleteLine(line, cnt-1)
	}
	return err
}

// Indent will indent the current line
func (b *Buffer) Indent(line, cnt int) error {
	st := strings.Repeat("\t", cnt)
	for i := line; i <= line+cnt; i++ {
		b.txt.Insert(i, 0, st)
	}
	return nil
}

// Outdent will decrease the indent level
func (b *Buffer) Outdent(line, cnt int) error {
	for i := line; i <= line+cnt; i++ {
		r, _ := b.txt.RuneAt(b.cur.Line(), 0)
		if r == '\t' {
			b.txt.Delete(line, 0, 1)
		}
	}
	return nil
}

// Undo will undo the last edit
func (b *Buffer) Undo() error {
	return b.txt.Undo()
}

// Redo will redo the last edit
func (b *Buffer) Redo() error {
	return b.txt.Redo()
}

// StartGroupUndo will group edits into a single undo
func (b *Buffer) StartGroupUndo() {
	b.txt.BeginEdit()
}

// StopGroupUndo will stop grouping undos
func (b *Buffer) StopGroupUndo() {
	b.txt.EndEdit()
}
