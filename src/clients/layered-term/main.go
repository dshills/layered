package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/buffer"
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/editor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/logger"
	"github.com/dshills/layered/palette"
	"github.com/dshills/layered/register"
	"github.com/dshills/layered/syntax"
	"github.com/dshills/layered/textobject"
	"github.com/dshills/layered/textstore"
	"github.com/dshills/layered/undo"
	"github.com/eiannone/keyboard"
)

const rtpath = "/Users/dshills/Development/projects/layered/runtime"

func main() {
	f, err := os.Create("./layered.log")
	if err != nil {
		logger.ErrorErr(err)
		os.Exit(1)
	}
	defer f.Close()
	log.SetOutput(f)
	logger.Message("Starting layered")

	if err := keyboard.Open(); err != nil {
		logger.ErrorErr(err)
		os.Exit(1)
	}
	defer keyboard.Close()

	logger.Message("Creating editor")
	ed, err := editor.New(undo.New, textstore.New, buffer.New, cursor.New, syntax.New, filetype.New, textobject.New, register.New, rtpath)
	if err != nil {
		logger.ErrorErr(err)
		os.Exit(1)
	}

	logger.Message("Loading colors")
	colors := palette.NewColorList()
	cp := filepath.Join(rtpath, "config", "colors.json")
	if err = colors.Load(cp); err != nil {
		logger.ErrorErr(err)
	}
	logger.Message("Loading palette")
	pal := palette.NewPalette()
	cp = filepath.Join(rtpath, "config", "palette.json")
	if err = pal.Load(cp, &colors); err != nil {
		logger.ErrorErr(err)
	}

	logger.Message("Creating screen")
	sc, err := newScreen(ed, &pal)
	if err != nil {
		logger.ErrorErr(err)
		os.Exit(1)
	}

	logger.Message("Loading file")
	trans := action.NewTransaction("")
	act := action.New(action.OpenFile)
	act.SetParam("/Users/dshills/Development/projects/goed-core/testdata/scanner.go")
	trans.Set(act)
	resp, err := ed.Exec(trans)
	if err != nil {
		logger.ErrorErr(err)
		os.Exit(1)
	}
	trans.SetBuffer(resp.Buffer)
	bufid := resp.Buffer

	act = action.New(action.BufferList)
	trans.Set(act)
	resp, err = ed.Exec(trans)
	if err != nil {
		logger.ErrorErr(err)
	}
	for i := range resp.Results {
		logger.Debugf("%+v", resp.Results[i])
	}

	act = action.New(action.Content)
	act.SetLine(0)
	act.SetCount(25)
	trans.Set(act)
	resp, err = ed.Exec(trans)
	if err != nil {
		logger.ErrorErr(err)
	}
	logger.Debugf("Content len %v", len(resp.Content))

	sc.newWindow(bufid)
	sc.draw()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			logger.ErrorErr(err)
			continue
		}
		//logger.Debugf("%v %v", char, key)
		if key == keyboard.KeyEsc {
			break
		}
	}
}
