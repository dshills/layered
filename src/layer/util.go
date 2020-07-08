package layer

import (
	"fmt"
	"strings"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

func convertActions(acts []actJSON) ([]action.Action, error) {
	actions := []action.Action{}
	errs := []string{}
	for _, a := range acts {
		act, err := action.StrToAction(a.Action)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		act.Target = a.Target
		actions = append(actions, act)
	}
	if len(errs) > 0 {
		return actions, fmt.Errorf("%v", strings.Join(errs, ", "))
	}
	return actions, nil
}

func convertKeys(kstrs []string) ([]key.Keyer, error) {
	keys := []key.Keyer{}
	errs := []string{}
	for _, k := range kstrs {
		akey, err := key.StrToKeyer(k)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		keys = append(keys, akey)
	}
	if len(errs) > 0 {
		return keys, fmt.Errorf("%v", strings.Join(errs, ", "))
	}
	return keys, nil
}
