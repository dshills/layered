package layer

import (
	"strings"
	"unicode"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

type keyAct struct {
	name     string
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

type keyMatch struct {
	r        rune
	k        int
	matcher  string
	original string
}

func (km keyMatch) match(keys []key.Keyer) ([]key.Keyer, int) {
	if len(keys) == 0 {
		return keys, 0
	}
	if km.matcher != "" {
		cnt := 0
		for _, key := range keys {
			k, r := key.Key(), key.Rune()
			switch km.matcher {
			case Any:
				return nil, len(keys)
			case Printable:
				if k != 0 || !unicode.IsPrint(r) {
					return keys[cnt:], cnt
				}
			case Control:
				if (r != 0 || k == 0) && !unicode.IsControl(r) {
					return keys[cnt:], cnt
				}
			case Digit:
				if k > 0 || !unicode.IsDigit(r) {
					return keys[cnt:], cnt
				}
			case Letter:
				if k > 0 || !unicode.IsLetter(r) {
					return keys[cnt:], cnt
				}
			case Lower:
				if k > 0 || !unicode.IsLower(r) {
					return keys[cnt:], cnt
				}
			case Upper:
				if k > 0 || !unicode.IsUpper(r) {
					return keys[cnt:], cnt
				}
			case NonBlank:
				if k > 0 || unicode.IsSpace(r) {
					return keys[cnt:], cnt
				}
			case Pattern:
			}
			cnt++
		}
		return nil, cnt
	}
	//logger.Debugf("Testing: %v (%v), %v == %v (%v), %v", keys[0].Rune(), keys[0].Rune(), keys[0].Key(), km.r, km.r, km.k)
	if km.r == keys[0].Rune() && km.k == keys[0].Key() {
		return keys[1:], 1
	}
	return keys, 0
}

func newKeyMatch(str string) (keyMatch, error) {
	kl := strings.ToLower(str)
	switch kl {
	case Any:
		return keyMatch{r: 0, k: 0, matcher: kl, original: str}, nil
	case Printable:
		return keyMatch{r: 0, k: 0, matcher: kl, original: str}, nil
	case Control:
		return keyMatch{r: 0, k: 0, matcher: kl, original: str}, nil
	case Digit:
		return keyMatch{r: 0, k: 0, matcher: kl, original: str}, nil
	case Letter:
		return keyMatch{r: 0, k: 0, matcher: kl, original: str}, nil
	case Lower:
		return keyMatch{r: 0, k: 0, matcher: kl, original: str}, nil
	case Upper:
		return keyMatch{r: 0, k: 0, matcher: kl, original: str}, nil
	case NonBlank:
		return keyMatch{r: 0, k: 0, matcher: kl, original: str}, nil
	case Pattern:
		return keyMatch{r: 0, k: 0, matcher: kl, original: str}, nil
	}
	r, k, err := key.StrToKey(str)
	if err != nil {
		return keyMatch{original: str}, err
	}
	return keyMatch{r: r, k: k, original: str}, nil
}
