package textstore

/*
	textstore manages text storage for an editor
	textstore is a zero based interface meaning that lines and columns start at 0.
	Generaly an editor will use lines and columns starting with 1. This can be a bit confusing
*/

import (
	"fmt"
	"io"

	"github.com/dshills/layered/undo"
)

// Factory is a function the returns a new LineBuf
type Factory func(undo.Factory) TextStorer

// TextStorer is a generalized text store
type TextStorer interface {
	Undoer
	StoreWriter
	LineWriter
	StoreReader
	LineReader
	fmt.Stringer
	StoreSubscriber
}

// Undoer is generalized undo functionality
type Undoer interface {
	BeginEdit()
	EndEdit()
	AddUndoSet(undo.ChangeSetter)
	Undo() error
	Redo() error
}

// StoreWriter are store write functionality
type StoreWriter interface {
	Reset(s string) uint64
	NewLine(line int, s string)
	DeleteLine(line int) (string, error)
	ResetLine(line int, s string) (string, error)
	SetLineDelim(str string)
	ReadFrom(r io.Reader) (int64, error)
}

// LineWriter are line write functions
type LineWriter interface {
	Replace(line, fromcol, tocol int, s string) error
	Insert(line, col int, s string) error
	Delete(line, col, cnt int) error
}

// StoreReader store reading functions
type StoreReader interface {
	Hash64() uint64
	Len() int
	LineDelim() string
	WriteTo(w io.Writer) (n int64, err error)
	LineRange(line, cnt int) ([]string, error)
	LineString(line int) (string, error)
	NumLines() int
}

// LineReader is line reading functionality
type LineReader interface {
	LineLen(col int) int
	RuneAt(line, col int) (rune, error)
}

// StoreSubscriber is subscriber functionality
type StoreSubscriber interface {
	Subscribe(id string, up chan uint64)
	Unsubscribe(id string)
}

// Reader implements the io.Reader, io.ReaderAt, io.Seeker, io.WriterTo, io.ByteScanner,
// and io.RuneScanner interfaces by reading from a textstore.
type Reader interface {
	Len() int
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	ReadByte() (byte, error)
	ReadRune() (ch rune, size int, err error)
	Reset(s string)
	Seek(offset int64, whence int) (int64, error)
	Size() int64
	UnreadByte() error
	UnreadRune() error
	WriteTo(w io.Writer) (n int64, err error)
}
