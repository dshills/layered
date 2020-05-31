# cursor
--
    import "."


## Usage

#### type Cursor

```go
type Cursor struct {
}
```

Cursor is a window cursor

#### func (*Cursor) AsRange

```go
func (c *Cursor) AsRange() []int
```
AsRange returns line, col as an int array

#### func (*Cursor) Bottom

```go
func (c *Cursor) Bottom() bool
```
Bottom will move to the last line in the buffer

#### func (Cursor) Column

```go
func (c Cursor) Column() int
```
Column will return the current column

#### func (*Cursor) Down

```go
func (c *Cursor) Down(cnt int) bool
```
Down moves the Cursor down cnt lines

#### func (*Cursor) EndTrack

```go
func (c *Cursor) EndTrack()
```
EndTrack will save the ending position

#### func (Cursor) Get

```go
func (c Cursor) Get() (int, int)
```
Get will return the current line and column

#### func (*Cursor) GotoLine

```go
func (c *Cursor) GotoLine(ln int) bool
```
GotoLine will move the Cursor to the specified line

#### func (Cursor) Line

```go
func (c Cursor) Line() int
```
Line will return the current line

#### func (*Cursor) MoveValid

```go
func (c *Cursor) MoveValid(line, col int) bool
```
MoveValid will move the Cursor to line, col insuring it is a valid position

#### func (*Cursor) Next

```go
func (c *Cursor) Next(cnt int) bool
```
Next moves the Cursor forward cnt chars

#### func (*Cursor) Prev

```go
func (c *Cursor) Prev(cnt int) bool
```
Prev moves the Cursor back cnt chars

#### func (*Cursor) StartTrack

```go
func (c *Cursor) StartTrack()
```
StartTrack will save the current position

#### func (*Cursor) Top

```go
func (c *Cursor) Top() bool
```
Top will move the Cursor to 0, 0

#### func (*Cursor) Tracked

```go
func (c *Cursor) Tracked() [][]int
```
Tracked will return the start and end position

#### func (*Cursor) Up

```go
func (c *Cursor) Up(cnt int) bool
```
Up moves the Cursor up cnt lines

#### type Cursorer

```go
type Cursorer interface {
	AsRange() []int
	Line() int
	Column() int
	GotoLine(ln int) bool
	Top() bool
	Bottom() bool
	Get() (int, int)
	Down(cnt int) bool
	Up(cnt int) bool
	Prev(cnt int) bool
	Next(cnt int) bool
	MoveValid(line, col int) bool
	StartTrack()
	EndTrack()
	Tracked() [][]int
}
```

Cursorer is an editor cursor

#### func  New

```go
func New(txt textstore.TextStorer) Cursorer
```
New will return a new cursor

#### type Factory

```go
type Factory func(txt textstore.TextStorer) Cursorer
```

Factory is a function that returns a new cursor
