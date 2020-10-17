# buffer
--
    import "."


## Usage

#### type Buffer

```go
type Buffer struct {
}
```

Buffer is a text buffer

#### func (*Buffer) BeginSelect

```go
func (b *Buffer) BeginSelect()
```
BeginSelect will save the current cursor position

#### func (*Buffer) Content

```go
func (b *Buffer) Content(start, cnt int) ([]string, error)
```
Content will return the content within a range

#### func (*Buffer) Cursor

```go
func (b *Buffer) Cursor() cursor.Cursor
```
Cursor will return the buffer's cursor

#### func (*Buffer) DeleteChar

```go
func (b *Buffer) DeleteChar(line, col, cnt int) error
```
DeleteChar will delete the next char

#### func (*Buffer) DeleteCharBack

```go
func (b *Buffer) DeleteCharBack(line, col, cnt int) error
```
DeleteCharBack will delete the prev char

#### func (*Buffer) DeleteLine

```go
func (b *Buffer) DeleteLine(line, cnt int) error
```
DeleteLine will remove a line

#### func (*Buffer) DeleteObject

```go
func (b *Buffer) DeleteObject(line, col int, obj textobject.TextObjecter, cnt int) error
```
DeleteObject will remove the next occurence of an object

#### func (*Buffer) Dirty

```go
func (b *Buffer) Dirty() bool
```
Dirty will return true if the buffer is unsaved

#### func (*Buffer) Down

```go
func (b *Buffer) Down(cnt int)
```
Down will move the cursor down cnt

#### func (*Buffer) EndSelect

```go
func (b *Buffer) EndSelect()
```
EndSelect will save the current position

#### func (*Buffer) Filename

```go
func (b *Buffer) Filename() string
```
Filename will return the buffer filename

#### func (*Buffer) Filetype

```go
func (b *Buffer) Filetype() string
```
Filetype returns the buffer's file type

#### func (*Buffer) ID

```go
func (b *Buffer) ID() string
```
ID will return the identifier for the buffer

#### func (*Buffer) Indent

```go
func (b *Buffer) Indent(line, cnt int) error
```
Indent will indent the current line

#### func (*Buffer) InsertString

```go
func (b *Buffer) InsertString(line, col int, st string) error
```
InsertString will insert a string at line, col

#### func (*Buffer) Move

```go
func (b *Buffer) Move(cnt int, obj textobject.TextObjecter) error
```
Move will move to the next cnt objects

#### func (*Buffer) MoveEnd

```go
func (b *Buffer) MoveEnd(cnt int, obj textobject.TextObjecter) error
```
MoveEnd will move cnt object end

#### func (*Buffer) MovePrev

```go
func (b *Buffer) MovePrev(cnt int, obj textobject.TextObjecter) error
```
MovePrev will move cnt prev objects

#### func (*Buffer) MovePrevEnd

```go
func (b *Buffer) MovePrevEnd(cnt int, obj textobject.TextObjecter) error
```
MovePrevEnd will move cnt previous objects at end

#### func (*Buffer) MoveTo

```go
func (b *Buffer) MoveTo(line, col int) error
```
MoveTo will move to a line and column

#### func (*Buffer) NewLineAbove

```go
func (b *Buffer) NewLineAbove(line int, st string, cnt int) error
```
NewLineAbove will add a line above line with string st

#### func (*Buffer) NewLineBelow

```go
func (b *Buffer) NewLineBelow(line int, st string, cnt int) error
```
NewLineBelow will add a line below line with string st

#### func (*Buffer) Next

```go
func (b *Buffer) Next(cnt int)
```
Next will move the cursor forward by cnt

#### func (*Buffer) OpenFile

```go
func (b *Buffer) OpenFile(path string) error
```
OpenFile will open

#### func (*Buffer) Outdent

```go
func (b *Buffer) Outdent(line, cnt int) error
```
Outdent will decrease the indent level

#### func (*Buffer) Position

```go
func (b *Buffer) Position() []int
```
Position will return the current cursor position

#### func (*Buffer) Prev

```go
func (b *Buffer) Prev(cnt int)
```
Prev will move the curosr back by cnt

#### func (*Buffer) Redo

```go
func (b *Buffer) Redo() error
```
Redo will redo the last edit

#### func (*Buffer) RenameFile

```go
func (b *Buffer) RenameFile(path string) error
```
RenameFile will rename the file

#### func (*Buffer) ReplaceObject

```go
func (b *Buffer) ReplaceObject(line, col int, obj textobject.TextObjecter, s string, cnt int) error
```
ReplaceObject will replace an object with s

#### func (*Buffer) Reset

```go
func (b *Buffer) Reset(st string)
```
Reset will reset the buffer content

#### func (*Buffer) SaveBuffer

```go
func (b *Buffer) SaveBuffer(path string) error
```
SaveBuffer will save the buffer to disk if path is specified it is used
otherwise it uses the current name

#### func (*Buffer) ScrollDown

```go
func (b *Buffer) ScrollDown()
```
ScrollDown will scroll the cursor down

#### func (*Buffer) ScrollUp

```go
func (b *Buffer) ScrollUp()
```
ScrollUp will scroll the cursor up

#### func (*Buffer) Search

```go
func (b *Buffer) Search(s string) ([]SearchResult, error)
```
Search will search the current textstore

#### func (*Buffer) SearchResults

```go
func (b *Buffer) SearchResults() []SearchResult
```
SearchResults will return the current search results

#### func (*Buffer) Selection

```go
func (b *Buffer) Selection() [][]int
```
Selection will return the cursor's selection

#### func (*Buffer) SetFilename

```go
func (b *Buffer) SetFilename(n string)
```
SetFilename will set the buffers file name

#### func (*Buffer) SetFiletype

```go
func (b *Buffer) SetFiletype(ft string)
```
SetFiletype will set the buffer's file type

#### func (*Buffer) StartGroupUndo

```go
func (b *Buffer) StartGroupUndo()
```
StartGroupUndo will group edits into a single undo

#### func (*Buffer) StopGroupUndo

```go
func (b *Buffer) StopGroupUndo()
```
StopGroupUndo will stop grouping undos

#### func (*Buffer) SyntaxResults

```go
func (b *Buffer) SyntaxResults(update bool, fgrps ...string) []syntax.Resulter
```
SyntaxResults returns the syntax scanning results

#### func (*Buffer) SyntaxResultsRange

```go
func (b *Buffer) SyntaxResultsRange(ln, cnt int, update bool, fgrps ...string) []syntax.Resulter
```
SyntaxResultsRange returns the syntax scanning results

#### func (*Buffer) TextStore

```go
func (b *Buffer) TextStore() textstore.TextStorer
```
TextStore will return the buffer's text store

#### func (*Buffer) Undo

```go
func (b *Buffer) Undo() error
```
Undo will undo the last edit

#### func (*Buffer) Up

```go
func (b *Buffer) Up(cnt int)
```
Up will move the cursor up cnt

#### type Bufferer

```go
type Bufferer interface {
	ID() string
	Content(start, count int) ([]string, error)
	TextStore() textstore.TextStorer
	Cursor() cursor.Cursor
	Filer
	Mover
	TextEditor
	Selector
	SyntaxResults(update bool, filterGroups ...string) []syntax.Resulter
	SyntaxResultsRange(ln, cnt int, update bool, filterGroups ...string) []syntax.Resulter
	SearchResults() []SearchResult
	Search(string) ([]SearchResult, error)
}
```

Bufferer is a text buffer

#### func  New

```go
func New(txt textstore.TextStorer, cur cursor.Cursor, m syntax.Manager, ftd filetype.Manager, reg register.Registerer) Bufferer
```
New will return a new Buffer

#### type Factory

```go
type Factory func(txt textstore.TextStorer, cur cursor.Cursor, m syntax.Manager, ftd filetype.Manager, reg register.Registerer) Bufferer
```

Factory is a function that returns new bufferers

#### type Filer

```go
type Filer interface {
	Filename() string
	SetFilename(n string)
	Filetype() string
	SetFiletype(ft string)
	SaveBuffer(path string) error
	OpenFile(path string) error
	RenameFile(path string) error
	Dirty() bool
}
```

Filer is file functions

#### type Mover

```go
type Mover interface {
	Move(cnt int, obj textobject.TextObjecter) error
	MoveEnd(cnt int, obj textobject.TextObjecter) error
	MovePrev(cnt int, obj textobject.TextObjecter) error
	MovePrevEnd(cnt int, obj textobject.TextObjecter) error
	MoveTo(line, col int) error
	Up(int)
	Down(int)
	Prev(int)
	Next(int)
	ScrollDown()
	ScrollUp()
	Position() []int
}
```

Mover is cursor movement functions

#### type SearchResult

```go
type SearchResult struct {
	Line    int
	Matches [][]int
}
```

SearchResult is the results for a line

#### type Selector

```go
type Selector interface {
	BeginSelect()
	EndSelect()
	Selection() [][]int
}
```

Selector is selection functions

#### type TextEditor

```go
type TextEditor interface {
	Reset(string)
	ReplaceObject(line, col int, obj textobject.TextObjecter, s string, cnt int) error
	DeleteChar(line, col, cnt int) error
	DeleteCharBack(line, col, cnt int) error
	DeleteLine(line, cnt int) error
	DeleteObject(line, col int, obj textobject.TextObjecter, cnt int) error
	InsertString(line, col int, st string) error
	NewLineAbove(line int, st string, cnt int) error
	NewLineBelow(line int, st string, cnt int) error
	Indent(ln, cnt int) error
	Outdent(ln, cnt int) error
	Undo() error
	Redo() error
	StartGroupUndo()
	StopGroupUndo()
}
```

TextEditor is text editing functions
