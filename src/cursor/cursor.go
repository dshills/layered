package cursor

import "github.com/dshills/layered/textstore"

// Cursor is a window cursor
type Cursor struct {
	line      int
	col       int
	targetCol int
	movePast  bool
	txt       textstore.TextStorer
}

// AsRange returns line, col as an int array
func (c *Cursor) AsRange() []int {
	return []int{c.line, c.col}
}

// Line will return the current line
func (c Cursor) Line() int { return c.line }

// Column will return the current column
func (c Cursor) Column() int { return c.col }

// GotoLine will move the Cursor to the specified line
func (c *Cursor) GotoLine(ln int) bool {
	return c.MoveValid(ln, c.targetCol)
}

// Top will move the Cursor to 0, 0
func (c *Cursor) Top() bool {
	return c.MoveValid(0, c.targetCol)
}

// Bottom will move to the last line in the buffer
func (c *Cursor) Bottom() bool {
	return c.MoveValid(c.txt.NumLines()-1, c.targetCol)
}

// Get will return the current line and column
func (c Cursor) Get() (int, int) {
	return c.line, c.col
}

// Down moves the Cursor down cnt lines
func (c *Cursor) Down(cnt int) bool {
	return c.MoveValid(c.line+cnt, c.targetCol)
}

// Up moves the Cursor up cnt lines
func (c *Cursor) Up(cnt int) bool {
	return c.MoveValid(c.line-cnt, c.targetCol)
}

// Prev moves the Cursor back cnt chars
func (c *Cursor) Prev(cnt int) bool {
	return c.MoveValid(c.line, c.col-cnt)
}

// Next moves the Cursor forward cnt chars
func (c *Cursor) Next(cnt int) bool {
	return c.MoveValid(c.line, c.col+cnt)
}

// MoveValid will move the Cursor to line, col insuring it is a valid position
func (c *Cursor) MoveValid(line, col int) bool {
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

// New will return a new cursor
func New(txt textstore.TextStorer) Cursorer {
	return &Cursor{txt: txt}
}
