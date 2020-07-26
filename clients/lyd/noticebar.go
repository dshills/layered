package main

import (
	"image"

	"github.com/gdamore/tcell"
)

// Notice will set the notice bar message
func (n *Noticebar) Notice(msg string) {
	n.draw(msg, n.style)
}

// Draw will draw the status bar
func (n *Noticebar) draw(msg string, st tcell.Style) {
	if n.hidden {
		return
	}
	n.clear()
	n.drawString(n.region.Min.X+1, n.region.Min.Y, msg, st)
	n.screen.Show()
}

// Clear will clear the status bar
func (n *Noticebar) clear() {
	for x := n.region.Min.X; x < n.region.Max.X; x++ {
		n.screen.SetContent(x, n.region.Min.Y, ' ', nil, n.style)
	}
}

// SetHidden will hide or show the status bar
func (n *Noticebar) SetHidden(h bool) {
	n.hidden = h
	n.clear()
}

func (n *Noticebar) drawString(x, y int, str string, st tcell.Style) int {
	for _, r := range str {
		n.screen.SetContent(x, y, r, nil, st)
		x++
	}
	return x
}

// NewNoticebar will return a status bar
func NewNoticebar(sc tcell.Screen, rgn image.Rectangle) *Noticebar {
	return &Noticebar{screen: sc, region: rgn, style: tcell.StyleDefault}
}

// Noticebar is the statusbar
type Noticebar struct {
	hidden  bool
	region  image.Rectangle
	screen  tcell.Screen
	style   tcell.Style
	message string
}
