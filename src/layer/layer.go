package layer

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
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

// Map will add a keys / actions mapping
func (l *Layer) Map(name string, keys []string, actions []action.Action) error {
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

// Unmap will remove a mapping
func (l *Layer) Unmap(name string) {
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

// Load will load a layer from a reader
func (l *Layer) Load(r io.Reader) error {
	js := jsLayer{}
	if err := json.NewDecoder(r).Decode(&js); err != nil {
		return err
	}
	errs := []string{}
	l.name = js.Name
	for _, a := range js.OnBeginActions {
		act, err := a.asAction()
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		l.beginActions = append(l.beginActions, act)
	}
	for _, a := range js.OnEndActions {
		act, err := a.asAction()
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		l.endActions = append(l.endActions, act)
	}
	for _, a := range js.NoMatchActions {
		act, err := a.asAction()
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		l.noMatchActions = append(l.noMatchActions, act)
	}
	for _, a := range js.PartialMatchActions {
		act, err := a.asAction()
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		l.partialActions = append(l.partialActions, act)
	}
	for _, kact := range js.Commands {
		kas, err := kact.asKeyAction()
		if err != nil {
			errs = append(errs, err.Error())
		}
		l.keyActs = append(l.keyActs, kas)
	}

	if len(errs) > 0 {
		return fmt.Errorf("%v", strings.Join(errs, ", "))
	}
	return nil
}

type jsLayer struct {
	Name                string        `json:"name"`
	Default             bool          `json:"default"`
	OnBeginActions      []jsAction    `json:"on_begin_actions"`
	OnEndActions        []jsAction    `json:"on_end_actions"`
	NoMatchActions      []jsAction    `json:"no_match_actions"`
	PartialMatchActions []jsAction    `json:"partial_match_actions"`
	Commands            []jsKeyAction `json:"commands"`
}

type jsAction struct {
	Action string `json:"action"`
	Target string `json:"target"`
}

func (a jsAction) asAction() (action.Action, error) {
	act, err := action.StrToAction(a.Action)
	if err != nil {
		return act, fmt.Errorf("%v", a.Action)
	}
	act.Target = a.Target
	act.Count = 1
	return act, nil
}

type jsKeyAction struct {
	Name    string     `json:"name"`
	Keys    []string   `json:"keys"`
	Actions []jsAction `json:"actions"`
}

func (ka jsKeyAction) asKeyAction() (keyAct, error) {
	actions := []action.Action{}
	errs := []string{}
	for _, a := range ka.Actions {
		act, err := a.asAction()
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		actions = append(actions, act)
	}
	keys := []keyMatch{}
	for i := range ka.Keys {
		km, err := newKeyMatch(ka.Keys[i])
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		keys = append(keys, km)
	}
	if len(errs) > 0 {
		return newKeyAct(keys, actions), fmt.Errorf("asKeyAction: %v", strings.Join(errs, ", "))
	}
	return newKeyAct(keys, actions), nil
}
