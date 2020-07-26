package main

import (
	"github.com/dshills/layered/palette"
	"github.com/gdamore/tcell"
)

func styleEntry(pal *palette.Palette, name string, def tcell.Style) tcell.Style {
	ent, err := pal.Entry(name)
	if err != nil {
		return def
	}
	return entryToStyle(ent)
}

func colorConv(clr palette.Color) tcell.Color {
	r, g, b, _ := clr.RGBA()
	return tcell.NewRGBColor(int32(r), int32(g), int32(b))
}

func entryToStyle(ent palette.Entry) tcell.Style {
	st := tcell.StyleDefault
	if !ent.Bck.Transparent {
		st = st.Background(colorConv(ent.Bck))
	}
	if !ent.Fgd.Transparent {
		st = st.Foreground(colorConv(ent.Fgd))
	}
	return st
}
