package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dshills/layered/buffer"
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

	ed, err := editor.New(undo.New, textstore.New, buffer.New, cursor.New, syntax.New, filetype.New, textobject.New, register.New, layer.New, rtpath)
	if err != nil {
		logger.ErrorErr(err)
		os.Exit(1)
	}

	colors := palette.NewColorList()
	cp := filepath.Join(rtpath, "config", "colors.json")
	if err = colors.Load(cp); err != nil {
		logger.ErrorErr(err)
	}
	pal := palette.NewPalette()
	cp = filepath.Join(rtpath, "config", "palette.json")
	if err = pal.Load(cp, &colors); err != nil {
		logger.ErrorErr(err)
	}

	screen, err := newScreen(ed, &pal)
	if err != nil {
		logger.ErrorErr(err)
		os.Exit(1)
	}
	go screen.processKeys()
	<-screen.doneC
}
