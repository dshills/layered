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
	Name string
	Fgd  Color
	Bck  Color
}

// Palette is a color palette
type Palette struct {
	entries map[string]Entry
	m       sync.RWMutex
}

// Entry will return a palette entry by name
func (p *Palette) Entry(name string) (Entry, error) {
	org := name
	p.m.RLock()
	defer p.m.RUnlock()
	for name != "" {
		e, ok := p.entries[strings.ToLower(name)]
		if !ok {
			name = generalize(name)
			continue
		}
		return e, nil
	}
	return Entry{}, fmt.Errorf("Palette.Entry: %v Not found", org)
}

// HasPrefix will return palette items with names starting with pre
func (p *Palette) HasPrefix(pre string) []Entry {
	p.m.RLock()
	defer p.m.RUnlock()
	pre = strings.ToLower(pre)
	ents := []Entry{}
	for k, v := range p.entries {
		if strings.HasPrefix(k, pre) {
			ents = append(ents, v)
		}
	}
	return ents
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

	none := Color{}
	none.Transparent = true
	for _, pal := range js {
		e := Entry{Name: pal.Name}
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
func NewPalette() *Palette {
	return &Palette{entries: make(map[string]Entry)}
}

type jsPalette []struct {
	Name       string `json:"name"`
	Foreground string `json:"foreground"`
	Background string `json:"background"`
	Modifier   string `json:"modifier,omitempty"`
}
