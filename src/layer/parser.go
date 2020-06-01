package layer

import (
	"fmt"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

// Scanner evaluates keys within a layer
type Scanner struct {
	layers  Collectioner
	prev    []string
	current Layerer
	partial []key.Keyer
	param   string
}

// Init will initialize the scanner
func (s *Scanner) Init() {
	s.param = ""
	s.partial = []key.Keyer{}
}

// Scan will match keys in the current layer
func (s *Scanner) Scan(key key.Keyer) ([]action.Action, ParseStatus, error) {
	if s.current.PartialAsParam() && s.current.PartialTrigger().IsEqual(key) {
		if s.current.PartialIncludeTrigger() {
			s.param += string(key.Rune())
		}
		acts := s.current.PartialMatchActions()
		if len(acts) == 0 {
			return nil, NoMatch, fmt.Errorf("Scanner.Scan: PartialAsParam set but no actions defined")
		}
		for i := range acts {
			if acts[i].Param == "" {
				acts[i].Param = s.param
			}
		}
		s.Init()
		s.layerChange(acts)
		return acts, Match, nil
	}

	s.partial = append(s.partial, key)
	acts, status := s.current.Match(s.partial)
	switch status {
	case Match:
		s.Init()
		s.layerChange(acts)
		return acts, status, nil
	case NoMatch:
		s.layerChange(s.current.NoMatchActions())
		return s.current.NoMatchActions(), status, nil
	case PartialMatch:
		s.layerChange(s.current.PartialMatchActions())
		return s.current.PartialMatchActions(), status, nil
	}
	return nil, NoMatch, fmt.Errorf("Unexpected match status")
}

func (s *Scanner) layerChange(acts []action.Action) {
	for _, act := range acts {
		if act.Name == action.ChangeLayer {
			s.setPrev(s.current.Name())
			lay, err := s.layers.Layer(act.Param)
			if err != nil {
				return
			}
			s.current = lay
		}
	}
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
func NewScanner(layers Collectioner) (*Scanner, error) {
	def := layers.Default()
	if def == nil {
		return nil, fmt.Errorf("Collectioner has no default layer")
	}
	return &Scanner{
		layers:  layers,
		current: def,
	}, nil
}
