package layer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
	"github.com/dshills/layered/logger"
)

type interpriter struct {
	layers       []Layer
	partial      []key.Keyer
	lastStatus   MatchStatus
	stack        []Layer
	active       Layer
	lastPartial  string
	actDefs      action.Definitions
	defaultLayer string
}

func (i *interpriter) push(l Layer) {
	if len(i.stack) == 0 {
		i.stack = append(i.stack, l)
		return
	}
	if i.stack[len(i.stack)-1].Name() == l.Name() {
		return
	}
	if l.NotStacked() {
		return
	}
	i.stack = append(i.stack, l)
}

func (i *interpriter) pop() Layer {
	if len(i.stack) == 1 {
		return i.stack[0]
	}
	var l Layer
	l, i.stack = i.stack[len(i.stack)-1], i.stack[:len(i.stack)-1]
	return l
}

func (i *interpriter) Layers() []Layer { return i.layers }

func (i *interpriter) Active() Layer {
	return i.active
}

func (i *interpriter) addPartial(kk ...key.Keyer) {
	i.partial = append(i.partial, kk...)
}

func (i *interpriter) Partial() string {
	str := ""
	for _, k := range i.partial {
		if k.Rune() != 0 {
			str += string(k.Rune())
		}
	}
	return str
}
func (i *interpriter) Status() MatchStatus { return i.lastStatus }
func (i *interpriter) Add(ls ...Layer) {
	for _, l := range ls {
		i.layers = append(i.layers, l)
	}
}

func (i *interpriter) Remove(name string) {
	idx := -1
	for i, l := range i.layers {
		if l.Name() == name {
			idx = i
			break
		}
	}
	if idx != -1 {
		i.layers[idx] = i.layers[len(i.layers)-1]
		i.layers = i.layers[:len(i.layers)-1]
	}
}

func (i *interpriter) LoadDirectory(dir string) error {
	fi, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	errs := []string{}
	for _, f := range fi {
		if f.IsDir() {
			continue
		}
		if strings.ToLower(path.Ext(f.Name())) == ".json" {
			file, err := os.Open(path.Join(dir, f.Name()))
			if err != nil {
				errs = append(errs, err.Error())
				continue
			}
			defer file.Close()
			lay := layer{}
			if err := lay.Load(i.actDefs, file); err != nil {
				errs = append(errs, err.Error())
				continue
			}
			i.Add(&lay)
		}
	}
	if i.active == nil {
		i.active = i.getLayer(i.defaultLayer)
		i.push(i.active)
	}
	if len(errs) > 0 {
		return fmt.Errorf("Interpriter.LoadDirectory: %v", strings.Join(errs, ", "))
	}
	return nil
}

func (i *interpriter) Match(keys ...key.Keyer) []action.Action {
	if i.active == nil {
		logger.Errorf("No default layer set")
		return nil
	}
	if len(keys) == 0 {
		return nil
	}
	i.addPartial(keys...)

	ak := keys[len(keys)-1]
	acts, end := i.active.MatchSpecial(ak)
	if end {
		i.clearPartial()
		return i.processActions(acts)
	}

	aa, st := i.active.Match(i.partial)
	if st != PartialMatch {
		i.clearPartial()
	}
	acts = append(acts, aa...)

	return i.processActions(acts)
}

func (i *interpriter) processActions(acts []action.Action) []action.Action {
	active := i.Active()
	pre := []action.Action{}
	post := []action.Action{}
	for ii, act := range acts {
		if act.Name == action.ChangeLayer {
			i.clearPartial()
			pre = append(pre, active.OnExitLayer()...)
			active = i.getLayer(act.Target)
			if active == nil {
				active = i.stack[0]
			}
			i.push(i.active)
			i.active = active
			post = append(post, i.active.OnEnterLayer()...)
			edact := action.Action{Name: action.CursorMovePast, Target: "false"}
			if i.active.AllowCursorPastEnd() {
				edact.Target = "true"
			}
			// Add cursor in pre so edit commands have it
			pre = append(pre, edact)
		}

		if act.Target == "input" {
			logger.Debugf("Target==input, %v", i.lastPartial)
			acts[ii].Target = i.lastPartial
		}

		acts[ii].Line = -1
		acts[ii].Column = -1
		acts[ii].Count = 0
	}
	acts = append(pre, acts...)
	acts = append(acts, post...)
	return acts
}

func (i *interpriter) getLayer(name string) Layer {
	name = strings.ToLower(name)
	if name == "previous" {
		return i.pop()
	}
	for _, l := range i.layers {
		if l.Name() == name {
			return l
		}
	}
	return nil
}

func (i *interpriter) clearPartial() {
	i.lastPartial = i.Partial()
	i.partial = []key.Keyer{}
}

// NewInterpriter returns an interpriter
func NewInterpriter(ad action.Definitions, deflayer string) Interpriter {
	return &interpriter{actDefs: ad, defaultLayer: deflayer}
}
