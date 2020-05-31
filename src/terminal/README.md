# terminal
--
    import "."


## Usage

```go
const (
	ColorModeTrueColor = iota
	ColorMode256
	ColorMode16
	ColorMode8
)
```
Color Modes

#### type TermWriter

```go
type TermWriter struct {
}
```

TermWriter writes to a terminal

#### func  NewTermWriter

```go
func NewTermWriter(w io.Writer, clrmode int) *TermWriter
```
NewTermWriter will return a new terminal writer

#### func (TermWriter) Back

```go
func (tw TermWriter) Back(cnt int)
```
Back will move the cursor back cnt columns

#### func (TermWriter) BeginningLineDown

```go
func (tw TermWriter) BeginningLineDown(cnt int)
```
BeginningLineDown will move to the beginning of count lines below

#### func (TermWriter) BeginningLineUp

```go
func (tw TermWriter) BeginningLineUp(cnt int)
```
BeginningLineUp will move to the beginning of count lines above

#### func (TermWriter) Clear

```go
func (tw TermWriter) Clear()
```
Clear will clear the entire screen

#### func (TermWriter) ClearEntireLine

```go
func (tw TermWriter) ClearEntireLine()
```
ClearEntireLine will clear the line

#### func (TermWriter) ClearLine

```go
func (tw TermWriter) ClearLine()
```
ClearLine will clear the current line

#### func (TermWriter) ClearToEnd

```go
func (tw TermWriter) ClearToEnd()
```
ClearToEnd will clear from cursor to end of screen

#### func (TermWriter) ClearToEndOfLine

```go
func (tw TermWriter) ClearToEndOfLine()
```
ClearToEndOfLine will clear to EOL

#### func (TermWriter) ClearToStart

```go
func (tw TermWriter) ClearToStart()
```
ClearToStart will clear from cursor to start of screen

#### func (TermWriter) ClearToStartOfLine

```go
func (tw TermWriter) ClearToStartOfLine()
```
ClearToStartOfLine will clear to BOL

#### func (TermWriter) Down

```go
func (tw TermWriter) Down(cnt int)
```
Down will move the cursor down cnt lines

#### func (TermWriter) Forward

```go
func (tw TermWriter) Forward(cnt int)
```
Forward will advance the cursor cnt columns

#### func (TermWriter) Home

```go
func (tw TermWriter) Home()
```
Home moves the cursor to 0, 0

#### func (TermWriter) ResetStyle

```go
func (tw TermWriter) ResetStyle()
```
ResetStyle will clear any styles set

#### func (TermWriter) RestorePos

```go
func (tw TermWriter) RestorePos()
```
RestorePos will restore the cursor to the saved position

#### func (TermWriter) SavePos

```go
func (tw TermWriter) SavePos()
```
SavePos will save the current cursor position

#### func (TermWriter) SetBackground

```go
func (tw TermWriter) SetBackground(c palette.Color)
```
SetBackground will set the background color

#### func (TermWriter) SetBckCode

```go
func (tw TermWriter) SetBckCode(c int)
```
SetBckCode will set the background color based on the color mode 8, 16, 256,
truecolor

#### func (*TermWriter) SetColorMode

```go
func (tw *TermWriter) SetColorMode(clrmode int)
```
SetColorMode will set the writers color mode Used in SetFgdCode and SetBckCode
methods

#### func (TermWriter) SetFgdCode

```go
func (tw TermWriter) SetFgdCode(c int)
```
SetFgdCode will set the foreground color based on the color mode 8, 16, 256,
truecolor

#### func (TermWriter) SetForeground

```go
func (tw TermWriter) SetForeground(c palette.Color)
```
SetForeground will set the foreground color

#### func (TermWriter) SetStyle

```go
func (tw TermWriter) SetStyle(fgd, bck palette.Color)
```
SetStyle will set the foreground and backgroud color

#### func (TermWriter) To

```go
func (tw TermWriter) To(line, col int)
```
To will move to line, column

#### func (TermWriter) ToColumn

```go
func (tw TermWriter) ToColumn(c int)
```
ToColumn will move the cursor to column c of the current line

#### func (TermWriter) Up

```go
func (tw TermWriter) Up(cnt int)
```
Up will move the cursor up cnt lines

#### func (TermWriter) Write

```go
func (tw TermWriter) Write(b []byte) (int, error)
```
Write is a io.Writer

#### func (TermWriter) WriteEscString

```go
func (tw TermWriter) WriteEscString(s string) (int, error)
```
WriteEscString will write an escape followed by the string

#### func (TermWriter) WriteRune

```go
func (tw TermWriter) WriteRune(r rune) (int, error)
```
WriteRune will write a rune at the current cursor position

#### func (TermWriter) WriteString

```go
func (tw TermWriter) WriteString(s string) (int, error)
```
WriteString will write a string at the current cursor position

#### type WindowWriter

```go
type WindowWriter struct {
	BorderFgd  palette.Color
	BorderBck  palette.Color
	ContentFgd palette.Color
	ContentBck palette.Color
}
```

WindowWriter is a terminal writer supporting windowing

#### func  NewWindowWriter

```go
func NewWindowWriter(w *TermWriter, reg image.Rectangle) *WindowWriter
```
NewWindowWriter returns a writer

#### func (WindowWriter) Clear

```go
func (ww WindowWriter) Clear()
```
Clear will clear the content area

#### func (WindowWriter) Draw

```go
func (ww WindowWriter) Draw()
```
Draw will draw the window

#### func (WindowWriter) Fill

```go
func (ww WindowWriter) Fill(r rune)
```
Fill will fill the content area, including padding areas with character r

#### func (*WindowWriter) Resize

```go
func (ww *WindowWriter) Resize(reg image.Rectangle)
```
Resize will move the window to reg

#### func (*WindowWriter) SetBorderChars

```go
func (ww *WindowWriter) SetBorderChars(vert, horz, tl, tr, bl, br rune)
```
SetBorderChars will set the border characters

#### func (*WindowWriter) SetBordered

```go
func (ww *WindowWriter) SetBordered(on bool)
```
SetBordered turns the border on / off

#### func (*WindowWriter) SetPadding

```go
func (ww *WindowWriter) SetPadding(top, right, bottom, left int)
```
SetPadding will set the window padding

#### func (*WindowWriter) TermWriter

```go
func (ww *WindowWriter) TermWriter() *TermWriter
```
TermWriter will return the term writer

#### func (WindowWriter) WriteStringAt

```go
func (ww WindowWriter) WriteStringAt(line, col int, s string) (int, error)
```
WriteStringAt will write a string relative to the window

#### func (WindowWriter) WriteStyledStringAt

```go
func (ww WindowWriter) WriteStyledStringAt(line, col int, fgd, bck palette.Color, s string) (int, error)
```
WriteStyledStringAt will write a styled string at a location a negative value
for any r,g,b will not set fgd, bkd where it appears
