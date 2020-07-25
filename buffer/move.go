package buffer

import (
	"fmt"

	"github.com/dshills/layered/logger"
	"github.com/dshills/layered/textobject"
)

const scrollDistance = 35

// Position will return the current cursor position
func (b *Buffer) Position() []int {
	return []int{b.cur.Line(), b.cur.Column()}
}

// Move will move to the next cnt objects
func (b *Buffer) Move(cnt int, obj textobject.TextObjecter) error {
	col := b.cur.Column()
	nl := b.txt.NumLines()
	logger.Debugf("Move: %v Line %v Col %v", obj.Name(), b.cur.Line(), col)
	for i := b.cur.Line(); i < nl; i++ {
		str, err := b.txt.LineString(i)
		if err != nil {
			return fmt.Errorf("Buffer.Move: %v", err)
		}
		mtch := obj.FindAfter(str, col)
		if len(mtch) > 0 {
			logger.Debugf("Move: Line %v Matches %v", i, mtch)
			b.cur.MoveValid(i, mtch[0][0])
			cnt--
			if cnt == 0 {
				return nil
			}
		}
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
		mtch := obj.FindAfter(str, col)
		if len(mtch) > 0 {
			logger.Debugf("MoveEnd: Line %v Matches %v", i, mtch)
			b.cur.MoveValid(i, mtch[0][1])
			cnt--
			if cnt == 0 {
				return nil
			}
		}
		col = -1
	}
	return fmt.Errorf("Buffer.MoveEnd: Not found")
}

// MovePrev will move cnt prev objects
func (b *Buffer) MovePrev(cnt int, obj textobject.TextObjecter) error {
	col := b.cur.Column()
	ln := b.cur.Line()
	for i := ln; i >= 0; i-- {
		str, err := b.txt.LineString(i)
		if err != nil {
			return fmt.Errorf("Buffer.MovePrev: %v", err)
		}
		if i != ln {
			col = len(str) + 1
		}
		matches := obj.FindBefore(str, col)
		logger.Debugf("MovePrev: Line %v Matches %v", i, matches)
		if len(matches) == 0 {
			continue
		}
		m := matches[len(matches)-1]
		b.cur.MoveValid(i, m[0])
		cnt--
		if cnt == 0 {
			return nil
		}
	}
	return fmt.Errorf("Buffer.MovePrev: Not found")
}

// MovePrevEnd will move cnt previous objects at end
func (b *Buffer) MovePrevEnd(cnt int, obj textobject.TextObjecter) error {
	col := b.cur.Column()
	ln := b.cur.Line()
	for i := ln; i >= 0; i-- {
		str, err := b.txt.LineString(i)
		if err != nil {
			return fmt.Errorf("Buffer.MovePrev: %v", err)
		}
		if i != ln {
			col = len(str) + 1
		}
		matches := obj.FindBefore(str, col)
		logger.Debugf("MovePrevEnd: Line %v Matches %v", i, matches)
		if len(matches) == 0 {
			continue
		}
		m := matches[len(matches)-1]
		b.cur.MoveValid(i, m[1])
		cnt--
		if cnt == 0 {
			return nil
		}
	}
	return fmt.Errorf("Buffer.MovePrevEnd: Not found")
}

// MoveTo will move to a line and column
func (b *Buffer) MoveTo(line, col int) error {
	b.cur.MoveValid(line, col)
	return nil
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

// ScrollDown will scroll the cursor down
func (b *Buffer) ScrollDown() {
	b.cur.MoveValid(b.cur.Line()+scrollDistance, b.cur.Column())
}

// ScrollUp  will scroll the cursor up
func (b *Buffer) ScrollUp() {
	b.cur.MoveValid(b.cur.Line()-scrollDistance, b.cur.Column())
}

// BeginSelect will save the current cursor position
func (b *Buffer) BeginSelect() {
	b.cur.StartTrack()
}

// EndSelect will save the current position
func (b *Buffer) EndSelect() {
	b.cur.EndTrack()
}

// Selection will return the cursor's selection
func (b *Buffer) Selection() [][]int {
	return b.cur.Tracked()
}
