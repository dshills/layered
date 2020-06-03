package layer

import (
	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
	"github.com/dshills/layered/logger"
)

// Layer is a mapping of key strokes to actions
type Layer struct {
	name           string
	keyActs        []keyAct
	noMatchActions []action.Action
	partialActions []action.Action
	beginActions   []action.Action
	endActions     []action.Action
}

func actCopy(acts []action.Action) []action.Action {
	a := make([]action.Action, len(acts))
	copy(a, acts)
	return a
}

// Match will attempt to map keys to actions
func (l *Layer) Match(keys []key.Keyer) ([]action.Action, ParseStatus) {
	hasPartial := false
	logger.Debugf("Layer.Match: %v", l.name)
	for i := range l.keyActs {
		switch l.keyActs[i].match(keys) {
		case Match:
			return actCopy(l.keyActs[i].actions), Match
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
func (l *Layer) Add(keys []string, actions []action.Action) error {
	kms := []keyMatch{}
	for _, s := range keys {
		km, err := newKeyMatch(s)
		if err != nil {
			return err
		}
		kms = append(kms, km)
	}
	l.keyActs = append(l.keyActs, keyAct{matchers: kms, actions: actions})
	return nil
}

// Remove will remove a mapping
func (l *Layer) Remove(keys []string) {
	// TODO
}

// BeginActions will return the actions that are returned when switching to layer
func (l *Layer) BeginActions() []action.Action { return actCopy(l.beginActions) }

// EndActions will return the actions that are returned when leaving layer
func (l *Layer) EndActions() []action.Action { return actCopy(l.endActions) }

// PartialMatchActions returns the partial match actions
func (l *Layer) PartialMatchActions() []action.Action { return actCopy(l.partialActions) }

// NoMatchActions returns actions when keys do not match
func (l *Layer) NoMatchActions() []action.Action { return actCopy(l.noMatchActions) }

type jsLayer struct {
	Name                string        `json:"name"`
	Default             bool          `json:"default"`
	OnBeginActions      []jsAction    `json:"on_begin_actions"`
	OnEndActions        []jsAction    `json:"on_end_actions"`
	NoMatchActions      []jsAction    `json:"no_match_actions"`
	PartialMatchActions []jsAction    `json:"partial_match_actions"`
	Commands            []jsKeyAction `json:"commands"`
}

func (l *jsLayer) asLayer() (Layerer, error) {
	lay := Layer{name: l.Name}
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
		kas, err := kact.asKeyAction()
		if err != nil {
			return &lay, err
		}
		lay.keyActs = append(lay.keyActs, kas)
	}

	return &lay, nil
}

type jsAction struct {
	Action string `json:"action"`
	Target string `json:"target"`
}

func (a jsAction) asAction() action.Action {
	return action.Action{
		Name:   a.Action,
		Target: a.Target,
	}
}

type jsKeyAction struct {
	Keys    []string   `json:"keys"`
	Actions []jsAction `json:"actions"`
}

func (ka jsKeyAction) asKeyAction() (keyAct, error) {
	actions := []action.Action{}
	for i := range ka.Actions {
		actions = append(actions, ka.Actions[i].asAction())
	}
	keys := []keyMatch{}
	for i := range ka.Keys {
		km, err := newKeyMatch(ka.Keys[i])
		if err != nil {
			return keyAct{}, err
		}
		keys = append(keys, km)
	}
	return newKeyAct(keys, actions), nil
}
