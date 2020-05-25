package syntax

import (
	"io"
	"regexp"
)

// Rule types
const (
	MatchRule   = "match"
	RegionRule  = "region"
	KeywordRule = "keyword"
)

// Rule is a syntax matching rules
type Rule struct {
	grp, typ                        string
	st, end, mat, sm, sk            string
	contains, keywords              []string
	contained, display, trans, fold bool
	stRx, enRx                      *regexp.Regexp
	priority                        int
}

// Group will return the rules group
func (r *Rule) Group() string {
	return r.grp
}

// Type returns the rule type
func (r *Rule) Type() string {
	return r.typ
}

// Match will return match results
func (r *Rule) Match(s io.Reader, ln int) Resulter {
	return nil
}
