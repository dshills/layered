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

type interpreter struct {
	status       MatchStatus
	layers       []Layer
	active       Layer
	prevPartial  []key.Keyer
	partial      []key.Keyer
	actDefs      action.Definitions
	layerStack   []Layer
	defaultLayer string
}

func (in *interpreter) Match(ks ...key.Keyer) ([]action.Action, error) {
	return in._match([]action.Action{}, ks)
}

func (in *interpreter) _match(pacts []action.Action, ks []key.Keyer) ([]action.Action, error) {
	acts := pacts
	if in.active == nil {
		in.status = ErrorMatch
		return acts, fmt.Errorf("No active layer")
	}
	if len(ks) == 0 {
		in.status = ErrorMatch
		return acts, fmt.Errorf("Nothing to process")
	}

	in.addPartial(ks...)

	mi := in.active.Match(in.partial...)
	in.status = mi.Status
	//logger.Debugf("_match: %q => %q Val: %q Rem: %q %+v", in.partial, in.status, in.partToString(mi.MatchValue), in.partToString(mi.Remaining), mi.Actions)

	if in.status != PartialMatch {
		in.clearPartial()
	}
	la := in.fortifyActions(mi.Actions, in.partToString(mi.MatchValue))
	acts = append(acts, la...)
	if len(mi.Remaining) > 0 {
		return in._match(acts, mi.Remaining)
	}
	return acts, nil
}

func (in *interpreter) fortifyActions(acts []action.Action, part string) []action.Action {
	pre := []action.Action{}
	post := []action.Action{}
	active := in.active
	for i, act := range acts {
		if act.Target == "input" {
			acts[i].Target = part
		}
		if act.Name == action.ChangeLayer {
			if err := in.changeLayer(act.Target); err != nil {
				logger.Errorf("fortifyActions: %v", err)
				continue
			}
			pre = append(pre, active.OnExitLayer()...)
			post = append(post, in.active.OnEnterLayer()...)
		}
		if act.Name == action.ChangePrevLayer {
			if err := in.changeLayer("previous"); err != nil {
				logger.Errorf("fortifyActions: %v", err)
				continue
			}
			pre = append(pre, active.OnExitLayer()...)
			post = append(post, in.active.OnEnterLayer()...)
		}
	}
	fin := make([]action.Action, len(pre)+len(acts)+len(post), len(pre)+len(acts)+len(post))
	idx := 0
	for i := range pre {
		fin[idx] = pre[i]
		idx++
	}
	for i := range acts {
		fin[idx] = acts[i]
		idx++
	}
	for i := range post {
		fin[idx] = post[i]
		idx++
	}
	return fin
}

func (in *interpreter) changeLayer(n string) error {
	if strings.ToLower(n) == "previous" {
		lay := in.pop()
		if lay == nil {
			return fmt.Errorf("Layer stack is empty")
		}
		in.active = lay
		return nil
	}

	lay := in.getLayer(n)
	if lay == nil {
		return fmt.Errorf("Layer not found")
	}
	in.push(in.active)
	in.active = lay
	return nil
}

func (in *interpreter) push(l Layer) {
	if l == nil {
		return
	}
	if len(in.layerStack) == 0 {
		in.layerStack = append(in.layerStack, l)
		return
	}
	if in.layerStack[len(in.layerStack)-1].Name() == l.Name() {
		return
	}
	if l.NotStacked() {
		return
	}
	in.layerStack = append(in.layerStack, l)
}

func (in *interpreter) pop() Layer {
	if len(in.layerStack) == 0 {
		return nil
	}
	var x Layer
	x, in.layerStack = in.layerStack[len(in.layerStack)-1], in.layerStack[:len(in.layerStack)-1]
	return x
}

func (in *interpreter) getLayer(n string) Layer {
	for _, l := range in.layers {
		if l.Name() == n {
			return l
		}
	}
	return nil
}

func (in *interpreter) Layers() []Layer {
	return in.layers
}

func (in *interpreter) Active() Layer {
	return in.active
}

func (in *interpreter) Partial() string {
	return in.partToString(in.partial)
}

func (in *interpreter) partToString(p []key.Keyer) string {
	builder := strings.Builder{}
	for _, p := range p {
		builder.WriteString(p.String())
	}
	return builder.String()
}

func (in *interpreter) Status() MatchStatus {
	return in.status
}

func (in *interpreter) Add(ll ...Layer) {
	in.layers = append(in.layers, ll...)
}

func (in *interpreter) Remove(name string) {
	for i := range in.layers {
		if in.layers[i].Name() == name {
			in.layers = append(in.layers[:i], in.layers[i+1:]...)
			return
		}
	}
}

func (in *interpreter) LoadDirectory(dir string) error {
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
			if err := lay.Load(in.actDefs, file); err != nil {
				errs = append(errs, err.Error())
				continue
			}
			in.Add(&lay)
		}
	}
	if in.active == nil {
		in.changeLayer(in.defaultLayer)
	}
	if len(errs) > 0 {
		return fmt.Errorf("Interpriter.LoadDirectory: %v", strings.Join(errs, ", "))
	}
	return nil
}

func (in *interpreter) addPartial(ks ...key.Keyer) {
	in.partial = append(in.partial, ks...)
}

func (in *interpreter) clearPartial() {
	copy(in.prevPartial, in.partial)
	in.partial = []key.Keyer{}
}

// NewInterpreter returns an interpreter
func NewInterpreter(ad action.Definitions, deflayer string) Interpriter {
	return &interpreter{actDefs: ad, defaultLayer: deflayer}
}
