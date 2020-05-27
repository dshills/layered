package syntax

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/dshills/layered/textstore"
)

// Matcher is syntax matcher
type Matcher struct {
	runtimes []string
	rules    []Ruler
}

// LoadFileType will load a syntax file by file type
func (m *Matcher) LoadFileType(ft string) error {
	ft = strings.ToLower(ft) + ".json"
	for i := len(m.runtimes) - 1; i >= 0; i-- {
		path := filepath.Join(m.runtimes[i], ft)
		if m.fileExists(path) {
			return m.LoadFile(path)
		}
	}
	return fmt.Errorf("Matcher.LoadFileType: Not found %v", ft)
}

func (m *Matcher) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// LoadFile will load a syntax file
func (m *Matcher) LoadFile(path string) error {
	m.rules = []Ruler{}
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	jrs := []jsRule{}
	if err = json.NewDecoder(f).Decode(&jrs); err != nil {
		return err
	}
	errs := []string{}
	rules := []Ruler{}
	for i := range jrs {
		r, err := jrs[i].asRuler()
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		rules = append(rules, r)
	}
	m.rules = rules
	if len(errs) > 0 {
		return fmt.Errorf("LoadFile: %v", strings.Join(errs, ", "))
	}
	return nil
}

// Add will add a rule to the matcher
func (m *Matcher) Add(r Ruler) {
	m.rules = append(m.rules, r)
}

// Parse will return a list of results for the text store
func (m *Matcher) Parse(ts textstore.TextStorer) []Resulter {
	errs := []string{}
	results := []Resulter{}
	cnt := ts.NumLines()
	for i := 0; i < cnt; i++ {
		for ii := range m.rules {
			reader, err := ts.LineAt(i)
			if err != nil {
				errs = append(errs, err.Error())
				continue
			}
			go func(rule Ruler, ln int) {
				results = append(results, rule.Match(reader, ln))
			}(m.rules[ii], i)
		}
	}
	sort.Sort(resultList(results))
	return results
}

// New returns a new syntax matcher
func New(rt ...string) Matcherer {
	return &Matcher{runtimes: rt}
}
