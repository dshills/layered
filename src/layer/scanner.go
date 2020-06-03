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
	layers  Collectioner
	prev    []string
	current Layerer
	partial []key.Keyer
}

// Init will initialize the scanner
func (s *Scanner) Init() {
	s.partial = []key.Keyer{}
}

// Scan will match keys in the current layer
func (s *Scanner) Scan(key key.Keyer) ([]action.Action, ParseStatus, error) {
	s.partial = append(s.partial, key)
	acts, status := s.current.Match(s.partial)
	switch status {
	case Match:
		acts = s.needTarget(acts)
		s.layerChange(acts)
		s.Init()
		return acts, status, nil
	case NoMatch:
		s.layerChange(s.current.NoMatchActions())
		s.Init()
		return s.current.NoMatchActions(), status, nil
	case PartialMatch:
		s.layerChange(s.current.PartialMatchActions())
		return s.current.PartialMatchActions(), status, nil
	}
	return nil, NoMatch, fmt.Errorf("Unexpected match status")
}

func (s *Scanner) needTarget(acts []action.Action) []action.Action {
	for i := range acts {
		if acts[i].Target == targetInput {
			acts[i].Target = s.keyerToTarget(s.partial)
		}
	}
	return acts
}

func (s *Scanner) keyerToTarget(keys []key.Keyer) string {
	builder := strings.Builder{}
	for i := range keys {
		r := keys[i].Rune()
		if r > 0 {
			builder.WriteRune(r)
		}
	}
	logger.Debugf("Scanner.keyerToTarget: %v", builder.String())
	return builder.String()
}

func (s *Scanner) layerChange(acts []action.Action) {
	for _, act := range acts {
		if strings.ToLower(act.Name) == action.ChangeLayer {
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
	}
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
func NewScanner(layers Collectioner, stLayer string) (*Scanner, error) {
	lay, err := layers.Layer(stLayer)
	if err != nil {
		return nil, err
	}
	return &Scanner{layers: layers, current: lay, def: stLayer}, nil
}
