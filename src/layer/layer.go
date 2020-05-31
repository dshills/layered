package layer

import (
	"encoding/json"
	"io"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

type keyAction struct {
	k []key.Keyer
	a []action.Actioner
}

// Layer is a keyboard action map
type Layer struct {
	n                    string
	ba, ea, pma, nma, ma []action.Actioner
	ka                   []keyAction
}

// Name returns the layer's name
func (l *Layer) Name() string { return l.n }

// Add will map keys to actions
func (l *Layer) Add(keys []key.Keyer, actions []action.Actioner) {
	l.ka = append(l.ka, keyAction{k: keys, a: actions})
}

// Remove will remove a key mapping
func (l *Layer) Remove(keys []key.Keyer) {}

// NewParser returns a new key parser
func (l *Layer) NewParser() Parserer {
	return &Parser{l: l}
}

// BeginActions returns actions that occur when switching to the layer
func (l *Layer) BeginActions() []action.Actioner { return l.ba }

// EndActions returns action that occur when switching away from layer
func (l *Layer) EndActions() []action.Actioner { return l.ea }

// PartialMatchActions returns actions that occur when a partial match is made
func (l *Layer) PartialMatchActions() []action.Actioner { return l.pma }

// NoMatchActions returns actions the occur when a match is not made
func (l *Layer) NoMatchActions() []action.Actioner { return l.nma }

// MatchActions returns actions that occur when a match is made
// they are in addition to key mapped actions
func (l *Layer) MatchActions() []action.Actioner { return l.ma }

// Load will load a layer from a reader
func (l *Layer) Load(r io.Reader) error {
	return json.NewDecoder(r).Decode(l)
}

// Parser is a key stroke parser specific to a layer
type Parser struct {
	l       *Layer
	partial []key.Keyer
}

// Parse will take key strokes and will return actions when matches
func (p *Parser) Parse(keys ...key.Keyer) (actions []action.Actioner, status ParseStatus) {
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

type jsLayer struct {
	Name                string        `json:"name"`
	OnBeginActions      []jsAction    `json:"on_begin_actions"`
	OnEndActions        []jsAction    `json:"on_end_actions"`
	NoMatchActions      []jsAction    `json:"no_match_actions"`
	PartialMatchActions []jsAction    `json:"partial_match_actions"`
	MatchActions        []jsAction    `json:"match_actions"`
	Commands            []jsKeyAction `json:"commands"`
}

func (l *jsLayer) asLayer() Layerer {
	lay := Layer{n: l.Name}
	for _, act := range l.OnBeginActions {
		lay.ba = append(lay.ba, act.asAction())
	}
	for _, act := range l.OnEndActions {
		lay.ea = append(lay.ea, act.asAction())
	}
	for _, act := range l.NoMatchActions {
		lay.nma = append(lay.nma, act.asAction())
	}
	for _, act := range l.PartialMatchActions {
		lay.pma = append(lay.pma, act.asAction())
	}
	for _, act := range l.MatchActions {
		lay.ma = append(lay.ma, act.asAction())
	}
	for _, kact := range l.Commands {
		lay.ka = append(lay.ka, kact.asKeyAction())
	}

	return &lay
}

type jsAction struct {
	Action string `json:"action"`
	Target string `json:"target"`
	Param  string `json:"param"`
}

func (a jsAction) asAction() action.Actioner {
	act := action.New(a.Action)
	act.SetTarget(a.Target)
	act.SetParam(a.Param)
	return act
}

type jsKeyAction struct {
	Keys    []string   `json:"keys"`
	Actions []jsAction `json:"actions"`
}

func (ka jsKeyAction) asKeyAction() keyAction {
	nka := keyAction{}
	for i := range ka.Actions {
		nka.a = append(nka.a, ka.Actions[i].asAction())
	}
	for i := range ka.Keys {
		nka.k = append(nka.k, key.StrToKey(ka.Keys[i]))
	}
	return nka
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
