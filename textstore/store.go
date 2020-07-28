package textstore

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"io"
	"strings"

	"github.com/dshills/layered/undo"
)

// Store is a TextStore
type Store struct {
	lines     []*line
	delim     string
	undoFac   undo.Factory
	undo      []undo.ChangeSetter
	redo      []undo.ChangeSetter
	active    []undo.ChangeSetter
	grpUndo   bool
	isUndoing bool
	subs      []subscription
	txthash   uint64
}

// --------  edits

// BeginEdit will start edit tracking
func (s *Store) BeginEdit() {
	if s.grpUndo {
		return
	}
	s.active = []undo.ChangeSetter{}
	s.grpUndo = true
}

// EndEdit will save undos since begin
func (s *Store) EndEdit() {
	if !s.grpUndo {
		return
	}
	s.grpUndo = false
	cs := undo.New()
	for i := range s.active {
		cs.AddChanges(s.active[i].Changes()...)
	}
	s.active = []undo.ChangeSetter{}
	s.AddUndoSet(cs)
}

// AddUndoSet will add an undo change set to the store
func (s *Store) AddUndoSet(cs undo.ChangeSetter) {
	s.notifyChange()
	switch {
	case s.grpUndo:
		s.active = append(s.active, cs)
	case s.isUndoing:
		s.redo = append(s.redo, cs)
	default:
		s.undo = append(s.undo, cs)
	}
}

// ReadFrom will read from an io.Reader
func (s *Store) ReadFrom(r io.Reader) (int64, error) {
	nw := false
	if len(s.lines) == 0 {
		nw = true
	}
	scanner := bufio.NewScanner(r)
	s.lines = []*line{}
	for scanner.Scan() {
		s.lines = append(s.lines, newLine(scanner.Text()))
	}
	if nw {
		s.hash()
	} else {
		s.notifyChange()
	}
	return int64(s.Len()), scanner.Err()
}

// Replace replaces text at ln, col
func (s *Store) Replace(ln, from, to int, st string) error {
	if ln < 0 || ln >= len(s.lines) {
		return fmt.Errorf("Replace: line %v out of range", ln)
	}
	s.lines[ln].replaceString(from, to, st)
	s.AddUndoSet(s.lines[ln].changeSet(ln, s.undoFac))
	return nil
}

// Insert inserts text to line
func (s *Store) Insert(ln, col int, st string) error {
	if ln < 0 || ln >= len(s.lines) {
		return fmt.Errorf("Insert: line %v out of range", ln)
	}
	s.lines[ln].insertString(col, st)
	s.AddUndoSet(s.lines[ln].changeSet(ln, s.undoFac))
	return nil
}

// Delete will delete text at line, col
func (s *Store) Delete(ln, col, cnt int) error {
	if ln < 0 || ln >= len(s.lines) {
		return fmt.Errorf("Delete: line %v out of range", ln)
	}
	s.lines[ln].delete(col, cnt)
	s.AddUndoSet(s.lines[ln].changeSet(ln, s.undoFac))
	return nil
}

// NewLine creates a new line after line
func (s *Store) NewLine(ln int, st string) {
	if ln <= 0 {
		s.lines = append([]*line{newLine(st)}, s.lines...)
		return
	}
	if ln >= len(s.lines) {
		s.lines = append(s.lines, newLine(st))
		return
	}
	s.lines = append(s.lines[:ln], append([]*line{newLine(st)}, s.lines[ln:]...)...)
	cs := s.undoFac()
	cs.AddLine(ln)
	cs.ChangeLine(ln+1, "", st)
	s.AddUndoSet(cs)
	return
}

// Reset will set the Store to s
func (s *Store) Reset(st string) uint64 {
	r := strings.NewReader(st)
	s.ReadFrom(r)
	return s.txthash
}

// DeleteLine will remove a line
func (s *Store) DeleteLine(line int) (string, error) {
	if line < 0 || line >= len(s.lines) {
		return "", fmt.Errorf("newLine: Invalid line %v", line)
	}
	original := s.lines[line].String()
	if line < len(s.lines)-1 {
		copy(s.lines[line:], s.lines[line+1:])
	}
	s.lines[len(s.lines)-1] = nil // or the zero value of T
	s.lines = s.lines[:len(s.lines)-1]
	cs := s.undoFac()
	cs.ChangeLine(line, original, "")
	cs.RemoveLine(line)
	s.AddUndoSet(cs)
	return original, nil
}

// ResetLine will set the contents of a line
func (s *Store) ResetLine(line int, st string) (string, error) {
	if line < 0 || line >= len(s.lines) {
		return "", fmt.Errorf("ResetLine: Invalid offset %v", line)
	}
	original := s.lines[line].String()
	s.lines[line].reset(st)
	cs := s.undoFac()
	cs.ChangeLine(line, original, st)
	s.AddUndoSet(cs)
	return original, nil
}

// ------------ end edits

func (s *Store) String() string {
	builder := strings.Builder{}
	for _, ln := range s.lines {
		builder.WriteString(ln.delimited(s.delim))
	}
	return builder.String()
}

func (s *Store) notifyChange() {
	if len(s.subs) == 0 || !s.hash() {
		return
	}
	go func() {
		for i := range s.subs {
			s.subs[i].ch <- s.Hash64()
		}
	}()
}

func (s *Store) hash() bool {
	h := fnv.New64a()
	h.Write([]byte(s.String()))
	nh := h.Sum64()
	if s.txthash != nh {
		s.txthash = nh
		return true
	}
	return false
}

// Hash64 will return a 64bit hash
func (s *Store) Hash64() uint64 {
	return s.txthash
}

// Undo will undo the last set of edits
func (s *Store) Undo() error {
	if len(s.undo) == 0 {
		return fmt.Errorf("No more undo entries")
	}
	s.isUndoing = true
	defer func() {
		s.isUndoing = false
	}()
	var x undo.ChangeSetter
	x, s.undo = s.undo[len(s.undo)-1], s.undo[:len(s.undo)-1]
	changes := x.Changes()

	for i := len(changes) - 1; i >= 0; i-- {
		change := changes[i]
		switch change.Type() {
		case undo.DeleteLine:
			s.NewLine(change.Line(), change.Undo(""))

		case undo.AddLine:
			s.DeleteLine(change.Line())

		case undo.ChangeLine:
			txt, _ := s.LineString(change.Line())
			s.ResetLine(change.Line(), change.Undo(txt))
		}
	}
	return nil
}

// Redo will undo an undo
func (s *Store) Redo() error {
	if len(s.redo) == 0 {
		return fmt.Errorf("No more redo entries")
	}
	var x undo.ChangeSetter
	x, s.redo = s.redo[len(s.redo)-1], s.redo[:len(s.redo)-1]
	changes := x.Changes()

	for i := len(changes) - 1; i >= 0; i-- {
		change := changes[i]
		switch change.Type() {
		case undo.DeleteLine:
			s.NewLine(change.Line(), change.Undo(""))

		case undo.AddLine:
			s.DeleteLine(change.Line())

		case undo.ChangeLine:
			txt, _ := s.LineString(change.Line())
			s.ResetLine(change.Line(), change.Undo(txt))
		}
	}
	return nil
}

// SetLineDelim will set the line delimeter
func (s *Store) SetLineDelim(str string) {
	s.delim = str
}

// LineDelim will return the current linedelimeter
func (s *Store) LineDelim() string {
	return s.delim
}

// Len will return the total length with delimeters
func (s *Store) Len() int {
	cnt := 0
	dl := len(s.delim)
	for i := range s.lines {
		cnt += s.lines[i].Len() + dl
	}
	return cnt
}

// WriteTo will write the store to w
func (s *Store) WriteTo(w io.Writer) (int64, error) {
	cnt, err := w.Write([]byte(s.String()))
	return int64(cnt), err
}

// LineRange returns a range of lines
func (s *Store) LineRange(ln, cnt int) ([]string, error) {
	if ln < 0 || ln+cnt >= len(s.lines) {
		return nil, fmt.Errorf("LineRange: line %v out of range", ln)
	}
	strs := []string{}
	for i := 0; i < cnt; i++ {
		strs = append(strs, s.lines[ln+i].String())
	}
	return strs, nil
}

// LineString will return the line as a string
func (s *Store) LineString(line int) (string, error) {
	if line < 0 || line >= len(s.lines) {
		return "", fmt.Errorf("LineString: Invalid offset %v", line)
	}
	return s.lines[line].String(), nil
}

// NumLines returns the number of lines
func (s *Store) NumLines() int {
	return len(s.lines)
}

// LineLen returns the lkength of a line
func (s *Store) LineLen(line int) int {
	if line < 0 || line >= len(s.lines) {
		return -1
	}
	return s.lines[line].Len()
}

// RuneAt returns the rune at ln, col
func (s *Store) RuneAt(ln, col int) (rune, error) {
	if ln < 0 || ln >= len(s.lines) {
		return 0, fmt.Errorf("RuneAt: line %v out of range", ln)
	}
	if col < 0 || col >= s.lines[ln].Len() {
		return 0, fmt.Errorf("RuneAt: column %v out of range", col)
	}
	return s.lines[ln].runeAt(col), nil
}

// Subscribe will subscribe to updates
func (s *Store) Subscribe(id string, up chan uint64) {
	s.subs = append(s.subs, subscription{id: id, ch: up})
}

// Unsubscribe will remove a subscription
func (s *Store) Unsubscribe(id string) {
	for i := range s.subs {
		if s.subs[i].id == id {
			// does not maintain order
			s.subs[i] = s.subs[len(s.subs)-1]      // Copy last element to index i.
			s.subs[len(s.subs)-1] = subscription{} // Erase last element (write zero value).
			s.subs = s.subs[:len(s.subs)-1]        // Truncate slice.
		}
	}
}

type subscription struct {
	id string
	ch chan uint64
}

// New returns a TextStorer
func New(uf undo.Factory) TextStorer {
	return &Store{undoFac: uf, delim: "\n"}
}
