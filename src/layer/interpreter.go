package layer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

type interpriter struct {
	layers     []Layer
	active     Layer
	partial    []key.Keyer
	lastStatus MatchStatus
	lstack     []Layer
	defLayer   Layer
}

func (i *interpriter) Layers() []Layer      { return i.layers }
func (i *interpriter) Active() Layer        { return i.active }
func (i *interpriter) Partial() []key.Keyer { return i.partial }
func (i *interpriter) Status() MatchStatus  { return i.lastStatus }
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
			if err := lay.Load(file); err != nil {
				errs = append(errs, err.Error())
				continue
			}
			i.Add(&lay)
		}
	}
	i.active = i.getLayer("default")
	if len(errs) > 0 {
		return fmt.Errorf("Interpriter.LoadDirectory: %v", strings.Join(errs, ", "))
	}
	return nil
}

func (i *interpriter) Match(keys ...key.Keyer) []action.Action {
	if len(keys) == 0 {
		return nil
	}
	i.partial = append(i.partial, keys...)

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
	for _, act := range acts {
		if act.Name == action.ChangeLayer {
			i.clearPartial()
			acts = append(acts, i.active.OnExitLayer()...)
			if act.Target == "previous" {
				i.active = i.prevLayer()
			} else {
				i.pushLayer(i.active)
				i.active = i.getLayer(act.Target)
				acts = append(acts, i.active.OnEnterLayer()...)
			}
		}
	}
	return acts
}

func (i *interpriter) prevLayer() Layer {
	if len(i.lstack) == 0 {
		return i.defLayer
	}
	return i.lstack[len(i.lstack)-1]
}

func (i *interpriter) getLayer(name string) Layer {
	name = strings.ToLower(name)
	if name == "prev" {
		return i.prevLayer()
	}
	for _, l := range i.layers {
		if l.Name() == name {
			return l
		}
	}
	return i.defLayer
}

func (i *interpriter) pushLayer(l Layer) {
	if l.NotStacked() {
		return
	}
	if i.active == l {
		return
	}
	i.lstack = append(i.lstack, l)
}

func (i *interpriter) clearPartial() {
	i.partial = []key.Keyer{}
}

// NewInterpriter returns an interpriter
func NewInterpriter() Interpriter {
	return &interpriter{}
}
