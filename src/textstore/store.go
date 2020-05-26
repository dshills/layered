package textstore

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/dshills/layered/undo"
)

// Store is an implementation of LineBuffer
type Store struct {
	lines     []*Line
	delim     string
	undoFac   undo.Factory
	undo      []undo.ChangeSetter
	redo      []undo.ChangeSetter
	hold      []undo.ChangeSetter
	grpUndo   bool
	isUndoing bool
}

// Reset will set the Store to s
func (s *Store) Reset(st string) {
	s.lines = []*Line{}
	scanner := bufio.NewScanner(strings.NewReader(st))
	for scanner.Scan() {
		s.lines = append(s.lines, NewLine(scanner.Text()))
	}
}

// NewLine creates a new line after line
func (s *Store) NewLine(st string, line int) (int, error) {
	if line < 0 || line >= len(s.lines) {
		return 0, fmt.Errorf("NewLine: Invalid line %v", line)
	}
	s.lines = append(s.lines[:line], append([]*Line{NewLine(st)}, s.lines[line:]...)...)
	cs := s.undoFac()
	cs.AddLine(line)
	cs.ChangeLine(line+1, "", st)
	s.AddUndoSet(cs)
	return line, nil
}

// DeleteLine will remove a line
func (s *Store) DeleteLine(line int) (string, error) {
	if line < 0 || line >= len(s.lines) {
		return "", fmt.Errorf("NewLine: Invalid line %v", line)
	}
	original := s.lines[line].text
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
func (s *Store) ResetLine(st string, line int) (string, error) {
	if line < 0 || line >= len(s.lines) {
		return "", fmt.Errorf("ResetLine: Invalid offset %v", line)
	}
	original := s.lines[line].text
	s.lines[line].Reset(st)
	cs := s.undoFac()
	cs.ChangeLine(line, original, st)
	s.AddUndoSet(cs)
	return original, nil
}

// LineString will return the line as a string
func (s *Store) LineString(line int) (string, error) {
	if line < 0 || line >= len(s.lines) {
		return "", fmt.Errorf("LineString: Invalid offset %v", line)
	}
	return s.lines[line].text, nil
}

// LineRangeString will return an array of line content
func (s *Store) LineRangeString(line, cnt int) ([]string, error) {
	nl := len(s.lines)
	if line < 0 || line >= nl {
		return nil, fmt.Errorf("LineRangeString: Invalid offset %v", line)
	}
	list := []string{}
	for i := line; i < line+cnt; i++ {
		if i >= nl {
			break
		}
		list = append(list, s.lines[i].text)
	}
	return list, nil
}

// NumLines returns the number of lines
func (s *Store) NumLines() int {
	return len(s.lines)
}

// LineLen returns the length of a line without delimeters
func (s *Store) LineLen(line int) int {
	if line < 0 || line >= len(s.lines) {
		return -1
	}
	return s.lines[line].Len()
}

// String will return all lines with delimeters
func (s *Store) String() string {
	builder := strings.Builder{}
	for i := range s.lines {
		builder.WriteString(s.lines[i].String())
		builder.WriteString(s.delim)
	}
	return builder.String()
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

// ReadRuneAt will return the run within a line
func (s *Store) ReadRuneAt(line, col int) (rune, int, error) {
	if line < 0 || line >= len(s.lines) {
		return 0, 0, fmt.Errorf("ReadRuneAt: Invalid line %v", line)
	}
	if col >= s.lines[line].Len() {
		return 0, 0, fmt.Errorf("ReadRuneAt: Invalid col %v", col)
	}
	reader, err := s.LineAt(line)
	if err != nil {
		return 0, 0, err
	}
	return reader.ReadRuneAt(int64(col))
}

// LineAt returns the line at line
func (s *Store) LineAt(line int) (LineReader, error) {
	if line < 0 || line >= len(s.lines) {
		return nil, fmt.Errorf("LineAt: Invalid line %v", line)
	}
	return NewReader(s.lines[line]), nil
}

// LineWriterAt returns a writer for line at line
func (s *Store) LineWriterAt(line int) (LineWriter, error) {
	if line < 0 || line >= len(s.lines) {
		return nil, fmt.Errorf("LineWriterAt: Invalid line %v", line)
	}
	return NewWriter(s.lines[line], line, s), nil
}

// SetLineDelim will set the line delimeter
func (s *Store) SetLineDelim(str string) {
	s.delim = str
}

// LineDelim will return the current linedelimeter
func (s *Store) LineDelim() string {
	return s.delim
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
			s.NewLine(change.Undo(""), change.Line())

		case undo.AddLine:
			s.DeleteLine(change.Line())

		case undo.ChangeLine:
			txt, _ := s.LineString(change.Line())
			s.ResetLine(change.Undo(txt), change.Line())
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
			s.NewLine(change.Undo(""), change.Line())

		case undo.AddLine:
			s.DeleteLine(change.Line())

		case undo.ChangeLine:
			txt, _ := s.LineString(change.Line())
			s.ResetLine(change.Undo(txt), change.Line())
		}
	}
	return nil
}

// StartGroupUndo will defer undo save until stopped
// grouping all undos together
func (s *Store) StartGroupUndo() {
	s.hold = []undo.ChangeSetter{}
	s.grpUndo = true
}

// StopGroupUndo will save undos as a simgle undo
func (s *Store) StopGroupUndo() {
	s.grpUndo = false
	cs := undo.New()
	for i := range s.hold {
		cs.AddChanges(s.hold[i].Changes()...)
	}
	s.hold = []undo.ChangeSetter{}
	s.AddUndoSet(cs)
}

// AddUndoSet will add an undo change set to the store
func (s *Store) AddUndoSet(cs undo.ChangeSetter) {
	switch {
	case s.grpUndo:
		s.hold = append(s.hold, cs)
	case s.isUndoing:
		s.redo = append(s.redo, cs)
	default:
		s.undo = append(s.undo, cs)
	}
}

// New returns a new Store
func New(fac undo.Factory) TextStorer {
	return &Store{undoFac: fac, delim: "\n"}
}
