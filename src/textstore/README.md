# textstore
--
    import "."


## Usage

#### type Factory

```go
type Factory func(undo.Factory) TextStorer
```

Factory is a function the returns a new LineBuf

#### type Line

```go
type Line struct {
}
```

Line is a linebuf line of text

#### func  NewLine

```go
func NewLine(s string) *Line
```
NewLine returns a new line set to s

#### func (*Line) InsertAt

```go
func (l *Line) InsertAt(p []byte, offset int64) (int, error)
```
InsertAt will insert buffer p at offset offset greater then text length will be
appended offset <= 0 will write to the beginning of the line

#### func (*Line) InsertRuneAt

```go
func (l *Line) InsertRuneAt(r rune, offset int) error
```
InsertRuneAt will insert rune r at offset offset <= 0 r will be written to the
beginning of the line offset >= len line r will be appended to the eol

#### func (*Line) Len

```go
func (l *Line) Len() int
```
Len returns the length of the line text not including delimeters in bytes

#### func (*Line) ReplaceAt

```go
func (l *Line) ReplaceAt(p []byte, offset, cnt int64) (int, error)
```
ReplaceAt will replace offset - offset + cnt with buffer p if offset <= 0 will
text will be replaced starting at the beginning of the line

#### func (*Line) Reset

```go
func (l *Line) Reset(s string)
```
Reset will reset the line to string s

#### func (*Line) String

```go
func (l *Line) String() string
```

#### func (*Line) WriteAt

```go
func (l *Line) WriteAt(p []byte, offset int64) (int, error)
```
WriteAt will insert buffer p at offset existing content will be replaced and
remaining will be appended WriteAt implements the io.WriterAt interface

#### func (*Line) WriteRuneAt

```go
func (l *Line) WriteRuneAt(r rune, offset int) error
```
WriteRuneAt will insert rune r at rune offset (not byte ofset) offset >=
len([]rune) r will be appended to the eol

#### type LineReader

```go
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
```

LineReader provides advanced reading for lines

#### type LineWriter

```go
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
```

LineWriter provides advanced editing for lines

#### type Liner

```go
type Liner interface {
	LineReader
	LineWriter
}
```

Liner is a generalized text storage interface for a single line of text

#### type Reader

```go
type Reader struct {
}
```

A Reader implements the io.Reader, io.ReaderAt, io.Seeker, io.WriterTo,
io.ByteScanner, and io.RuneScanner interfaces by reading from a string. The zero
value for Reader operates like a Reader of an empty string.

#### func  NewReader

```go
func NewReader(l *Line) *Reader
```
NewReader returns a new Reader reading from s. It is similar to
bytes.NewBufferString but more efficient and read-only.

#### func (*Reader) Len

```go
func (r *Reader) Len() int
```
Len returns the number of bytes of the unread portion of the string.

#### func (*Reader) Read

```go
func (r *Reader) Read(b []byte) (n int, err error)
```

#### func (*Reader) ReadAt

```go
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
```
ReadAt will into p at position off

#### func (*Reader) ReadByte

```go
func (r *Reader) ReadByte() (byte, error)
```
ReadByte will read the next byte from the seek position

#### func (*Reader) ReadRune

```go
func (r *Reader) ReadRune() (ch rune, size int, err error)
```
ReadRune will read the next rune from the seek position

#### func (*Reader) ReadRuneAt

```go
func (r *Reader) ReadRuneAt(offset int64) (rune, int, error)
```
ReadRuneAt returns the rune at rune offset (not byte offset)

#### func (*Reader) Seek

```go
func (r *Reader) Seek(offset int64, whence int) (int64, error)
```
Seek implements the io.Seeker interface.

#### func (*Reader) Size

```go
func (r *Reader) Size() int64
```
Size returns the original length of the underlying string. Size is the number of
bytes available for reading via ReadAt. The returned value is always the same
and is not affected by calls to any other method.

#### func (*Reader) String

```go
func (r *Reader) String() string
```

#### func (*Reader) UnreadByte

```go
func (r *Reader) UnreadByte() error
```
UnreadByte will decrement the seek position by 1 byte

#### func (*Reader) UnreadRune

```go
func (r *Reader) UnreadRune() error
```
UnreadRune will decrement the seek position by size of last rune

#### func (*Reader) WriteTo

```go
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```
WriteTo implements the io.WriterTo interface.

#### type StReader

```go
type StReader interface {
	LineString(pos int) (string, error)
	LineRangeString(line, cnt int) ([]string, error)
	ReadRuneAt(line, col int) (rune, int, error)
	LineDelim() string
	String() string // adds delimeters
}
```

StReader is reader functionality

#### type StWriter

```go
type StWriter interface {
	Reset(s string) uint64
	ReadFrom(r io.Reader) (int64, error)
	NewLine(s string, line int) (int, error)
	DeleteLine(line int) (string, error)
	ResetLine(s string, line int) (string, error)
	SetLineDelim(str string)
}
```

StWriter is writer functionality

#### type Store

```go
type Store struct {
}
```

Store is an implementation of LineBuffer

#### func (*Store) AddUndoSet

```go
func (s *Store) AddUndoSet(cs undo.ChangeSetter)
```
AddUndoSet will add an undo change set to the store

#### func (*Store) DeleteLine

```go
func (s *Store) DeleteLine(line int) (string, error)
```
DeleteLine will remove a line

#### func (*Store) Hash64

```go
func (s *Store) Hash64() uint64
```
Hash64 will return a 64bit hash

#### func (*Store) Len

```go
func (s *Store) Len() int
```
Len will return the total length with delimeters

#### func (*Store) LineAt

```go
func (s *Store) LineAt(line int) (LineReader, error)
```
LineAt returns the line at line

#### func (*Store) LineDelim

```go
func (s *Store) LineDelim() string
```
LineDelim will return the current linedelimeter

#### func (*Store) LineLen

```go
func (s *Store) LineLen(line int) int
```
LineLen returns the length of a line without delimeters

#### func (*Store) LineRangeString

```go
func (s *Store) LineRangeString(line, cnt int) ([]string, error)
```
LineRangeString will return an array of line content

#### func (*Store) LineString

```go
func (s *Store) LineString(line int) (string, error)
```
LineString will return the line as a string

#### func (*Store) LineWriterAt

```go
func (s *Store) LineWriterAt(line int) (LineWriter, error)
```
LineWriterAt returns a writer for line at line

#### func (*Store) NewLine

```go
func (s *Store) NewLine(st string, line int) (int, error)
```
NewLine creates a new line after line

#### func (*Store) NumLines

```go
func (s *Store) NumLines() int
```
NumLines returns the number of lines

#### func (*Store) ReadFrom

```go
func (s *Store) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom will read from an io.Reader

#### func (*Store) ReadRuneAt

```go
func (s *Store) ReadRuneAt(line, col int) (rune, int, error)
```
ReadRuneAt will return the run within a line

#### func (*Store) Redo

```go
func (s *Store) Redo() error
```
Redo will undo an undo

#### func (*Store) Reset

```go
func (s *Store) Reset(st string) uint64
```
Reset will set the Store to s

#### func (*Store) ResetLine

```go
func (s *Store) ResetLine(st string, line int) (string, error)
```
ResetLine will set the contents of a line

#### func (*Store) SetLineDelim

```go
func (s *Store) SetLineDelim(str string)
```
SetLineDelim will set the line delimeter

#### func (*Store) StartGroupUndo

```go
func (s *Store) StartGroupUndo()
```
StartGroupUndo will defer undo save until stopped grouping all undos together

#### func (*Store) StopGroupUndo

```go
func (s *Store) StopGroupUndo()
```
StopGroupUndo will save undos as a simgle undo

#### func (*Store) String

```go
func (s *Store) String() string
```
String will return all lines with delimeters

#### func (*Store) Subscribe

```go
func (s *Store) Subscribe(id string, up chan uint64)
```
Subscribe will subscribe to updates

#### func (*Store) Undo

```go
func (s *Store) Undo() error
```
Undo will undo the last set of edits

#### func (*Store) Unsubscribe

```go
func (s *Store) Unsubscribe(id string)
```
Unsubscribe will remove a subscription

#### type StoreLengther

```go
type StoreLengther interface {
	NumLines() int
	LineLen(line int) int
	Len() int
}
```

StoreLengther is length functionality

#### type StoreSubscriber

```go
type StoreSubscriber interface {
	Subscribe(id string, up chan uint64)
	Unsubscribe(id string)
}
```

StoreSubscriber is subscriber functionality

#### type TextStorer

```go
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
```

TextStorer is a generalized text storage interface

#### func  New

```go
func New(fac undo.Factory) TextStorer
```
New returns a new Store

#### type Undoer

```go
type Undoer interface {
	StartGroupUndo()
	StopGroupUndo()
	AddUndoSet(undo.ChangeSetter)
	Undo() error
	Redo() error
}
```

Undoer is generalized undo functionality

#### type Writer

```go
type Writer struct {
}
```

Writer is a line writer

#### func  NewWriter

```go
func NewWriter(l *Line, line int, store *Store) *Writer
```
NewWriter returns a new line writer

#### func (*Writer) Flush

```go
func (w *Writer) Flush()
```
Flush will set the line to the buffered text

#### func (*Writer) Insert

```go
func (w *Writer) Insert(p []byte) (int, error)
```
Insert will insert buffer p at offset

#### func (*Writer) InsertAt

```go
func (w *Writer) InsertAt(p []byte, offset int64) (int, error)
```
InsertAt will insert buffer p at offset offset greater then text length will be
appended offset <= 0 will write to the beginning of the line

#### func (*Writer) InsertRuneAt

```go
func (w *Writer) InsertRuneAt(r rune, offset int64) (int, error)
```
InsertRuneAt will insert rune r at offset offset <= 0 r will be written to the
beginning of the line offset >= len line r will be appended to the eol

#### func (*Writer) Len

```go
func (w *Writer) Len() int
```
Len returns the number of bytes from the current seek position

#### func (*Writer) Replace

```go
func (w *Writer) Replace(p []byte, cnt int64) (int, error)
```
Replace will replace from seek offset to offset + cnt with buffer p

#### func (*Writer) ReplaceAt

```go
func (w *Writer) ReplaceAt(p []byte, offset, cnt int64) (int, error)
```
ReplaceAt will replace offset - offset + cnt with buffer p if offset <= 0
wilw.text will be replaced starting at the beginning of the line

#### func (*Writer) Reset

```go
func (w *Writer) Reset(s string)
```
Reset will replace the line content

#### func (*Writer) Seek

```go
func (w *Writer) Seek(offset int64, whence int) (int64, error)
```
Seek implements the io.Seeker interface.

#### func (*Writer) String

```go
func (w *Writer) String() string
```

#### func (*Writer) Write

```go
func (w *Writer) Write(p []byte) (int, error)
```
Write will write buffer p to the line at the seek position depending on current
mode it will either insert inside the current text or replace the text at the
current seek position

#### func (*Writer) WriteAt

```go
func (w *Writer) WriteAt(p []byte, offset int64) (int, error)
```
WriteAt will insert buffer p at offset existing content will be replaced and
remaining will be appended WriteAt implements the io.WriterAt interface

#### func (*Writer) WriteByte

```go
func (w *Writer) WriteByte(b byte) error
```
WriteByte will write a byte at the seek position

#### func (*Writer) WriteRune

```go
func (w *Writer) WriteRune(r rune) (int, error)
```
WriteRune will write a rune at the seek position

#### func (*Writer) WriteRuneAt

```go
func (w *Writer) WriteRuneAt(r rune, offset int64) (int, error)
```
WriteRuneAt will insert rune r at rune offset (not byte ofset)

#### func (*Writer) WriteString

```go
func (w *Writer) WriteString(s string) (int, error)
```
WriteString will write a string at the seek position
