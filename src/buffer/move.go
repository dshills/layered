package buffer

import (
	"fmt"

	"github.com/dshills/layered/textobject"
)

// Move will move to the next cnt objects
func (b *Buffer) Move(cnt int, obj textobject.TextObjecter) error {
	col := b.cur.Column()
	nl := b.txt.NumLines()
	for i := b.cur.Line(); i < nl; i++ {
		str, err := b.txt.LineString(i)
		if err != nil {
			return fmt.Errorf("Buffer.Move: %v", err)
		}
		obj.FindAfter(str, col)
		col = -1
	}
	return fmt.Errorf("Buffer.Move: Not found")
}

// MoveEnd will move cnt object end
func (b *Buffer) MoveEnd(cnt int, obj textobject.TextObjecter) error {
	col := b.cur.Column()
	nl := b.txt.NumLines()
	for i := b.cur.Line(); i < nl; i++ {
		str, err := b.txt.LineString(i)
		if err != nil {
			return fmt.Errorf("Buffer.MoveEnd: %v", err)
		}
		obj.FindAfter(str, col)
		col = -1
	}
	return fmt.Errorf("Buffer.MoveEnd: Not found")
}

// MovePrev will move cnt prev objects
func (b *Buffer) MovePrev(cnt int, obj textobject.TextObjecter) error {
	col := b.cur.Column()
	nl := b.txt.NumLines()
	for i := b.cur.Line(); i < nl; i++ {
		str, err := b.txt.LineString(i)
		if err != nil {
			return fmt.Errorf("Buffer.MovePrev: %v", err)
		}
		obj.FindAfter(str, col)
		col = -1
	}
	return fmt.Errorf("Buffer.MovePrev: Not found")
}

// MovePrevEnd will move cnt previous objects at end
func (b *Buffer) MovePrevEnd(cnt int, obj textobject.TextObjecter) error {
	col := b.cur.Column()
	nl := b.txt.NumLines()
	for i := b.cur.Line(); i < nl; i++ {
		str, err := b.txt.LineString(i)
		if err != nil {
			return fmt.Errorf("Buffer.MovePrev: %v", err)
		}
		obj.FindAfter(str, col)
		col = -1
	}
	return fmt.Errorf("Buffer.MovePrev: Not found")
}

// Up will move the cursor up cnt
func (b *Buffer) Up(cnt int) {
	b.cur.Up(cnt)
}

// Down will move the cursor down cnt
func (b *Buffer) Down(cnt int) {
	b.cur.Down(cnt)
}

// Prev will move the curosr back by cnt
func (b *Buffer) Prev(cnt int) {
	b.cur.Prev(cnt)
}

// Next will move the cursor forward by cnt
func (b *Buffer) Next(cnt int) {
	b.cur.Next(cnt)
}
