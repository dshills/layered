package layer

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

type layer struct {
	name         string
	editable     bool
	wait         bool
	nostack      bool
	cancelKey    key.Keyer
	completeKey  key.Keyer
	prevlayerKey key.Keyer
	any          []action.Action
	anyprint     []action.Action
	anynonprint  []action.Action
	onEnter      []action.Action
	onExit       []action.Action
	onComplete   []action.Action
	onMatch      []action.Action
	onNoMatch    []action.Action
	onPartial    []action.Action
	keyActions   []KeyAction
}

func (l *layer) Name() string                                               { return l.name }
func (l *layer) Map(name string, keys []key.Keyer, actions []action.Action) {}
func (l *layer) UnMap(name string)                                          {}
func (l *layer) Editable() bool                                             { return l.editable }
func (l *layer) WaitForComplete() bool                                      { return l.wait }
func (l *layer) NotStacked() bool                                           { return l.nostack }
func (l *layer) CancelKey() key.Keyer                                       { return l.cancelKey }
func (l *layer) PrevLayerKey() key.Keyer                                    { return l.prevlayerKey }
func (l *layer) CompleteKey() key.Keyer                                     { return l.completeKey }

func (l *layer) OnAnyKey() []action.Action         { return l.any }
func (l *layer) OnPrintableKey() []action.Action   { return l.anyprint }
func (l *layer) OnNonPritableKey() []action.Action { return l.anyprint }
func (l *layer) OnEnterLayer() []action.Action     { return l.onEnter }
func (l *layer) OnExitLayer() []action.Action      { return l.onExit }
func (l *layer) OnComplete() []action.Action       { return l.onComplete }
func (l *layer) OnMatch() []action.Action          { return l.onMatch }
func (l *layer) OnNoMatch() []action.Action        { return l.onNoMatch }
func (l *layer) OnPartialMatch() []action.Action   { return l.onPartial }
func (l *layer) KeyActions() []KeyAction           { return l.keyActions }
func (l *layer) Load(r io.Reader) error {
	errs := []string{}
	var err error
	lay := layJSON{}
	if err := json.NewDecoder(r).Decode(&lay); err != nil {
		return fmt.Errorf("%v", err)
	}
	l.name = strings.ToLower(lay.Name)
	l.editable = lay.Editable
	l.wait = lay.WaitForComplete
	l.prevlayerKey, _ = key.StrToKeyer(lay.PrevLayerOnKey)
	l.cancelKey, _ = key.StrToKeyer(lay.CancelOnKey)
	l.completeKey, _ = key.StrToKeyer(lay.CompleteOnKey)

	l.any, err = convertActions(lay.OnAnyKey)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.anyprint, err = convertActions(lay.OnPrintableKey)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.anynonprint, err = convertActions(lay.OnNonPritableKey)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onEnter, err = convertActions(lay.OnEnterLayer)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onExit, err = convertActions(lay.OnExitLayer)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onComplete, err = convertActions(lay.OnComplete)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onMatch, err = convertActions(lay.OnMatch)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onNoMatch, err = convertActions(lay.OnNoMatch)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onPartial, err = convertActions(lay.OnPartialMatch)
	if err != nil {
		errs = append(errs, err.Error())
	}

	for _, ka := range lay.Actions {
		acts, err := convertActions(ka.Actions)
		if err != nil {
			errs = append(errs, err.Error())
		}
		keys, err := convertKeys(ka.Keys)
		if err != nil {
			errs = append(errs, err.Error())
		}
		if len(acts) > 0 && len(keys) > 0 {
			l.keyActions = append(l.keyActions, NewKeyAction(ka.Name, keys, acts))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("%v", strings.Join(errs, ", "))
	}

	return nil
}

func (l *layer) Match(keys []key.Keyer) ([]action.Action, MatchStatus) {
	status := NoMatch
	acts := []action.Action{}

	for _, ka := range l.keyActions {
		switch ka.Match(keys) {
		case Match:
			acts = append(acts, ka.Actions()...)
			status = Match
			break
		case PartialMatch:
			status = PartialMatch
		}
	}

	switch status {
	case Match:
		acts = append(acts, l.OnMatch()...)
	case NoMatch:
		acts = append(acts, l.OnNoMatch()...)
	case PartialMatch:
		acts = append(acts, l.OnPartialMatch()...)
	}

	return acts, status
}

func (l *layer) MatchSpecial(k key.Keyer) ([]action.Action, bool) {
	acts := []action.Action{}
	if l.cancelKey != nil && l.cancelKey.Eq(k) {
		return []action.Action{action.Action{Name: action.ChangeLayer, Target: "default"}}, true
	}
	if l.completeKey != nil && l.completeKey.Eq(k) {
		return l.onComplete, true
	}
	if l.prevlayerKey != nil && l.prevlayerKey.Eq(k) {
		return []action.Action{action.Action{Name: action.ChangeLayer, Target: "previous"}}, true
	}
	acts = append(acts, l.OnAnyKey()...)
	if unicode.IsPrint(k.Rune()) {
		acts = append(acts, l.OnPrintableKey()...)
	} else {
		acts = append(acts, l.OnNonPritableKey()...)
	}
	return acts, false
}

// NewLayer will return a Layer
func NewLayer(name string) Layer {
	return &layer{name: name}
}
