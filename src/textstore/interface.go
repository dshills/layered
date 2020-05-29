package textstore

import (
	"fmt"
	"io"

	"github.com/dshills/layered/undo"
)

// Factory is a function the returns a new LineBuf
type Factory func(undo.Factory) TextStorer

// TextStorer is a generalized text storage interface
type TextStorer interface {
	Undoer
	StoreLengther
	StReader
	StWriter
	StoreSubscriber
	LineAt(line int) (LineReader, error)
	LineWriterAt(line int) (LineWriter, error)
	Hash64() uint64
}

// StoreSubscriber is subscriber functionality
type StoreSubscriber interface {
	Subscribe(id string, up chan uint64)
	Unsubscribe(id string)
}

// StoreLengther is length functionality
type StoreLengther interface {
	NumLines() int
	LineLen(line int) int
	Len() int
}

// StReader is reader functionality
type StReader interface {
	LineString(pos int) (string, error)
	LineRangeString(line, cnt int) ([]string, error)
	ReadRuneAt(line, col int) (rune, int, error)
	LineDelim() string
	String() string // adds delimeters
}

// StWriter is writer functionality
type StWriter interface {
	Reset(s string) uint64
	ReadFrom(r io.Reader) (int64, error)
	NewLine(s string, line int) (int, error)
	DeleteLine(line int) (string, error)
	ResetLine(s string, line int) (string, error)
	SetLineDelim(str string)
}

// Undoer is generalized undo functionality
type Undoer interface {
	StartGroupUndo()
	StopGroupUndo()
	AddUndoSet(undo.ChangeSetter)
	Undo() error
	Redo() error
}

// Liner is a generalized text storage interface
// for a single line of text
type Liner interface {
	LineReader
	LineWriter
}

// LineReader provides advanced reading for lines
type LineReader interface {
	Len() int
	Size() int64
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	ReadRuneAt(offset int64) (rune, int, error)
	ReadByte() (byte, error)
	UnreadByte() error
	ReadRune() (ch rune, size int, err error)
	UnreadRune() error
	Seek(offset int64, whence int) (int64, error)
	WriteTo(w io.Writer) (n int64, err error)
	fmt.Stringer
}

// LineWriter provides advanced editing for lines
type LineWriter interface {
	Flush()
	Len() int
	String() string
	Seek(offset int64, whence int) (int64, error)
	Reset(s string)
	Write(p []byte) (int, error)
	Replace(p []byte, cnt int64) (int, error)
	Insert(p []byte) (int, error)
	WriteByte(b byte) error
	WriteRune(r rune) (int, error)
	WriteString(s string) (int, error)
	WriteRuneAt(r rune, offset int64) (int, error)
	InsertRuneAt(r rune, offset int64) (int, error)
	WriteAt(p []byte, offset int64) (int, error)
	InsertAt(p []byte, offset int64) (int, error)
	ReplaceAt(p []byte, offset, cnt int64) (int, error)
}
