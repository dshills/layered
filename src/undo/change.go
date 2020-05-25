package undo

import "github.com/sergi/go-diff/diffmatchpatch"

// Change is a change to a text store
type Change struct {
	pos []int
	d   bool
	t   int
	l   int
	p   []diffmatchpatch.Patch
}

// Cursor returns the cursor position before the change
func (c *Change) Cursor() []int { return c.pos }

// Dirty returns the dirty status before the change
func (c *Change) Dirty() bool { return c.d }

// Type will return the change type
func (c *Change) Type() int { return c.t }

// Line will return the line the change was made
func (c *Change) Line() int { return c.l }

// Undo will return the text before the change
func (c *Change) Undo(after string) string {
	dmp := diffmatchpatch.New()
	txt, _ := dmp.PatchApply(c.p, after)
	return txt
}
