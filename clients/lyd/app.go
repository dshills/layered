package main

import (
	"fmt"
	"image"
	"path/filepath"
	"strings"
	"time"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/conf"
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/editor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/layer"
	"github.com/dshills/layered/logger"
	"github.com/dshills/layered/palette"
	"github.com/dshills/layered/register"
	"github.com/dshills/layered/syntax"
	"github.com/dshills/layered/textobject"
	"github.com/dshills/layered/textstore"
	"github.com/dshills/layered/undo"
	"github.com/gdamore/tcell"
)

const (
	rt          = "/Users/dshills/Development/projects/layered/runtime"
	demof       = "/Users/dshills/Development/projects/layered/testdata/scanner_test.go"
	scrollLines = 20
)

var errDone = fmt.Errorf("DONE")

// App is the primary application
type App struct {
	done          chan struct{}
	status        *Statusbar
	notice        *Noticebar
	screen        tcell.Screen
	windows       []*Window
	current       *Window
	ed            editor.Editorer
	reqC          chan action.Request
	respC         chan editor.Response
	width, height int
	pal           *palette.Palette
	keyscanner    layer.Interpriter
	bnum          int
	actDefs       action.Definitions
}

func (a *App) init() error {
	logger.Debugf("Starting editor...")
	config := conf.New()
	config.AddRuntime(rt)
	var err error
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	a.screen, err = tcell.NewScreen()
	if err != nil {
		a.screen.Fini()
		return fmt.Errorf("App.init: NewScreen %v", err)
	}
	if err := a.screen.Init(); err != nil {
		a.screen.Fini()
		return fmt.Errorf("App.init: Screen init %v", err)
	}
	a.screen.EnableMouse()

	clrs := palette.NewColorList(config)
	clrs.Load(filepath.Join(rt, "config", "colors.json"))
	a.pal = palette.NewPalette()
	err = a.pal.Load(filepath.Join(rt, "palette", "default.json"), clrs)
	if err != nil {
		a.screen.Fini()
		return fmt.Errorf("App.init: %v", err)
	}

	a.actDefs = action.New()
	a.ed, err = editor.New(undo.New, textstore.New, buffer.New, cursor.New, syntax.New, filetype.New, textobject.New, register.New, a.actDefs, config)
	if err != nil {
		a.screen.Fini()
		return fmt.Errorf("App.init: Editor.New %v", err)
	}
	a.keyscanner = layer.NewInterpriter(a.actDefs, "normal")
	err = a.keyscanner.LoadDirectory(filepath.Join(rt, "layers"))
	if err != nil {
		a.screen.Fini()
		return fmt.Errorf("App.init: layer.LoadDirectory %v", err)
	}
	a.reqC = make(chan action.Request, 10)
	a.respC = make(chan editor.Response, 10)
	a.ed.ExecChan(a.reqC, a.respC, a.done)
	go a.listen()

	a.screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite))
	a.screen.Clear()

	a.width, a.height = a.screen.Size()
	a.status = NewStatusbar(a.screen, image.Rect(0, a.height-2, a.width, a.height-2), a.done)
	a.initStatusbar()
	a.notice = NewNoticebar(a.screen, image.Rect(0, a.height-1, a.width, a.height-1))
	a.notice.Notice("I'm a notice bar")

	a.reqC <- action.NewRequest("", action.Action{Name: action.NewBuffer})
	logger.Debugf("%v", a)
	return nil
}

func (a *App) initStatusbar() {
	st := tcell.StyleDefault.Foreground(tcell.NewRGBColor(255, 255, 255)).Background(tcell.NewRGBColor(15, 77, 118))
	a.status.style = styleEntry(a.pal, "sbar-default", st)
	a.status.Palette(a.pal)

	a.status.AddItem(SBItem{Key: SBLayer, Pre: " ", Post: " "}, 1000)
	a.status.AddItem(SBItem{Key: SBBufNum, Pre: " ", Post: " "}, 1000)
	a.status.AddItem(SBItem{Key: SBFilename, Pre: " ", Post: " "}, 1000)
	a.status.AddItem(SBItem{Key: SBDirty, Post: " "}, 1000)

	a.status.AddItem(SBItem{Key: SBFiletype, Right: true, Post: " "}, 1000)
	a.status.AddItem(SBItem{Key: SBLine, Right: true, Pre: " ", Post: ":"}, 1000)
	a.status.AddItem(SBItem{Key: SBColumn, Right: true, Post: " "}, 1000)
	a.status.AddItem(SBItem{Key: SBNumLines, Right: true, Pre: " ", Post: " "}, 1000)
	a.status.AddItem(SBItem{Key: SBPercent, Right: true, Pre: " ", Post: " "}, 1000)
	a.status.AddItem(SBItem{Key: SBClock, Right: true, Pre: " ", Post: " "}, 1000)

	// Default values
	a.status.SetValue(SBLayer, "DEFAULT")
	a.status.SetValue(SBFilename, "No file")
	a.status.SetValue(SBFiletype, "unknown")
	a.status.SetValue(SBLine, "1")
	a.status.SetValue(SBColumn, "1")
	a.status.SetValue(SBNumLines, "0")
	a.status.SetValue(SBPercent, "0%")
	a.status.SetValue(SBClock, time.Now().Format(time.Kitchen))
}

func logRequest(req action.Request) {
	logger.Debugf(fmt.Sprintf("Request: LineOffset: %v, LineCount: %v, Buf: %v", req.LineOffset, req.LineCount, req.BufferID))
	for i, act := range req.Actions {
		builder := strings.Builder{}
		builder.WriteString(fmt.Sprintf("%v. Action: %q, Line: %v, Col: %v", i+1, act.Name, act.Line, act.Column))
		if act.Count > 0 {
			builder.WriteString(fmt.Sprintf(", Count: %v", act.Count))
		}
		if act.Target != "" {
			builder.WriteString(fmt.Sprintf(", Target: %v", act.Target))
		}
		logger.Debugf(builder.String())
	}
}

func logResponse(resp editor.Response) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Response: %q, Line: %v, Col: %v", resp.Action, resp.Line, resp.Column))
	if resp.ContentChanged {
		builder.WriteString(", Content changed")
	}
	if resp.CursorChanged {
		builder.WriteString(", Cursor changed")
	}
	if resp.NewBuffer {
		builder.WriteString(", BufferID created")
	}
	if resp.CloseBuffer {
		builder.WriteString(", Close buffer")
	}
	if resp.InfoChanged {
		builder.WriteString(", Info Changed")
	}
	if resp.BufferID != "" {
		builder.WriteString(fmt.Sprintf(", Buf: %q", resp.BufferID))
	}
	if resp.Err != nil {
		builder.WriteString(fmt.Sprintf(", Err: %v", resp.Err))
	}
	logger.Debugf(builder.String())
}

// ProcessKeys will process keystrokes
func (a *App) ProcessKeys() error {
	defer a.screen.Fini()
	for {
		ev := a.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if err := a.handleKeyEvent(ev); err != nil {
				logger.Errorf("handleKeyEvent: %v", err)
				if err == errDone {
					return nil
				}
			}

		case *tcell.EventMouse:
			if err := a.handleMouseEvent(ev); err != nil {
				logger.Errorf("handleMouseEvent: %v", err)
			}

		case *tcell.EventResize:
			a.screen.Sync()
		}
	}
}

func (a *App) handleKeyEvent(ev *tcell.EventKey) error {
	switch ev.Key() {
	case tcell.KeyHome:
		a.quit()
		return errDone
	}
	k, err := convertKey(ev)
	if err != nil {
		logger.Errorf("handleKeyEvent: convertKey %v", err)
		return nil
	}
	if a.keyscanner == nil {
		panic("handleKeyEvent: no interpritter")
	}
	acts := a.keyscanner.Match(k)
	a.status.SetValue("layer", strings.ToUpper(a.keyscanner.Active().Name()))
	a.Notify(a.keyscanner.Partial())

	if len(acts) > 0 {
		req := action.Request{}
		req.Add(acts...)
		if a.current != nil {
			req.BufferID = a.current.id
			req.LineOffset = a.current.offset
			req.LineCount = a.current.region.Dy()
		}
		logRequest(req)
		a.reqC <- req
	}
	return nil
}

func (a *App) handleMouseEvent(ev *tcell.EventMouse) error {
	mask := ev.Buttons()
	switch {
	case mask == tcell.Button1:
		x, y := ev.Position()
		win := a.windowAtPoint(x, y)
		if win != nil {
			a.current = win
			ln, col := win.pointToPos(x, y)
			a.reqC <- action.NewRequest(a.windows[0].id, action.Action{Name: action.Move, Line: ln, Column: col})
		}
	case mask&tcell.WheelUp != 0:
		if a.current != nil {
			a.reqC <- action.NewRequest(a.windows[0].id, action.Action{Name: action.ScrollUp})
		}
	case mask&tcell.WheelDown != 0:
		if a.current != nil {
			a.reqC <- action.NewRequest(a.windows[0].id, action.Action{Name: action.ScrollDown})
		}
	}
	return nil
}

func (a *App) listen() {
	for {
		select {
		case <-a.done:
			return
		case resp := <-a.respC:
			logResponse(resp)
			if resp.Err != nil {
				logger.ErrorErr(resp.Err)
				a.Notify(resp.Err.Error())
				continue
			}
			go a.updateStatusResp(resp)
			switch {
			case resp.CloseBuffer:
				for i, win := range a.windows {
					if win.id == resp.BufferID {
						if win.id == a.current.id {
							a.current = nil
						}
						copy(a.windows[i:], a.windows[i+1:])
						a.windows[len(a.windows)-1] = nil
						a.windows = a.windows[:len(a.windows)-1]
						if len(a.windows) == 0 {
							a.reqC <- action.NewRequest("", action.Action{Name: action.NewBuffer})
						}
					}
				}
			case resp.NewBuffer:
				win := a.newWindow(image.Rect(0, 0, a.width, a.height-3))
				a.current = win
				win.SetResponse(resp)
				a.Notifyf("BufferID %v created", win.bnum)
				req := action.NewRequest(resp.BufferID, action.Action{Name: action.OpenFile, Target: demof})
				req.LineCount = a.height - 3
				logRequest(req)
				a.reqC <- req
			default:
				if win := a.window(resp.BufferID); win != nil {
					win.SetResponse(resp)
				}
			}
		}
	}
}

func (a *App) updateStatusResp(resp editor.Response) {
	if resp.BufferID == "" {
		return
	}
	if a.current != nil {
		a.status.SetValue(SBBufNum, fmt.Sprintf("%v", a.current.bnum))
	}
	a.status.SetValue(SBLine, fmt.Sprintf("%v", resp.Line))
	a.status.SetValue(SBColumn, fmt.Sprintf("%v", resp.Column))
	a.status.SetValue(SBNumLines, fmt.Sprintf("%v", resp.NumLines))
	if resp.NumLines == 0 {
		a.status.SetValue(SBPercent, "0%")
	} else {
		pct := (float64(resp.Line) / float64(resp.NumLines)) * 100
		a.status.SetValue(SBPercent, fmt.Sprintf("%.0f%%", pct))
	}
	ft := resp.Filetype
	if ft == "" {
		ft = "unknown"
	}
	fn := resp.Filename
	if fn == "" {
		fn = "No file"
	}
	a.status.SetValue(SBFiletype, ft)
	a.status.SetValue(SBFilename, fn)
	if resp.Dirty {
		a.status.SetValue(SBDirty, "+")
	} else {
		a.status.SetValue(SBDirty, "")
	}
}

func (a *App) windowAtPoint(x, y int) *Window {
	for _, win := range a.windows {
		if win.hasPoint(x, y) {
			return win
		}
	}
	return nil
}

func (a *App) quit() {
	a.done <- struct{}{} // Self
	a.done <- struct{}{} // Statusbar
	a.done <- struct{}{} // editor
}

// Notifyf set a formatted notification message
func (a *App) Notifyf(format string, i interface{}) {
	a.notice.Notice(fmt.Sprintf(format, i))
}

// Notify set a notification message
func (a *App) Notify(msg string) {
	a.notice.Notice(msg)
}

func (a *App) newWindow(rgn image.Rectangle) *Window {
	win := NewWindow(a.screen, rgn, a.pal, a.reqC)
	a.bnum++
	win.bnum = a.bnum
	a.windows = append(a.windows, win)
	return win
}

func (a *App) window(id string) *Window {
	for _, w := range a.windows {
		if w.id == id {
			return w
		}
	}
	return nil
}
