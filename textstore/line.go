package textstore

import (
	"fmt"
	"strings"
)

// Line is a linebuf line of text
type Line struct {
	text string
}

// NewLine returns a new line set to s
func NewLine(s string) *Line {
	return &Line{text: s}
}

// Len returns the length of the line text not including delimeters in bytes
func (l *Line) Len() int {
	return len(l.text)
}

func (l *Line) String() string {
	return l.text
}

// Reset will reset the line to string s
func (l *Line) Reset(s string) {
	l.text = s
}

// WriteRuneAt will insert rune r at rune offset (not byte ofset)
// offset >= len([]rune) r will be appended to the eol
func (l *Line) WriteRuneAt(r rune, offset int) error {
	if offset < 0 {
		return fmt.Errorf("WriteRuneAt: offset out of range %v", offset)
	}
	builder := strings.Builder{}
	if offset > 0 {
		builder.WriteString(string([]rune(l.text)[:offset]))
	}
	builder.WriteRune(r)
	if offset < len([]rune(l.text)) {
		builder.WriteString(string([]rune(l.text)[offset+1:]))
	}
	l.text = builder.String()
	return nil
}

// InsertRuneAt will insert rune r at offset
// offset <= 0 r will be written to the beginning of the line
// offset >= len line r will be appended to the eol
func (l *Line) InsertRuneAt(r rune, offset int) error {
	builder := strings.Builder{}
	if offset <= 0 {
		offset = 0
		builder.WriteString(string([]rune(l.text)[:offset]))
	}
	builder.WriteRune(r)
	if offset < len([]rune(l.text)) {
		builder.WriteString(string([]rune(l.text)[offset:]))
	}
	l.text = builder.String()
	return nil
}

// WriteAt will insert buffer p at offset
// existing content will be replaced and remaining will be appended
// WriteAt implements the io.WriterAt interface
func (l *Line) WriteAt(p []byte, offset int64) (int, error) {
	lnp := int64(len(p))
	ln := int64(len(l.text))
	if offset < 0 || offset >= ln {
		return 0, fmt.Errorf("WriteAt: offset out of range %v", offset)
	}

	bld := strings.Builder{}
	txt := []rune(l.text)
	if offset > 0 {
		bld.WriteString(string(txt[:offset]))
	}
	bld.WriteString(string(p))
	if offset+lnp < ln {
		bld.WriteString(string(txt[offset+lnp:]))
	}
	l.text = bld.String()
	return len(p), nil
}

// InsertAt will insert buffer p at offset
// offset greater then text length will be appended
// offset <= 0 will write to the beginning of the line
func (l *Line) InsertAt(p []byte, offset int64) (int, error) {
	if offset < 0 {
		l.text = string(p) + l.text
		return len(p), nil
	}
	ln := int64(len(l.text))
	if offset >= ln || ln == 0 {
		l.text += string(p)
		return len(p), nil
	}
	bld := strings.Builder{}
	if offset > 0 {
		bld.WriteString(string([]rune(l.text)[:offset]))
	}
	bld.WriteString(string(p))
	bld.WriteString(string([]rune(l.text)[offset:]))
	l.text = bld.String()
	return len(p), nil
}

// ReplaceAt will replace offset - offset + cnt with buffer p
// if offset <= 0 will text will be replaced starting at the beginning of the line
func (l *Line) ReplaceAt(p []byte, offset, cnt int64) (int, error) {
	if cnt <= 0 {
		return 0, fmt.Errorf("ReplaceAt: Count <= 0")
	}
	builder := strings.Builder{}
	if offset > 0 {
		builder.WriteString(l.text[:offset])
	}
	builder.WriteString(string(p))
	if offset+cnt < int64(len(l.text)) {
		builder.WriteString(l.text[offset+cnt:])
	}
	l.text = builder.String()
	return len(p), nil
}
