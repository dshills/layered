package screen

import (
	"strings"
	"sync"
)

type statusbar struct {
	hidden     bool
	pos        BarPos
	components []SBComponent
	values     map[string]string
	m          sync.RWMutex
}

func (sb *statusbar) SetHidden(h bool)          { sb.hidden = h }
func (sb *statusbar) Hidden() bool              { return sb.hidden }
func (sb *statusbar) SetPosition(bp BarPos)     { sb.pos = bp }
func (sb *statusbar) Position() BarPos          { return sb.pos }
func (sb *statusbar) Components() []SBComponent { return sb.components }
func (sb *statusbar) SetComponentValue(key string, val string) {
	sb.m.Lock()
	defer sb.m.Unlock()
	sb.values[key] = val
}
func (sb *statusbar) ComponentValue(key string) string {
	sb.m.RLock()
	defer sb.m.RUnlock()
	return sb.values[key]
}
func (sb *statusbar) Add(cp ...SBComponent) {
	sb.components = append(sb.components, cp...)
}
func (sb *statusbar) Remove(key string) {
	idx := -1
	for i := range sb.components {
		if sb.components[i].Key == key {
			idx = i
			break
		}
	}
	if idx == -1 {
		return
	}
	sb.components = append(sb.components[:idx], sb.components[idx+1:]...)
}

func (sb *statusbar) Output(width int) string {
	left := strings.Builder{}
	right := strings.Builder{}
	for _, comp := range sb.components {
		if !comp.Right {
			left.WriteString(comp.Pre)
			left.WriteString(sb.ComponentValue(comp.Key))
			left.WriteString(comp.Post)
		} else {
			right.WriteString(comp.Pre)
			right.WriteString(sb.ComponentValue(comp.Key))
			right.WriteString(comp.Post)
		}
	}

	pad := ""
	if left.Len()+right.Len() < width {
		pad = strings.Repeat(" ", width-left.Len()+right.Len())
	}
	return left.String() + pad + right.String()
}

// NewStatusbar will return a statusbar
func NewStatusbar() Statusbar {
	return &statusbar{
		pos:    Bottom,
		values: make(map[string]string),
	}
}
