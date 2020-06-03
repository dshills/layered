package layer

import (
	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

type keyAct struct {
	matchers []keyMatch
	actions  []action.Action
}

func (ka keyAct) match(keys []key.Keyer) ParseStatus {
	if len(keys) == 0 {
		return NoMatch
	}
	count := 0
	var cnt int
	mr := 0
	for _, km := range ka.matchers {
		mr++
		keys, cnt = km.match(keys)
		count += cnt
		switch {
		case len(keys) == 0 && mr == len(ka.matchers):
			// we've run all matchers and all keys
			return Match
		case cnt == 0:
			// Failed a match
			return NoMatch
		case len(keys) == 0:
			// we've run all the keys but not all the matchers
			return PartialMatch
		}
	}
	// we've run all the matchers but still have keys
	return NoMatch
}

func newKeyAct(kms []keyMatch, actions []action.Action) keyAct {
	ka := keyAct{matchers: kms, actions: actions}
	return ka
}
