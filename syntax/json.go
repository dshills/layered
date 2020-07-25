package syntax

import (
	"fmt"
	"regexp"
	"strings"
)

type jsRule struct {
	Group       string   `json:"group"`
	Type        string   `json:"type"`
	Start       string   `json:"start,omitempty"`
	End         string   `json:"end,omitempty"`
	Contains    []string `json:"contains,omitempty"`
	Keywords    []string `json:"keywords,omitempty"`
	Match       string   `json:"match,omitempty"`
	Contained   bool     `json:"contained,omitempty"`
	Display     bool     `json:"display,omitempty"`
	Same        string   `json:"same,omitempty"`
	Skip        string   `json:"skip,omitempty"`
	Transparent bool     `json:"transparent,omitempty"`
	Fold        bool     `json:"fold,omitempty"`
}

func (jr jsRule) asRuler(rank int) (Ruler, error) {
	r := Rule{
		priority:    rank,
		group:       jr.Group,
		rtype:       jr.Type,
		start:       jr.Start,
		end:         jr.End,
		contains:    jr.Contains,
		keywords:    jr.Keywords,
		match:       jr.Match,
		contained:   jr.Contained,
		display:     jr.Display,
		same:        jr.Same,
		skip:        jr.Skip,
		transparent: jr.Transparent,
		fold:        jr.Fold,
	}
	switch jr.Type {
	case KeywordRule:
		match := fmt.Sprintf("\\b(%s)\\b", strings.Join(jr.Keywords, "|"))
		reg, err := regexp.Compile(match)
		if err != nil {
			return nil, fmt.Errorf("Rule %v %v %v", jr.Type, jr.Match, err)
		}
		r.stRx = reg

	case MatchRule:
		reg, err := regexp.Compile(jr.Match)
		if err != nil {
			return nil, fmt.Errorf("Rule %v %v %v", jr.Type, jr.Match, err)
		}
		r.stRx = reg

	case RegionRule:
		if jr.Same != "" {
			reg, err := regexp.Compile(jr.Same)
			if err != nil {
				return nil, fmt.Errorf("Rule %v %v %v", jr.Type, jr.Same, err)
			}
			r.stRx = reg
		} else {
			reg, err := regexp.Compile(jr.Start)
			if err != nil {
				return nil, fmt.Errorf("Rule %v %v %v", jr.Type, jr.Start, err)
			}
			r.stRx = reg
			reg, err = regexp.Compile(jr.End)
			if err != nil {
				return nil, fmt.Errorf("Rule %v %v %v", jr.Type, jr.End, err)
			}
			r.enRx = reg
		}
	}

	return &r, nil
}
