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
		//logger.Debugf("Match: NoMatch len %v %v %v", len(keys), len(k.keys), keys)
		return NoMatch
	}
	for i := range keys {
		if !k.keys[i].Eq(keys[i]) {
			//logger.Debugf("Match: NoMatch key %v %v %v", keys[i], k.keys[i], keys)
			return NoMatch
		}
	}
	if len(keys) == len(k.keys) {
		//logger.Debugf("Match: Match key %v %v", keys, k.keys)
		return Match
	}
	//logger.Debugf("Match: Partial %v", keys)
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
