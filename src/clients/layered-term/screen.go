package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/dshills/layered/editor"
	"github.com/dshills/layered/key"
	"github.com/dshills/layered/layer"
	"github.com/dshills/layered/logger"
	"github.com/dshills/layered/palette"
	"github.com/dshills/layered/terminal"
	"github.com/eiannone/keyboard"
)

type info struct {
	key   string
	value string
}

type screen struct {
	pal           *palette.Palette
	ed            editor.Editorer
	tw            *terminal.TermWriter
	width, length int
	windows       []window
	activeBufID   string
	noticePrefix  string
	respC         chan editor.Response
}

func (s *screen) handleResponse() {
	for {
		resp := <-s.respC
		if resp.Err != nil {
			logger.ErrorErr(resp.Err)
			s.notice(resp.Err.Error())
			continue
		}
		if resp.Status == layer.Match {
			logger.Debugf("%v %v %v", resp.Buffer, resp.Line, resp.Column)
		}
		if resp.Buffer != "" {
			s.activeBufID = resp.Buffer
		}
		if resp.NewBuffer {
			s.newWindow(resp.Buffer)
		}
		if resp.ContentChanged {
			s.draw()
		}
		if resp.CursorChanged {
			idx := s.winIdx(resp.Buffer)
			if idx != -1 {
				s.windows[idx].drawCursor(resp.Line, resp.Column)
			}
		}
		switch strings.ToUpper(resp.Layer) {
		case "COLON":
			s.noticePrefix = ":"
		case "SEARCH":
			s.noticePrefix = "/"
		default:
			s.noticePrefix = ""
		}
		s.status(strings.ToUpper(resp.Layer))
		s.notice(resp.Partial)
	}
}

func (s *screen) processKeys() {
	keys := []key.Keyer{}
	s.respC = make(chan editor.Response)
	keyC := s.ed.KeyChan()
	s.ed.SetRespChan(s.respC)
	go s.handleResponse()
	for {
		r, k, err := keyboard.GetKey()
		if err != nil {
			logger.ErrorErr(err)
			continue
		}
		if k == 32 {
			k = 0
			r = ' '
		}
		keyer := key.New(r, int(k))
		if keyer == nil {
			logger.Errorf("Unknown key press %v %v", r, k)
			continue
		}
		keyC <- keyer
		keys = append(keys, keyer)
		if k == keyboard.KeyHome {
			break
		}
	}
}

func (s *screen) winIdx(id string) int {
	for i := range s.windows {
		if s.windows[i].bufid == id {
			return i
		}
	}
	return -1
}

func (s *screen) draw() {
	for i := range s.windows {
		if err := s.windows[i].draw(); err != nil {
			logger.Errorf("draw: %v", err)
		}
	}
}

func (s *screen) newWindow(bufid string) {
	s.windows = append(s.windows, newWindow(bufid, s.tw, 0, 0, s.width, s.length-2, s.ed, s.pal))
}

func (s *screen) notice(str string) {
	s.tw.SavePos()
	str = s.noticePrefix + str
	s.tw.To(s.length, 0)
	s.tw.SetForeground(palette.NewColor(255, 255, 255))
	s.tw.WriteString(strings.Repeat(" ", s.width))
	s.tw.To(s.length, 2)
	s.tw.WriteString(str)
	s.tw.RestorePos()
}

func (s *screen) status(str string) {
	s.tw.SavePos()
	s.tw.To(s.length-1, 0)
	s.tw.SetBackground(palette.NewColor(40, 44, 52))
	s.tw.SetForeground(palette.NewColor(202, 194, 183))
	s.tw.WriteString(strings.Repeat(" ", s.width))
	s.tw.To(s.length-1, 2)
	s.tw.WriteString(str)
	s.tw.RestorePos()
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
	sc.notice("")
	sc.status("NORMAL")
	return sc, nil
}

func partialToString(keys []key.Keyer) string {
	builder := strings.Builder{}
	for i := range keys {
		r := keys[i].Rune()
		if r > 0 {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
