package cursor

import (
	"github.com/dshills/layered/textstore"
)

// BufCursor is a window cursor
type BufCursor struct {
	line                          int
	col                           int
	targetCol                     int
	movePast                      bool
	txt                           textstore.TextStorer
	trackLineStart, trackColStart int
	trackLineEnd, trackColEnd     int
}

// AsRange returns line, col as an int array
func (c *BufCursor) AsRange() []int {
	return []int{c.line, c.col}
}

// Line will return the current line
func (c BufCursor) Line() int { return c.line }

// Column will return the current column
func (c BufCursor) Column() int { return c.col }

// GotoLine will move the BufCursor to the specified line
func (c *BufCursor) GotoLine(ln int) bool {
	return c.MoveValid(ln, c.targetCol)
}

// Top will move the BufCursor to 0, 0
func (c *BufCursor) Top() bool {
	return c.MoveValid(0, c.targetCol)
}

// Bottom will move to the last line in the buffer
func (c *BufCursor) Bottom() bool {
	return c.MoveValid(c.txt.NumLines()-1, c.targetCol)
}

// Get will return the current line and column
func (c BufCursor) Get() (int, int) {
	return c.line, c.col
}

// Down moves the BufCursor down cnt lines
func (c *BufCursor) Down(cnt int) bool {
	return c.MoveValid(c.line+cnt, c.targetCol)
}

// Up moves the BufCursor up cnt lines
func (c *BufCursor) Up(cnt int) bool {
	return c.MoveValid(c.line-cnt, c.targetCol)
}

// Prev moves the BufCursor back cnt chars
func (c *BufCursor) Prev(cnt int) bool {
	return c.MoveValid(c.line, c.col-cnt)
}

// Next moves the BufCursor forward cnt chars
func (c *BufCursor) Next(cnt int) bool {
	return c.MoveValid(c.line, c.col+cnt)
}

// SetMovePast will allow cursor to move past the end of the line
func (c *BufCursor) SetMovePast(on bool) {
	c.movePast = on
}

// MovePast will return the move past flag
func (c *BufCursor) MovePast() bool { return c.movePast }

// MoveValid will move the BufCursor to line, col insuring it is a valid position
func (c *BufCursor) MoveValid(line, col int) bool {
	nl := c.txt.NumLines()
	c.targetCol = col
	switch {
	case line < 0:
		line = 0
	case line >= nl:
		line = nl - 1
		if line < 0 {
			line = 0
		}
	}

	cl := c.txt.LineLen(line)
	switch {
	case col < 0:
		col = 0
	case c.movePast:
		if col > cl {
			col = cl
		}
	case cl == 0:
		col = 0
	case col >= cl:
		col = cl - 1
	}

	changed := false
	if line != c.line {
		c.line = line
		changed = true
	}
	if col != c.col {
		c.col = col
		changed = true
	}
	return changed
}

// StartTrack will save the current position
func (c *BufCursor) StartTrack() {
	c.trackLineEnd = -1
	c.trackColEnd = -1
	c.trackLineStart = c.line
	c.trackColStart = c.col
}

// EndTrack will save the ending position
func (c *BufCursor) EndTrack() {
	c.trackLineEnd = c.line
	c.trackColEnd = c.col
}

// Tracked will return the start and end position
func (c *BufCursor) Tracked() [][]int {
	return [][]int{
		[]int{c.trackLineStart, c.trackColStart},
		[]int{c.trackLineEnd, c.trackColEnd},
	}
}

// New will return a new cursor
func New(txt textstore.TextStorer) Cursor {
	return &BufCursor{txt: txt}
}
