package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/dshills/layered/editor"
	"github.com/dshills/layered/palette"
	"github.com/dshills/layered/terminal"
)

type screen struct {
	pal           *palette.Palette
	ed            editor.Editorer
	tw            *terminal.TermWriter
	width, length int
	windows       []window
}

func (s *screen) draw() {
	for i := range s.windows {
		s.windows[i].draw()
	}
}

func (s *screen) newWindow(bufid string) {
	s.windows = append(s.windows, newWindow(bufid, s.tw, 0, 0, s.width, s.length-2, s.ed, s.pal))
}

func (s *screen) notice(str string) {
	s.tw.To(s.length, 0)
	s.tw.SetForeground(palette.NewColor(255, 255, 255))
	s.tw.WriteString(strings.Repeat(" ", s.width))
	s.tw.To(s.length, 2)
	s.tw.WriteString(str)
}

func (s *screen) status(str string) {
	s.tw.To(s.length-1, 0)
	s.tw.SetBackground(palette.NewColor(40, 44, 52))
	s.tw.SetForeground(palette.NewColor(202, 194, 183))
	s.tw.WriteString(strings.Repeat(" ", s.width))
	s.tw.To(s.length-1, 2)
	s.tw.WriteString(str)
}

func newScreen(ed editor.Editorer, pal *palette.Palette) (*screen, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	fmt.Println(string(out))
	sz := strings.TrimSpace(string(out))
	splits := strings.Split(sz, " ")
	if len(splits) != 2 {
		return nil, fmt.Errorf("Could not get term size")
	}
	splits[0] = strings.TrimSpace(splits[0])
	splits[1] = strings.TrimSpace(splits[1])

	l, err := strconv.Atoi(splits[0])
	if err != nil {
		return nil, fmt.Errorf("Failed to get length")
	}
	w, err := strconv.Atoi(splits[1])
	if err != nil {
		return nil, fmt.Errorf("Failed to get width")
	}

	tw := terminal.NewTermWriter(os.Stdin, terminal.ColorModeTrueColor)
	sc := &screen{
		pal:    pal,
		ed:     ed,
		width:  w,
		length: l,
		tw:     tw,
	}
	tw.Clear()
	sc.notice("This is a notice")
	sc.status("This is a status")

	return sc, nil
}
