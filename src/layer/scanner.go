package layer

import (
	"fmt"
	"strings"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
	"github.com/dshills/layered/logger"
)

const targetInput = "input"

// Scanner evaluates keys within a layer
type Scanner struct {
	def     string
	layers  Manager
	prev    []string
	current Layerer
	partial []key.Keyer
}

// Init will initialize the scanner
func (s *Scanner) Init() {
	s.partial = []key.Keyer{}
}

// Partial returns the partial keys
func (s *Scanner) Partial() string {
	return s.keyerToTarget(s.partial)
}

// LayerName will return name of the current layer
func (s *Scanner) LayerName() string {
	return s.current.Name()
}

// Scan will match keys in the current layer
func (s *Scanner) Scan(key key.Keyer) ([]action.Action, ParseStatus, error) {
	s.partial = append(s.partial, key)
	//logger.Debugf("%+v", s.partial)
	acts, status := s.current.Match(s.partial)
	switch status {
	case Match:
		acts = s.processActions(acts)
		s.Init()
		return acts, status, nil
	case NoMatch:
		acts = s.processActions(s.current.NoMatchActions())
		s.Init()
		return acts, status, nil
	case PartialMatch:
		acts = s.processActions(s.current.PartialMatchActions())
		return acts, status, nil
	}
	return nil, NoMatch, fmt.Errorf("Unexpected match status")
}

func (s *Scanner) log(key key.Keyer, st ParseStatus, acts []action.Action) {
	pp := ""
	if len(s.prev) > 0 {
		pp = s.prev[len(s.prev)-1]
	}
	logger.Debugf("Scanner.Scan:\nStatus: %v\nLayer: %v, Prev: %v\nKey: %v, Partial: %q\nActions: %v\n", st, s.current.Name(), pp, key, s.keyerToTarget(s.partial), acts)
}

func (s *Scanner) processActions(acts []action.Action) []action.Action {
	na := []action.Action{}
	idx := 0
	for {
		if idx >= len(acts) {
			break
		}
		name := strings.ToLower(acts[idx].Name)
		target := strings.ToLower(acts[idx].Target)
		switch {
		case name == action.ChangeLayer:
			s.layerChange(acts[idx])
		case name == action.RunCommand:
			tar := target
			if target == targetInput {
				tar = s.keyerToTarget(s.partial)
			}
			na = append(na, s.processCommand(tar)...)
		case target == targetInput:
			acts[idx].Target = s.keyerToTarget(s.partial)
			na = append(na, acts[idx])
		default:
			na = append(na, acts[idx])
		}
		idx++
	}
	return na
}

func (s *Scanner) processCommand(target string) []action.Action {
	target = strings.TrimSpace(target)
	splits := strings.Split(target, " ")
	act, err := action.StrToAction(splits[0])
	if err != nil {
		return nil
	}
	if len(splits) > 1 {
		act.Target = splits[1]
	}
	return []action.Action{act}
}

func (s *Scanner) keyerToTarget(keys []key.Keyer) string {
	builder := strings.Builder{}
	for i := range keys {
		r := keys[i].Rune()
		if r > 0 {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func (s *Scanner) layerChange(act action.Action) {
	nl := strings.ToLower(act.Target)
	if nl == "prev" {
		nl = s.prevNot(s.current.Name())
	}
	s.setPrev(s.current.Name())

	lay, err := s.layers.Layer(nl)
	if err != nil {
		return
	}
	s.current = lay
}

func (s *Scanner) prevNot(lay string) string {
	for i := len(s.prev) - 1; i >= 0; i-- {
		if s.prev[i] != lay {
			return s.prev[i]
		}
	}
	return s.def
}

func (s *Scanner) setPrev(lay string) {
	lp := len(s.prev)
	if lp == 0 {
		s.prev = append(s.prev, lay)
		return
	}
	if s.prev[lp-1] == lay {
		return
	}
	s.prev = append(s.prev, lay)
	lp++
	if lp > 10 {
		s.prev = s.prev[lp-10:]
	}
}

// NewScanner returns a layer scanner
func NewScanner(layers Manager, stLayer string) (*Scanner, error) {
	lay, err := layers.Layer(stLayer)
	if err != nil {
		return nil, err
	}
	return &Scanner{layers: layers, current: lay, def: stLayer}, nil
}
