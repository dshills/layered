package terminal

import (
	"fmt"
	"io"

	"github.com/dshills/layered/palette"
)

const (
	esc = byte(0x1b)
)

// Color Modes
const (
	ColorModeTrueColor = iota
	ColorMode256
	ColorMode16
	ColorMode8
)

// TermWriter writes to a terminal
type TermWriter struct {
	w         io.Writer
	colorMode int
}

// SetColorMode will set the writers color mode
// Used in SetFgdCode and SetBckCode methods
func (tw *TermWriter) SetColorMode(clrmode int) {
	tw.colorMode = clrmode
}

// Write is a io.Writer
func (tw TermWriter) Write(b []byte) (int, error) {
	return tw.w.Write(b)
}

// WriteString will write a string at the current cursor position
func (tw TermWriter) WriteString(s string) (int, error) {
	return tw.w.Write([]byte(s))
}

// WriteRune will write a rune at the current cursor position
func (tw TermWriter) WriteRune(r rune) (int, error) {
	return tw.w.Write([]byte(string(r)))
}

// WriteEscString will write an escape followed by the string
func (tw TermWriter) WriteEscString(s string) (int, error) {
	tw.Write([]byte{esc})
	return tw.Write([]byte(s))
}

// SetFgdCode will set the foreground color
// based on the color mode 8, 16, 256, truecolor
func (tw TermWriter) SetFgdCode(c int) {
	switch tw.colorMode {
	case ColorModeTrueColor:
		fallthrough
	case ColorMode256:
		tw.WriteEscString(fmt.Sprintf("[38;5;%dm", c))
	case ColorMode16:
		tw.WriteEscString(fmt.Sprintf("[%d;1m", c))
	case ColorMode8:
		tw.WriteEscString(fmt.Sprintf("[%dm", c))
	}

}

// SetBckCode will set the background color
// based on the color mode 8, 16, 256, truecolor
func (tw TermWriter) SetBckCode(c int) {
	switch tw.colorMode {
	case ColorModeTrueColor:
		fallthrough
	case ColorMode256:
		tw.WriteEscString(fmt.Sprintf("[38;5;%dm", c))
	case ColorMode16:
		tw.WriteEscString(fmt.Sprintf("[%d;1m", c))
	case ColorMode8:
		tw.WriteEscString(fmt.Sprintf("[%dm", c))
	}
}

// SetStyle will set the foreground and backgroud color
func (tw TermWriter) SetStyle(fgd, bck palette.Color) {
	tw.SetForeground(fgd)
	tw.SetBackground(bck)
}

// SetForeground will set the foreground color
func (tw TermWriter) SetForeground(c palette.Color) {
	if c.Transparent {
		return
	}
	tw.WriteEscString(fmt.Sprintf("[38;2;%d;%d;%dm", c.Red, c.Green, c.Blue))
}

// SetBackground will set the background color
func (tw TermWriter) SetBackground(c palette.Color) {
	if c.Transparent {
		return
	}
	tw.WriteEscString(fmt.Sprintf("[48;2;%d;%d;%dm", c.Red, c.Green, c.Blue))
}

// ResetStyle will clear any styles set
func (tw TermWriter) ResetStyle() {
	tw.WriteEscString("[0m")
}

// Home moves the cursor to 0, 0
func (tw TermWriter) Home() {
	tw.WriteEscString("[H")
}

// Up will move the cursor up cnt lines
func (tw TermWriter) Up(cnt int) {
	tw.WriteEscString(fmt.Sprintf("[%dA", cnt))
}

// Down will move the cursor down cnt lines
func (tw TermWriter) Down(cnt int) {
	tw.WriteEscString(fmt.Sprintf("[%dB", cnt))
}

// Forward will advance the cursor cnt columns
func (tw TermWriter) Forward(cnt int) {
	tw.WriteEscString(fmt.Sprintf("[%dC", cnt))
}

// Back will move the cursor back cnt columns
func (tw TermWriter) Back(cnt int) {
	tw.WriteEscString(fmt.Sprintf("[%dD", cnt))
}

// SavePos will save the current cursor position
func (tw TermWriter) SavePos() {
	tw.WriteEscString("[s")
}

// RestorePos will restore the cursor to the saved position
func (tw TermWriter) RestorePos() {
	tw.WriteEscString("[u")
}

// ToColumn will move the cursor to column c of the current line
func (tw TermWriter) ToColumn(c int) {
	tw.WriteEscString(fmt.Sprintf("[%dG", c))
}

// BeginningLineUp will move to the beginning of count lines above
func (tw TermWriter) BeginningLineUp(cnt int) {
	tw.WriteEscString(fmt.Sprintf("[%dF", cnt))
}

// BeginningLineDown will move to the beginning of count lines below
func (tw TermWriter) BeginningLineDown(cnt int) {
	tw.WriteEscString(fmt.Sprintf("[%dE", cnt))
}

// To will move to line, column
func (tw TermWriter) To(line, col int) {
	tw.Home()
	tw.WriteEscString(fmt.Sprintf("[%d;%df", line, col))
}

// Clear will clear the entire screen
func (tw TermWriter) Clear() {
	//tw.WriteEscString("[J")
	tw.WriteEscString("[2J")
}

// ClearToEnd will clear from cursor to end of screen
func (tw TermWriter) ClearToEnd() {
	tw.WriteEscString("[0J")
}

// ClearToStart will clear from cursor to start of screen
func (tw TermWriter) ClearToStart() {
	tw.WriteEscString("[1J")
}

// ClearLine will clear the current line
func (tw TermWriter) ClearLine() {
	tw.WriteEscString("[K")
}

// ClearToEndOfLine will clear to EOL
func (tw TermWriter) ClearToEndOfLine() {
	tw.WriteEscString("[0K")
}

// ClearToStartOfLine will clear to BOL
func (tw TermWriter) ClearToStartOfLine() {
	tw.WriteEscString("[1K")
}

// ClearEntireLine will clear the line
func (tw TermWriter) ClearEntireLine() {
	tw.WriteEscString("[2K")
}

// NewTermWriter will return a new terminal writer
func NewTermWriter(w io.Writer, clrmode int) *TermWriter {
	return &TermWriter{w: w, colorMode: clrmode}
}
