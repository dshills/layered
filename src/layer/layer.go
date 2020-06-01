package layer

import (
	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

type keyAct struct {
	keys    []key.Keyer
	actions []action.Action
}

func (ka keyAct) match(keys []key.Keyer) ParseStatus {
	if len(keys) > len(ka.keys) {
		return NoMatch
	}
	for i, k := range keys {
		if !ka.keys[i].IsEqual(k) {
			return NoMatch
		}
	}
	if len(keys) == len(ka.keys) {
		return Match
	}
	return PartialMatch
}

// Layer is a mapping of key strokes to actions
type Layer struct {
	name           string
	keyActs        []keyAct
	isDefault      bool
	partialAsParam bool
	papTrigger     key.Keyer
	includePap     bool
	noMatchActions []action.Action
	partialActions []action.Action
	beginActions   []action.Action
	endActions     []action.Action
}

// PartialAsParam returns true if the keys should be used as an action parameter
// requires a trigger key
func (l *Layer) PartialAsParam() bool { return l.partialAsParam }

// PartialIncludeTrigger will add the trigger to the param
func (l *Layer) PartialIncludeTrigger() bool { return l.includePap }

// PartialTrigger will trigger a match using previous partial keys as a parameter
func (l *Layer) PartialTrigger() key.Keyer { return l.papTrigger }

// Match will attempt to map keys to actions
func (l *Layer) Match(keys []key.Keyer) ([]action.Action, ParseStatus) {
	hasPartial := false
	for i := range l.keyActs {
		switch l.keyActs[i].match(keys) {
		case Match:
			return l.keyActs[i].actions, Match
		case PartialMatch:
			hasPartial = true
		}
	}
	if hasPartial {
		return nil, PartialMatch
	}
	return nil, NoMatch
}

// Name returns the layer name
func (l *Layer) Name() string { return l.name }

// Add will add a keys / actions map
func (l *Layer) Add(keys []key.Keyer, actions []action.Action) {
	l.keyActs = append(l.keyActs, keyAct{keys: keys, actions: actions})
}

// Remove will remove a mapping
func (l *Layer) Remove(keys []key.Keyer) {
	for i := range l.keyActs {
		if l.keyActs[i].match(keys) == Match {
			// Order not perserved
			l.keyActs[i] = l.keyActs[len(l.keyActs)-1]
			l.keyActs = l.keyActs[:len(l.keyActs)-1]
			return
		}
	}
}

// BeginActions will return the actions that are returned when switching to layer
func (l *Layer) BeginActions() []action.Action { return l.beginActions }

// EndActions will return the actions that are returned when leaving layer
func (l *Layer) EndActions() []action.Action { return l.endActions }

// PartialMatchActions returns the partial match actions
func (l *Layer) PartialMatchActions() []action.Action { return l.partialActions }

// NoMatchActions returns actions when keys do not match
func (l *Layer) NoMatchActions() []action.Action { return l.noMatchActions }

// IsDefault returns true if it is the default layer
func (l *Layer) IsDefault() bool { return l.isDefault }

type jsLayer struct {
	Name                string        `json:"name"`
	Default             bool          `json:"default"`
	OnBeginActions      []jsAction    `json:"on_begin_actions"`
	OnEndActions        []jsAction    `json:"on_end_actions"`
	NoMatchActions      []jsAction    `json:"no_match_actions"`
	PartialMatchActions []jsAction    `json:"partial_match_actions"`
	MatchActions        []jsAction    `json:"match_actions"`
	Commands            []jsKeyAction `json:"commands"`
}

func (l *jsLayer) asLayer() Layerer {
	lay := Layer{name: l.Name, isDefault: l.Default}
	for _, act := range l.OnBeginActions {
		lay.beginActions = append(lay.beginActions, act.asAction())
	}
	for _, act := range l.OnEndActions {
		lay.endActions = append(lay.endActions, act.asAction())
	}
	for _, act := range l.NoMatchActions {
		lay.noMatchActions = append(lay.noMatchActions, act.asAction())
	}
	for _, act := range l.PartialMatchActions {
		lay.partialActions = append(lay.partialActions, act.asAction())
	}
	for _, kact := range l.Commands {
		lay.keyActs = append(lay.keyActs, kact.asKeyAction())
	}

	return &lay
}

type jsAction struct {
	Action string `json:"action"`
	Target string `json:"target"`
	Param  string `json:"param"`
}

func (a jsAction) asAction() action.Action {
	return action.Action{
		Name:   a.Action,
		Target: a.Target,
		Param:  a.Param,
	}
}

type jsKeyAction struct {
	Keys    []string   `json:"keys"`
	Actions []jsAction `json:"actions"`
}

func (ka jsKeyAction) asKeyAction() keyAct {
	nka := keyAct{}
	for i := range ka.Actions {
		nka.actions = append(nka.actions, ka.Actions[i].asAction())
	}
	for i := range ka.Keys {
		nka.keys = append(nka.keys, key.StrToKey(ka.Keys[i]))
	}
	return nka
}
