package layer

import (
	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

type keyAction struct {
	k []key.Keyer
	a []action.Action
}

// OldLayer is a keyboard action map
type OldLayer struct {
	n                    string
	ba, ea, pma, nma, ma []action.Action
	ka                   []keyAction
	def                  bool
}

// Name returns the layer's name
func (l *OldLayer) Name() string { return l.n }

// Add will map keys to actions
func (l *OldLayer) Add(keys []key.Keyer, actions []action.Action) {
	l.ka = append(l.ka, keyAction{k: keys, a: actions})
}

// Remove will remove a key mapping
func (l *OldLayer) Remove(keys []key.Keyer) {}

// NewParser returns a new key parser
func (l *OldLayer) NewParser() Parserer {
	return &Parser{l: l}
}

// BeginActions returns actions that occur when switching to the layer
func (l *OldLayer) BeginActions() []action.Action { return l.ba }

// EndActions returns action that occur when switching away from layer
func (l *OldLayer) EndActions() []action.Action { return l.ea }

// PartialMatchActions returns actions that occur when a partial match is made
func (l *OldLayer) PartialMatchActions() []action.Action { return l.pma }

// NoMatchActions returns actions the occur when a match is not made
func (l *OldLayer) NoMatchActions() []action.Action { return l.nma }

// MatchActions returns actions that occur when a match is made
// they are in addition to key mapped actions
func (l *OldLayer) MatchActions() []action.Action { return l.ma }

// IsDefault returns true if this is the default layer
func (l *OldLayer) IsDefault() bool { return l.def }

// Parser is a key stroke parser specific to a layer
type Parser struct {
	l       *OldLayer
	partial []key.Keyer
}

// Parse will take key strokes and will return actions when matches
func (p *Parser) Parse(keys ...key.Keyer) (actions []action.Action, status ParseStatus) {
	p.partial = append(p.partial, keys...)
	for i := range p.l.ka {
		switch SameKeys(p.partial, p.l.ka[i].k) {
		case NoMatch:
		case PartialMatch:
			return p.l.pma, PartialMatch
		case Match:
			status = Match
			actions = append(actions, p.l.ka[i].a...)
			actions = append(actions, p.l.MatchActions()...)
			return
		}
	}
	return p.l.nma, NoMatch
}

// SameKeys compares two key lists
func SameKeys(a, b []key.Keyer) ParseStatus {
	if len(a) != len(b) {
		for i := range a {
			if i >= len(b) {
				return NoMatch
			}
		}
		return PartialMatch
	}
	for i := range a {
		if !a[i].IsEqual(b[i]) {
			return NoMatch
		}
	}
	return Match
}
