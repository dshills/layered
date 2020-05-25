package undo

import "github.com/sergi/go-diff/diffmatchpatch"

// ChangeSet is a set of editor changes
type ChangeSet struct {
	changes []Changer
}

// Changes returns the list of changes
func (cs *ChangeSet) Changes() []Changer { return cs.changes }

// RemoveLine will add a line deletion to the set
func (cs *ChangeSet) RemoveLine(ln int) {
	cs.changes = append(cs.changes, &Change{l: ln, t: DeleteLine})
}

// AddLine will add a line add to the set
func (cs *ChangeSet) AddLine(ln int) {
	cs.changes = append(cs.changes, &Change{l: ln, t: AddLine})
}

// ChangeLine will add a line change to the set
func (cs *ChangeSet) ChangeLine(ln int, before, after string) {
	dmp := diffmatchpatch.New()
	patches := dmp.PatchMake(after, before)
	cs.changes = append(cs.changes, &Change{l: ln, t: ChangeLine, p: patches})
}

// New will return a new ChangeSet
func New() ChangeSetter {
	return &ChangeSet{}
}
