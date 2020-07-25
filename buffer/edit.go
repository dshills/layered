package buffer

import (
	"fmt"

	"github.com/dshills/layered/logger"
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
		lw, err := b.txt.LineWriterAt(i)
		if err != nil {
			return fmt.Errorf("Buffer.DeleteObject %v", err)
		}
		if _, err := lw.ReplaceAt([]byte(s), int64(col), int64(matches[0][1])); err != nil {
			return fmt.Errorf("Buffer.DeleteChar: %v", err)
		}
		if cnt > 1 {
			b.ReplaceObject(i, col, obj, s, cnt-1)
		}
		lw.Flush()
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
		lw, err := b.txt.LineWriterAt(i)
		if err != nil {
			return fmt.Errorf("Buffer.DeleteObject %v", err)
		}
		if _, err := lw.ReplaceAt([]byte(""), int64(col), int64(matches[0][1])); err != nil {
			return fmt.Errorf("Buffer.DeleteChar: %v", err)
		}
		if cnt > 1 {
			b.DeleteObject(i, col, obj, cnt-1)
		}
		lw.Flush()
		return nil
	}
	return nil
}

// NewLineBelow will add a line below line with string st
func (b *Buffer) NewLineBelow(line int, st string, cnt int) error {
	if _, err := b.txt.NewLine(st, line+1); err != nil {
		return fmt.Errorf("Buffer.NewLineBelow: %v", err)
	}
	b.cur.MoveValid(b.cur.Line()+1, 0)
	if cnt > 1 {
		b.NewLineBelow(line, st, cnt-1)
	}
	return nil
}

// NewLineAbove will add a line above line with string st
func (b *Buffer) NewLineAbove(line int, st string, cnt int) error {
	if _, err := b.txt.NewLine(st, line); err != nil {
		return fmt.Errorf("Buffer.NewLineAbove: %v", err)
	}
	b.cur.MoveValid(b.cur.Line(), 0)
	if cnt > 1 {
		b.NewLineAbove(line, st, cnt-1)
	}
	return nil
}

// DeleteChar will delete the next char
func (b *Buffer) DeleteChar(line, col, cnt int) error {
	lw, err := b.txt.LineWriterAt(line)
	if err != nil {
		return fmt.Errorf("Buffer.DeleteChar %v", err)
	}
	if _, err := lw.ReplaceAt([]byte(""), int64(col), 1); err != nil {
		return fmt.Errorf("Buffer.DeleteChar: %v", err)
	}
	lw.Flush()
	if cnt > 1 {
		b.DeleteChar(line, col, cnt-1)
	}
	return nil
}

// DeleteCharBack will delete the prev char
func (b *Buffer) DeleteCharBack(line, col, cnt int) error {
	lw, err := b.txt.LineWriterAt(line)
	if err != nil {
		return fmt.Errorf("Buffer.DeleteCharBack %v", err)
	}
	b.cur.Prev(1)
	if _, err := lw.ReplaceAt([]byte(""), int64(col), 1); err != nil {
		return fmt.Errorf("Buffer.DeleteCharBack: %v", err)
	}
	lw.Flush()
	if cnt > 1 {
		b.DeleteCharBack(line, col, cnt-1)
	}
	return nil
}

// InsertString will insert a string at line, col
func (b *Buffer) InsertString(line, col int, st string) error {
	lw, err := b.txt.LineWriterAt(line)
	if err != nil {
		return fmt.Errorf("Buffer.InsertString %v", err)
	}
	if _, err := lw.InsertAt([]byte(st), int64(col)); err != nil {
		return fmt.Errorf("Buffer.InsertString: %v", err)
	}
	lw.Flush()
	b.cur.Next(len(st))
	logger.Debugf("InsertString: Pos: %v:%v, %v, After %v:%v CursorPast: %v", line, col, st, b.cur.Line(), b.cur.Column(), b.cur.MovePast())
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
	for i := line; i <= line+cnt; i++ {
		lw, err := b.txt.LineWriterAt(i)
		if err != nil {
			return fmt.Errorf("Buffer.Indent %v", err)
		}
		if _, err := lw.InsertRuneAt('\t', int64(0)); err != nil {
			return fmt.Errorf("Buffer.Indent: %v", err)
		}
		lw.Flush()
	}
	return nil
}

// Outdent will decrease the indent level
func (b *Buffer) Outdent(line, cnt int) error {
	for i := line; i <= line+cnt; i++ {
		r, _, err := b.txt.ReadRuneAt(b.cur.Line(), 0)
		if err != nil {
			return fmt.Errorf("Buffer.Outdent %v", err)
		}
		if r == '\t' {
			lw, err := b.txt.LineWriterAt(i)
			if err != nil {
				return fmt.Errorf("buffer.Outdent %v", err)
			}
			if _, err := lw.ReplaceAt([]byte(""), int64(b.cur.Line()), 1); err != nil {
				return fmt.Errorf("buffer.Outdent %v", err)
			}
			lw.Flush()
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
	b.txt.StartGroupUndo()
}

// StopGroupUndo will stop grouping undos
func (b *Buffer) StopGroupUndo() {
	b.txt.StopGroupUndo()
}
