# textstore
--
    import "."


## Usage

#### type Factory

```go
type Factory func(undo.Factory) TextStorer
```

Factory is a function the returns a new LineBuf

#### type LineReader

```go
type LineReader interface {
	LineLen(col int) int
	RuneAt(line, col int) (rune, error)
}
```

LineReader is line reading functionality

#### type LineWriter

```go
type LineWriter interface {
	Replace(line, fromcol, tocol int, s string) error
	Insert(line, col int, s string) error
	Delete(line, col, cnt int) error
}
```

LineWriter are line write functions

#### type Store

```go
type Store struct {
}
```

Store is a TextStore

#### func (*Store) AddUndoSet

```go
func (s *Store) AddUndoSet(cs undo.ChangeSetter)
```
AddUndoSet will add an undo change set to the store

#### func (*Store) BeginEdit

```go
func (s *Store) BeginEdit()
```
BeginEdit will start edit tracking

#### func (*Store) Delete

```go
func (s *Store) Delete(ln, col, cnt int) error
```
Delete will delete text at line, col

#### func (*Store) DeleteLine

```go
func (s *Store) DeleteLine(line int) (string, error)
```
DeleteLine will remove a line

#### func (*Store) EndEdit

```go
func (s *Store) EndEdit()
```
EndEdit will save undos since begin

#### func (*Store) Hash64

```go
func (s *Store) Hash64() uint64
```
Hash64 will return a 64bit hash

#### func (*Store) Insert

```go
func (s *Store) Insert(ln, col int, st string) error
```
Insert inserts text to line

#### func (*Store) Len

```go
func (s *Store) Len() int
```
Len will return the total length with delimeters

#### func (*Store) LineDelim

```go
func (s *Store) LineDelim() string
```
LineDelim will return the current linedelimeter

#### func (*Store) LineLen

```go
func (s *Store) LineLen(line int) int
```
LineLen returns the lkength of a line

#### func (*Store) LineRange

```go
func (s *Store) LineRange(ln, cnt int) ([]string, error)
```
LineRange returns a range of lines

#### func (*Store) LineString

```go
func (s *Store) LineString(line int) (string, error)
```
LineString will return the line as a string

#### func (*Store) NewLine

```go
func (s *Store) NewLine(ln int, st string)
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

#### func (*Store) Redo

```go
func (s *Store) Redo() error
```
Redo will undo an undo

#### func (*Store) Replace

```go
func (s *Store) Replace(ln, from, to int, st string) error
```
Replace replaces text at ln, col

#### func (*Store) Reset

```go
func (s *Store) Reset(st string) uint64
```
Reset will set the Store to s

#### func (*Store) ResetLine

```go
func (s *Store) ResetLine(line int, st string) (string, error)
```
ResetLine will set the contents of a line

#### func (*Store) RuneAt

```go
func (s *Store) RuneAt(ln, col int) (rune, error)
```
RuneAt returns the rune at ln, col

#### func (*Store) SetLineDelim

```go
func (s *Store) SetLineDelim(str string)
```
SetLineDelim will set the line delimeter

#### func (*Store) String

```go
func (s *Store) String() string
```

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

#### func (*Store) WriteTo

```go
func (s *Store) WriteTo(w io.Writer) (int64, error)
```
WriteTo will write the store to w

#### type StoreReader

```go
type StoreReader interface {
	Hash64() uint64
	Len() int
	LineDelim() string
	WriteTo(w io.Writer) (n int64, err error)
	LineRange(line, cnt int) ([]string, error)
	LineString(line int) (string, error)
	NumLines() int
}
```

StoreReader store reading functions

#### type StoreSubscriber

```go
type StoreSubscriber interface {
	Subscribe(id string, up chan uint64)
	Unsubscribe(id string)
}
```

StoreSubscriber is subscriber functionality

#### type StoreWriter

```go
type StoreWriter interface {
	Reset(s string) uint64
	NewLine(line int, s string)
	DeleteLine(line int) (string, error)
	ResetLine(line int, s string) (string, error)
	SetLineDelim(str string)
	ReadFrom(r io.Reader) (int64, error)
}
```

StoreWriter are store write functionality

#### type TextStorer

```go
type TextStorer interface {
	Undoer
	StoreWriter
	LineWriter
	StoreReader
	LineReader
	fmt.Stringer
	StoreSubscriber
}
```

TextStorer is a generalized text store

#### func  New

```go
func New(uf undo.Factory) TextStorer
```
New returns a TextStorer

#### type Undoer

```go
type Undoer interface {
	BeginEdit()
	EndEdit()
	AddUndoSet(undo.ChangeSetter)
	Undo() error
	Redo() error
}
```

Undoer is generalized undo functionality
