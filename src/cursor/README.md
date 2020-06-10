# cursor
--
    import "."


## Usage

#### type BufCursor

```go
type BufCursor struct {
}
```

BufCursor is a window cursor

#### func (*BufCursor) AsRange

```go
func (c *BufCursor) AsRange() []int
```
AsRange returns line, col as an int array

#### func (*BufCursor) Bottom

```go
func (c *BufCursor) Bottom() bool
```
Bottom will move to the last line in the buffer

#### func (BufCursor) Column

```go
func (c BufCursor) Column() int
```
Column will return the current column

#### func (*BufCursor) Down

```go
func (c *BufCursor) Down(cnt int) bool
```
Down moves the BufCursor down cnt lines

#### func (*BufCursor) EndTrack

```go
func (c *BufCursor) EndTrack()
```
EndTrack will save the ending position

#### func (BufCursor) Get

```go
func (c BufCursor) Get() (int, int)
```
Get will return the current line and column

#### func (*BufCursor) GotoLine

```go
func (c *BufCursor) GotoLine(ln int) bool
```
GotoLine will move the BufCursor to the specified line

#### func (BufCursor) Line

```go
func (c BufCursor) Line() int
```
Line will return the current line

#### func (*BufCursor) MoveValid

```go
func (c *BufCursor) MoveValid(line, col int) bool
```
MoveValid will move the BufCursor to line, col insuring it is a valid position

#### func (*BufCursor) Next

```go
func (c *BufCursor) Next(cnt int) bool
```
Next moves the BufCursor forward cnt chars

#### func (*BufCursor) Prev

```go
func (c *BufCursor) Prev(cnt int) bool
```
Prev moves the BufCursor back cnt chars

#### func (*BufCursor) StartTrack

```go
func (c *BufCursor) StartTrack()
```
StartTrack will save the current position

#### func (*BufCursor) Top

```go
func (c *BufCursor) Top() bool
```
Top will move the BufCursor to 0, 0

#### func (*BufCursor) Tracked

```go
func (c *BufCursor) Tracked() [][]int
```
Tracked will return the start and end position

#### func (*BufCursor) Up

```go
func (c *BufCursor) Up(cnt int) bool
```
Up moves the BufCursor up cnt lines

#### type Cursor

```go
type Cursor interface {
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

Cursor is an editor cursor

#### func  New

```go
func New(txt textstore.TextStorer) Cursor
```
New will return a new cursor

#### type Factory

```go
type Factory func(txt textstore.TextStorer) Cursor
```

Factory is a function that returns a new cursor
