package syntax

import (
	"regexp"

	"github.com/dshills/layered/logger"
	"github.com/dshills/layered/textstore"
)

// Rule types
const (
	MatchRule   = "match"
	RegionRule  = "region"
	KeywordRule = "keyword"
)

const (
	idxStart = 0
	idxEnd   = 1
)

// Rule is a syntax matching rules
type Rule struct {
	group, rtype, start, end, match, same, skip string
	contains, keywords                          []string
	contained, display, transparent, fold       bool
	stRx, enRx                                  *regexp.Regexp
	priority                                    int
}

// Group will return the rules group
func (r *Rule) Group() string {
	return r.group
}

// Type returns the rule type
func (r *Rule) Type() string {
	return r.rtype
}

// Match will return match results
func (r *Rule) Match(txt textstore.TextStorer) []Resulter {
	switch r.rtype {
	case KeywordRule:
		fallthrough
	case MatchRule:
		return r.matchSimple(txt)
	case RegionRule:
		if r.same == "" {
			return r.matchRegion(txt)
		}
		return r.matchRegionSame(txt)
	}
	return nil
}

func (r *Rule) matchSimple(txt textstore.TextStorer) []Resulter {
	results := []Resulter{}
	cnt := txt.NumLines()
	for i := 0; i < cnt; i++ {
		str, _ := txt.LineString(i)
		idxs := r.stRx.FindAllStringIndex(str, -1)
		if len(idxs) == 0 {
			continue
		}
		res := Result{ln: i, tok: r.group, rg: idxs, pr: r.priority}
		results = append(results, &res)
	}
	return results
}

func (r *Rule) matchRegion(txt textstore.TextStorer) []Resulter {
	if r.stRx == nil || r.enRx == nil {
		return nil
	}
	results := []Resulter{}
	open := 0
	cnt := txt.NumLines()
	for i := 0; i < cnt; i++ {
		str, _ := txt.LineString(i)
		stIdxs := r.stRx.FindAllStringIndex(str, -1)
		endIdxs := r.enRx.FindAllStringIndex(str, -1)
		if len(stIdxs) == 0 && len(endIdxs) == 0 {
			continue
		}
		res := Result{ln: i, tok: r.group, pr: r.priority}

		op := open
		for ii := 0; ii < op; ii++ {
			if len(endIdxs) > 0 {
				res.rg = append(res.rg, []int{0, endIdxs[0][idxEnd]})
				endIdxs = endIdxs[1:]
				open--
			}
		}

		ll := len(str)
		for _, sti := range stIdxs {
			if len(endIdxs) > 0 {
				res.rg = append(res.rg, []int{sti[idxStart], endIdxs[0][idxEnd]})
				endIdxs = endIdxs[1:]
			} else {
				res.rg = append(res.rg, []int{sti[idxStart], ll - 1})
				open++
			}
		}
		results = append(results, &res)
	}
	return results
}

func (r *Rule) matchRegionSame(txt textstore.TextStorer) []Resulter {
	if r.stRx == nil {
		return nil
	}
	results := []Resulter{}
	open := 0
	cnt := txt.NumLines()
	for i := 0; i < cnt; i++ {
		str, _ := txt.LineString(i)
		idxs := r.stRx.FindAllStringIndex(str, -1)
		if len(idxs) == 0 {
			continue
		}
		res := Result{ln: i, tok: r.group, pr: r.priority}

		op := open
		for ii := 0; ii < op; ii++ {
			if len(idxs) > 0 {
				res.rg = append(res.rg, []int{0, idxs[0][idxEnd]})
				idxs = idxs[1:]
				open--
			}
		}

		for {
			if len(idxs) < 2 {
				break
			}
			res.rg = append(res.rg, []int{idxs[0][idxStart], idxs[1][idxEnd]})
			idxs = idxs[2:]
		}

		if len(idxs) > 0 {
			res.rg = append(res.rg, []int{idxs[0][idxStart], len(str) - 1})
			open++
		}

		results = append(results, &res)
	}
	return results
}

func testResultsValid(rt string, results []Resulter) {
	for _, lnres := range results {
		for _, rg := range lnres.Range() {
			if len(rg) != 2 {
				logger.Debugf("%v Syntax Range len != 2 got %v", rt, len(rg))
				continue
			}
			if rg[0] > rg[1] {
				logger.Debugf("%v Syntax Range %v > %v", rt, rg[0], rg[1])
				continue
			}
		}
	}
}
