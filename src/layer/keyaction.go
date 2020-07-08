package layer

import (
	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

type keyact struct {
	name    string
	keys    []key.Keyer
	actions []action.Action
}

func (k *keyact) Keys() []key.Keyer {
	return k.keys
}

func (k *keyact) Actions() []action.Action {
	return k.actions
}

func (k *keyact) Match(keys []key.Keyer) MatchStatus {
	if len(keys) > len(k.keys) {
		return NoMatch
	}
	for i := range keys {
		if !k.keys[i].Eq(keys[i]) {
			return NoMatch
		}
	}
	if len(keys) == len(k.keys) {
		return Match
	}
	return PartialMatch
}

// NewKeyAction will return a key action structure
func NewKeyAction(name string, keys []key.Keyer, acts []action.Action) KeyAction {
	return &keyact{
		name:    name,
		keys:    keys,
		actions: acts,
	}
}
