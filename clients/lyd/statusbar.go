package main

import (
	"fmt"
	"image"
	"strings"
	"sync"
	"time"

	"github.com/dshills/layered/palette"
	"github.com/gdamore/tcell"
)

// Layer component default types
const (
	SBLayer    = "layer"
	SBFilename = "filename"
	SBFiletype = "filetype"
	SBDirty    = "dirty"
	SBLine     = "line"
	SBColumn   = "column"
	SBClock    = "clock"
	SBPercent  = "percent"
	SBNumLines = "numlines"
	SBBufNum   = "bufnum"
)

// Draw will draw the status bar
func (s *Statusbar) Draw() {
	if s.hidden {
		return
	}
	s.Clear()
	x := s.region.Min.X
	for _, it := range s.items {
		if !it.Right {
			st := s.itemStyle(it.Key, s.values[it.Key], s.style)
			x = s.drawString(x, s.region.Min.Y, it.ValStr(s.values[it.Key]), st)
		}
	}
	x = s.region.Max.X - 1
	for i := len(s.items) - 1; i >= 0; i-- {
		it := s.items[i]
		if it.Right {
			st := s.itemStyle(it.Key, s.values[it.Key], s.style)
			str := it.ValStr(s.values[it.Key])
			x -= len(str)
			s.drawString(x, s.region.Min.Y, str, st)
		}
	}
	s.screen.Show()
}

func (s *Statusbar) itemStyle(key, value string, def tcell.Style) tcell.Style {
	s.sm.RLock()
	defer s.sm.RUnlock()
	st := def
	ss, ok := s.styles[key]
	if ok {
		st = ss
	}
	ss, ok = s.styles[fmt.Sprintf("%v#%v", key, value)]
	if ok {
		st = ss
	}
	return st
}

// Clear will clear the status bar
func (s *Statusbar) Clear() {
	for x := s.region.Min.X; x < s.region.Max.X; x++ {
		s.screen.SetContent(x, s.region.Min.Y, ' ', nil, s.style)
	}
}

// AddItem will add an item to the status bar
func (s *Statusbar) AddItem(it SBItem, pos int) {
	switch {
	case pos > len(s.items):
		s.items = append(s.items, it)
	default:
		s.items = append(s.items, SBItem{})
		copy(s.items[pos+1:], s.items[pos:])
		s.items[pos] = it
	}
	s.Draw()
}

// SetValue will set a statusbar value
func (s *Statusbar) SetValue(key, val string) {
	s.m.Lock()
	defer s.m.Unlock()
	old := s.values[key]
	s.values[key] = val
	if s.hasItem(key) && old != val {
		s.Draw()
	}
}

// RemoveValue will remove a value
func (s *Statusbar) RemoveValue(key string) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.values, key)
}

// SetHidden will hide or show the status bar
func (s *Statusbar) SetHidden(h bool) {
	s.hidden = h
	s.Draw()
}

// Palette will load the palette colors
func (s *Statusbar) Palette(pal *palette.Palette) {
	s.sm.Lock()
	defer s.sm.Unlock()
	ents := pal.HasPrefix("sbar-")
	for _, ent := range ents {
		name := strings.TrimPrefix(ent.Name, "sbar-")
		s.styles[name] = entryToStyle(ent)
	}
}

func (s *Statusbar) clock() {
	tic := time.NewTicker(20 * time.Second)
	defer tic.Stop()
	for {
		select {
		case <-s.done:
			return
		case <-tic.C:
			s.SetValue("clock", time.Now().Format(time.Kitchen))
		}
	}
}

func (s *Statusbar) hasItem(key string) bool {
	for _, it := range s.items {
		if it.Key == key {
			return true
		}
	}
	return false
}

func (s *Statusbar) drawString(x, y int, str string, st tcell.Style) int {
	for _, r := range str {
		s.screen.SetContent(x, y, r, nil, st)
		x++
	}
	return x
}

// NewStatusbar will return a status bar
func NewStatusbar(sc tcell.Screen, rgn image.Rectangle, done chan struct{}) *Statusbar {
	sb := &Statusbar{
		done:   done,
		screen: sc,
		region: rgn,
		values: make(map[string]string),
		styles: make(map[string]tcell.Style),
	}
	go sb.clock()
	return sb
}

// Statusbar is the statusbar
type Statusbar struct {
	hidden bool
	items  []SBItem
	values map[string]string
	m      sync.RWMutex
	done   chan struct{}
	region image.Rectangle
	screen tcell.Screen
	style  tcell.Style
	styles map[string]tcell.Style
	sm     sync.RWMutex
}

// ValueStyle is used when the style should change as the
// statusbar item changes value
type ValueStyle struct {
	Value string
	Bck   *palette.Color
	Fgd   *palette.Color
}

// SBItem is a statusbar item
type SBItem struct {
	Key   string
	Right bool
	Pre   string
	Post  string
}

// ValStr will return the formatted the value string
func (i *SBItem) ValStr(val string) string {
	return fmt.Sprintf("%s%s%s", i.Pre, val, i.Post)
}

// NewSBItem will return a status bar item
func NewSBItem(key, pre, post string, right bool) SBItem {
	return SBItem{Key: key, Pre: pre, Post: post, Right: right}
}
