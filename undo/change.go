package undo

import "github.com/sergi/go-diff/diffmatchpatch"

// Change is a change to a text store
type Change struct {
	pos []int
	d   bool
	t   ChangeType
	l   int
	p   []diffmatchpatch.Patch
}

// Cursor returns the cursor position before the change
func (c *Change) Cursor() []int { return c.pos }

// Dirty returns the dirty status before the change
func (c *Change) Dirty() bool { return c.d }

// Type will return the change type
func (c *Change) Type() ChangeType { return c.t }

// Line will return the line the change was made
func (c *Change) Line() int { return c.l }

// Undo will return the text before the change
func (c *Change) Undo(after string) string {
	dmp := diffmatchpatch.New()
	txt, _ := dmp.PatchApply(c.p, after)
	return txt
}

// SetLine will set the line the change occured
func (c *Change) SetLine(l int) {
	c.l = l
}

// SetType will set the change type
func (c *Change) SetType(t ChangeType) {
	c.t = t
}

// SetCursor will set the cursor position
func (c *Change) SetCursor(cur []int) {
	c.pos = cur
}

// SetDirty will set the dirty flag
func (c *Change) SetDirty(di bool) {
	c.d = di
}

// GenChange will create a diff
func (c *Change) GenChange(before, after string) {
	dmp := diffmatchpatch.New()
	c.p = dmp.PatchMake(after, before)
}

// NewChange will return a Changer
func NewChange(before, after string) Changer {
	c := Change{}
	c.GenChange(before, after)
	return &c
}
