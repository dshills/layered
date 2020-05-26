package buffer

import (
	"fmt"

	"github.com/dshills/layered/textobject"
)

// DeleteObject will remove the next occurence of an object
func (b *Buffer) DeleteObject(line, col int, obj textobject.TextObjecter) error {
	if line == -1 {
		line = b.cur.Line()
	}
	if col == -1 {
		col = b.cur.Column()
	}
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
		return nil
	}
	return nil
}

// NewLineAbove will add a line above line with string st
func (b *Buffer) NewLineAbove(line int, st string) error {
	if line == -1 {
		line = b.cur.Line()
	}
	if _, err := b.txt.NewLine(st, line-1); err != nil {
		return fmt.Errorf("Buffer.NewLineAbove: %v", err)
	}
	return nil
}

// NewLineBelow will add a line below line with string st
func (b *Buffer) NewLineBelow(line int, st string) error {
	if line == -1 {
		line = b.cur.Line()
	}
	if _, err := b.txt.NewLine(st, line); err != nil {
		return fmt.Errorf("Buffer.NewLineBelow: %v", err)
	}
	return nil
}

// DeleteChar will delete the next char
func (b *Buffer) DeleteChar(line, col int) error {
	if line == -1 {
		line = b.cur.Line()
	}
	if col == -1 {
		col = b.cur.Column()
	}
	lw, err := b.txt.LineWriterAt(line)
	if err != nil {
		return fmt.Errorf("Buffer.DeleteChar %v", err)
	}
	if _, err := lw.ReplaceAt([]byte(""), int64(col), 1); err != nil {
		return fmt.Errorf("Buffer.DeleteChar: %v", err)
	}
	lw.Flush()
	return nil
}

// DeleteCharBack will delete the prev char
func (b *Buffer) DeleteCharBack(line, col int) error {
	if line == -1 {
		line = b.cur.Line()
	}
	if col == -1 {
		col = b.cur.Column()
	}
	lw, err := b.txt.LineWriterAt(line)
	if err != nil {
		return fmt.Errorf("Buffer.DeleteCharBack %v", err)
	}
	b.cur.Prev(1)
	if _, err := lw.ReplaceAt([]byte(""), int64(col), 1); err != nil {
		return fmt.Errorf("Buffer.DeleteCharBack: %v", err)
	}
	lw.Flush()
	return nil
}

// InsertString will insert a string at line, col
func (b *Buffer) InsertString(line, col int, st string) error {
	if line == -1 {
		line = b.cur.Line()
	}
	if col == -1 {
		col = b.cur.Column()
	}
	lw, err := b.txt.LineWriterAt(line)
	if err != nil {
		return fmt.Errorf("Buffer.InsertString %v", err)
	}
	if _, err := lw.InsertAt([]byte(st), int64(col)); err != nil {
		return fmt.Errorf("Buffer.InsertString: %v", err)
	}
	b.cur.Next(len(st))
	return nil
}

// DeleteLine will remove a line
func (b *Buffer) DeleteLine(line int) error {
	if line == -1 {
		line = b.cur.Line()
	}
	_, err := b.txt.DeleteLine(line)
	return err
}

// Indent will indent the current line
func (b *Buffer) Indent(line, cnt int) error {
	if line == -1 {
		line = b.cur.Line()
	}
	if cnt == 0 {
		cnt = 1
	}
	for i := line; i <= line+cnt; i++ {
		lw, err := b.txt.LineWriterAt(i)
		if err != nil {
			return fmt.Errorf("Buffer.Indent %v", err)
		}
		if _, err := lw.InsertRuneAt('\t', int64(0)); err != nil {
			return fmt.Errorf("Buffer.Indent: %v", err)
		}
	}
	return nil
}

// Outdent will decrease the indent level
func (b *Buffer) Outdent(line, cnt int) error {
	if line == -1 {
		line = b.cur.Line()
	}
	if cnt == 0 {
		cnt = 1
	}
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
