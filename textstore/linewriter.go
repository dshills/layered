package textstore

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/dshills/layered/logger"
	"github.com/dshills/layered/undo"
)

// Writer is a line writer
type Writer struct {
	line      int
	ln        *Line
	text      []rune
	i         int64 // <== rune index not byte
	store     *Store
	changeSet undo.ChangeSetter
}

// Flush will set the line to the buffered text
func (w *Writer) Flush() {
	w.changeSet.ChangeLine(w.line, w.ln.String(), string(w.text))
	w.ln.Reset(string(w.text))
	w.store.AddUndoSet(w.changeSet)
}

// Reset will replace the line content
func (w *Writer) Reset(s string) {
	w.text = []rune(s)
}

// Seek implements the io.Seeker interface.
func (w *Writer) Seek(offset int64, whence int) (int64, error) {
	var abs int64
	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = w.i + offset
	case io.SeekEnd:
		abs = int64(len(w.text)) + offset
	default:
		return 0, errors.New("Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("Seek: negative position")
	}
	w.i = abs
	return abs, nil
}

// Write will write buffer p to the line at the seek position
// depending on current mode it will either insert inside the current text
// or replace the text at the current seek position
func (w *Writer) Write(p []byte) (int, error) {
	lnpr := int64(len([]rune(string(p))))
	builder := strings.Builder{}
	if w.i > 0 {
		builder.WriteString(string(w.text[:w.i]))
	}
	builder.WriteString(string(p))
	if w.i+lnpr < int64(len(w.text)) {
		builder.WriteString(string(w.text[w.i:]))
	}
	w.text = []rune(builder.String())
	w.i += lnpr
	return len(p), nil
}

// Replace will replace from seek offset to offset + cnt with buffer p
func (w *Writer) Replace(p []byte, cnt int64) (int, error) {
	if cnt <= 0 {
		return 0, fmt.Errorf("ReplaceAt: Count <= 0")
	}
	builder := strings.Builder{}
	if w.i > 0 && w.i < int64(len(w.text)) {
		builder.WriteString(string(w.text[:w.i]))
	}
	builder.WriteString(string(p))
	if w.i+cnt < int64(len(w.text)) {
		builder.WriteString(string(w.text[w.i+cnt:]))
	}
	w.text = []rune(builder.String())
	return len(p), nil
}

// Insert will insert buffer p at offset
func (w *Writer) Insert(p []byte) (int, error) {
	ln := int64(len(w.text))
	bld := strings.Builder{}
	if w.i > 0 && w.i < ln {
		bld.WriteString(string(w.text[:w.i]))
	}
	bld.WriteString(string(p))
	if w.i < ln {
		bld.WriteString(string(w.text[w.i:]))
	}
	w.text = []rune(bld.String())
	return len(p), nil
}

// Len returns the number of bytes from the current seek position
func (w *Writer) Len() int {
	return len(w.text) - int(w.i)
}

func (w *Writer) String() string {
	return string(w.text)
}

// WriteByte will write a byte at the seek position
func (w *Writer) WriteByte(b byte) error {
	_, err := w.Write([]byte{b})
	return err
}

// WriteRune will write a rune at the seek position
func (w *Writer) WriteRune(r rune) (int, error) {
	return w.Write([]byte(string(r)))
}

// WriteString will write a string at the seek position
func (w *Writer) WriteString(s string) (int, error) {
	return w.Write([]byte(s))
}

// WriteRuneAt will insert rune r at rune offset (not byte ofset)
func (w *Writer) WriteRuneAt(r rune, offset int64) (int, error) {
	return w.WriteAt([]byte(string(r)), offset)
}

// InsertRuneAt will insert rune r at offset
// offset <= 0 r will be written to the beginning of the line
// offset >= len line r will be appended to the eol
func (w *Writer) InsertRuneAt(r rune, offset int64) (int, error) {
	return w.InsertAt([]byte(string(r)), offset)
}

// WriteAt will insert buffer p at offset
// existing content will be replaced and remaining will be appended
// WriteAt implements the io.WriterAt interface
func (w *Writer) WriteAt(p []byte, offset int64) (int, error) {
	lnp := int64(len(p))
	ln := int64(len(w.text))
	if offset < 0 || offset >= ln {
		return 0, fmt.Errorf("WriteAt: offset out of range %v", offset)
	}

	bld := strings.Builder{}
	if offset > 0 {
		bld.WriteString(string(w.text[:offset]))
	}
	bld.WriteString(string(p))
	if offset+lnp < ln {
		bld.WriteString(string(w.text[offset+lnp:]))
	}
	w.text = []rune(bld.String())
	return len(p), nil
}

// InsertAt will insert buffer p at offset
// offset greater then text length will be appended
// offset <= 0 will write to the beginning of the line
func (w *Writer) InsertAt(p []byte, offset int64) (int, error) {
	ln := int64(len(w.text))
	if offset < 0 {
		return 0, fmt.Errorf("InsertAt: offset out of range %v", offset)
	}
	logger.Debugf("Writer.InsertAt: %v => %v (%v) at %v", string(p), string(w.text), ln, offset)
	bld := strings.Builder{}
	if offset > 0 && offset < ln {
		bld.WriteString(string(w.text[:offset]))
	}
	bld.WriteString(string(p))
	if offset < ln {
		bld.WriteString(string(w.text[offset:]))
	}
	w.text = []rune(bld.String())
	return len(p), nil
}

// ReplaceAt will replace offset - offset + cnt with buffer p
// if offset <= 0 wilw.text will be replaced starting at the beginning of the line
func (w *Writer) ReplaceAt(p []byte, offset, cnt int64) (int, error) {
	if cnt <= 0 {
		return 0, fmt.Errorf("ReplaceAt: Count <= 0")
	}
	builder := strings.Builder{}
	if offset > 0 {
		builder.WriteString(string(w.text[:offset]))
	}
	builder.WriteString(string(p))
	if offset+cnt < int64(len(w.text)) {
		builder.WriteString(string(w.text[offset+cnt:]))
	}
	w.text = []rune(builder.String())
	return len(p), nil
}

// NewWriter returns a new line writer
func NewWriter(l *Line, line int, store *Store) *Writer {
	return &Writer{ln: l, text: []rune(l.String()), line: line, store: store, changeSet: store.undoFac()}
}
