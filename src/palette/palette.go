package palette

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Entry is a palette entry
type Entry struct {
	Fgd Color
	Bck Color
}

// Palette is a color palette
type Palette struct {
	entries map[string]Entry
	m       sync.RWMutex
}

// Entry will return a palette entry by name
func (p *Palette) Entry(name string) (Entry, error) {
	if name == "" {
		return Entry{}, fmt.Errorf("Palette.Entry: %v Not found", name)
	}
	p.m.RLock()
	defer p.m.RUnlock()
	e, ok := p.entries[strings.ToLower(name)]
	if !ok {
		return p.Entry(generalize(name))
	}
	return e, nil
}

func generalize(name string) string {
	i := len(name) - 1
	for i >= 0 && []rune(name)[i] != '.' {
		i--
	}
	if i >= 0 {
		return string([]rune(name)[:i])
	}
	return ""
}

// Add will add an entry to the palette
func (p *Palette) Add(name string, e Entry) {
	p.m.Lock()
	defer p.m.Unlock()
	p.entries[strings.ToLower(name)] = e
}

// Load will load a palette
func (p *Palette) Load(path string, cl *ColorList) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	js := jsPalette{}
	if err = json.NewDecoder(f).Decode(&js); err != nil {
		return err
	}
	errs := []string{}
	for i := range js.Colors {
		clr, err := parseRGB(js.Colors[i].Color)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		cl.Add(js.Colors[i].Name, clr)
	}

	none := Color{}
	none.Transparent = true
	for _, pal := range js.Palette {
		e := Entry{}
		if strings.ToLower(pal.Foreground) != "none" {
			clr, err := cl.Color(pal.Foreground)
			if err != nil {
				errs = append(errs, err.Error())
				continue
			}
			e.Fgd = clr
		} else {
			e.Fgd = none
		}
		if strings.ToLower(pal.Background) != "none" {
			clr, err := cl.Color(pal.Background)
			if err != nil {
				errs = append(errs, err.Error())
				continue
			}
			e.Bck = clr
		} else {
			e.Bck = none
		}
		p.Add(pal.Name, e)
	}

	if len(errs) > 0 {
		return fmt.Errorf("Palette.Load: %v", strings.Join(errs, ", "))
	}
	return nil
}

// NewPalette will return a palette
func NewPalette() Palette {
	return Palette{entries: make(map[string]Entry)}
}

type jsPalette struct {
	Colors  []jsColor `json:"colors"`
	Palette []struct {
		Name       string `json:"name"`
		Foreground string `json:"foreground"`
		Background string `json:"background"`
		Modifier   string `json:"modifier"`
	} `json:"palette"`
}
