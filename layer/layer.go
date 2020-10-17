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

// MatchInfo is information about the match
type MatchInfo struct {
	Actions    []action.Action
	Status     MatchStatus
	MatchValue []key.Keyer
	Remaining  []key.Keyer
}

type layer struct {
	name               string
	editable           bool // command is editable
	wait               bool // waitForComplete
	nostack            bool // don't add it to the layer stack
	allowCursorPastEnd bool // set cursor to allow past end
	alwaysPartialMatch bool // always return partial match till complete
	cancelKey          key.Keyer
	completeKey        key.Keyer
	prevlayerKey       key.Keyer
	any                []action.Action
	anyprint           []action.Action
	anynonprint        []action.Action
	onEnter            []action.Action
	onExit             []action.Action
	onComplete         []action.Action
	onMatch            []action.Action
	onNoMatch          []action.Action
	onPartial          []action.Action
	keyActions         []KeyAction
}

func (l *layer) Match(keys ...key.Keyer) MatchInfo {
	switch {
	case l.wait:
		return l.matchWait(keys...)
	case l.cancelKey != nil || l.completeKey != nil || l.prevlayerKey != nil:
		return l.matchWithSpecial(keys...)
	default:
		// No special keys, just test
		acts, status := l.testKeyActions(keys...)
		acts = append(acts, l.addMatchActions(status)...)
		return MatchInfo{Actions: acts, Status: status, MatchValue: keys}
	}
}

func (l *layer) matchWait(keys ...key.Keyer) MatchInfo {
	acts := []action.Action{}
	status := NoMatch
	for i, k := range keys {
		switch {
		case l.completeKey != nil && l.completeKey.Eq(k):
			// completed so test the keys
			acts, status = l.testKeyActions(keys[:i]...)
			acts = append(acts, l.onComplete...)
			acts = append(acts, l.addMatchActions(status)...)
			return MatchInfo{Actions: acts, Status: status, Remaining: keys[i+1:], MatchValue: keys[:i]}
		case l.cancelKey != nil && l.cancelKey.Eq(k):
			// remove actions and change layer
			acts = []action.Action{{Name: action.ChangeLayer, Target: "previous"}}
			return MatchInfo{Actions: acts, Status: CancelMatch}
		case l.prevlayerKey != nil && l.prevlayerKey.Eq(k):
			// remove actions and change layer
			acts = []action.Action{{Name: action.ChangeLayer, Target: "previous"}}
			return MatchInfo{Actions: acts, Status: CancelMatch}
		}
	}
	acts = append(acts, l.addMatchActions(status)...)
	return MatchInfo{Actions: acts, Status: PartialMatch, MatchValue: keys}
}

func (l *layer) matchWithSpecial(keys ...key.Keyer) MatchInfo {
	acts := []action.Action{}
	status := NoMatch
	val := []key.Keyer{}
	for i, k := range keys {
		switch {
		case l.completeKey != nil && l.completeKey.Eq(k):
			acts = append(acts, l.onComplete...)
			acts = append(acts, l.addMatchActions(status)...)
			if status == PartialMatch {
				status = Match
			}
			return MatchInfo{Actions: acts, Status: status, Remaining: keys[i+1:], MatchValue: val}
		case l.cancelKey != nil && l.cancelKey.Eq(k):
			// remove actions and change layer
			acts = []action.Action{{Name: action.ChangeLayer, Target: "previous"}}
			return MatchInfo{Actions: acts, Status: CancelMatch, MatchValue: []key.Keyer{k}}
		case l.prevlayerKey != nil && l.prevlayerKey.Eq(k):
			// remove actions and change layer
			acts = []action.Action{{Name: action.ChangeLayer, Target: "previous"}}
			return MatchInfo{Actions: acts, Status: CancelMatch, MatchValue: []key.Keyer{k}}
		}
		val = append(val, k)
		ta, st := l.testKeyActions(keys[:i+1]...)
		status = st
		acts = append(acts, ta...)
		if l.alwaysPartialMatch {
			status = PartialMatch
		}
		if status != PartialMatch {
			acts = append(acts, l.addMatchActions(status)...)
			return MatchInfo{Actions: acts, Status: status, MatchValue: val}
		}
	}

	acts = append(acts, l.addMatchActions(status)...)
	return MatchInfo{Actions: acts, Status: status, MatchValue: keys}
}

func (l *layer) testKeyActions(keys ...key.Keyer) ([]action.Action, MatchStatus) {
	status := NoMatch
	for _, ka := range l.keyActions {
		switch ka.Match(keys) {
		case Match:
			acts := ka.Actions()
			acts = append(acts, l.addMatchActions(Match)...)
			return acts, Match
		case PartialMatch:
			status = PartialMatch
		}
	}
	return nil, status
}

func (l *layer) addMatchActions(status MatchStatus) []action.Action {
	switch status {
	case Match:
		return l.OnMatch()
	case NoMatch:
		return l.OnNoMatch()
	case PartialMatch:
		return l.OnPartialMatch()
	}
	return nil
}

func (l *layer) MatchSpecial(k key.Keyer) ([]action.Action, bool) {
	acts := []action.Action{}
	if l.cancelKey != nil && l.cancelKey.Eq(k) {
		return []action.Action{{Name: action.ChangeLayer, Target: "previous"}}, true
	}
	if l.completeKey != nil && l.completeKey.Eq(k) {
		return l.onComplete, true
	}
	if l.prevlayerKey != nil && l.prevlayerKey.Eq(k) {
		return []action.Action{{Name: action.ChangeLayer, Target: "previous"}}, true
	}
	return acts, true
}

func (l *layer) MatchKeyType(k key.Keyer) []action.Action {
	acts := []action.Action{}
	acts = append(acts, l.OnAnyKey()...)
	if unicode.IsPrint(k.Rune()) {
		acts = append(acts, l.OnPrintableKey()...)
	} else {
		acts = append(acts, l.OnNonPritableKey()...)
	}
	return acts
}

func (l *layer) Name() string                                               { return l.name }
func (l *layer) Map(name string, keys []key.Keyer, actions []action.Action) {}
func (l *layer) UnMap(name string)                                          {}
func (l *layer) AllowCursorPastEnd() bool                                   { return l.allowCursorPastEnd }
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

func (l *layer) Load(defs action.Definitions, r io.Reader) error {
	errs := []string{}
	var err error
	lay := layJSON{}
	if err := json.NewDecoder(r).Decode(&lay); err != nil {
		return fmt.Errorf("%v", err)
	}
	l.name = strings.ToLower(lay.Name)
	l.editable = lay.Editable
	l.wait = lay.WaitForComplete
	l.alwaysPartialMatch = lay.AlwaysPartial
	l.prevlayerKey, _ = key.StrToKeyer(lay.PrevLayerOnKey)
	l.cancelKey, _ = key.StrToKeyer(lay.CancelOnKey)
	l.completeKey, _ = key.StrToKeyer(lay.CompleteOnKey)
	l.allowCursorPastEnd = lay.AllowCursorPastEnd

	l.any, err = convertActions(defs, lay.OnAnyKey)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.anyprint, err = convertActions(defs, lay.OnPrintableKey)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.anynonprint, err = convertActions(defs, lay.OnNonPritableKey)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onEnter, err = convertActions(defs, lay.OnEnterLayer)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onExit, err = convertActions(defs, lay.OnExitLayer)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onComplete, err = convertActions(defs, lay.OnComplete)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onMatch, err = convertActions(defs, lay.OnMatch)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onNoMatch, err = convertActions(defs, lay.OnNoMatch)
	if err != nil {
		errs = append(errs, err.Error())
	}
	l.onPartial, err = convertActions(defs, lay.OnPartialMatch)
	if err != nil {
		errs = append(errs, err.Error())
	}

	for _, ka := range lay.Actions {
		acts, err := convertActions(defs, ka.Actions)
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

// NewLayer will return a Layer
func NewLayer(name string) Layer {
	return &layer{name: name}
}
