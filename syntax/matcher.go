package syntax

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/dshills/layered/conf"
	"github.com/dshills/layered/textstore"
)

// Matcher is syntax matcher
type Matcher struct {
	rules  []Ruler
	config *conf.Configuration
}

// Parse will return a list of results for the text store
// optionally a list of rule groups to use
func (m *Matcher) Parse(ts textstore.TextStorer, groups ...string) []Resulter {
	results := []Resulter{}
	wg := sync.WaitGroup{}
	for _, rule := range m.filterRules(groups) {
		wg.Add(1)
		go func(rule Ruler, wg *sync.WaitGroup) {
			results = append(results, rule.Match(ts)...)
			wg.Done()
		}(rule, &wg)
	}
	wg.Wait()
	results = m.dependencies(results)
	sort.Sort(resultList(results))
	return results
}

// FilterResults will filter results by group
func (m *Matcher) FilterResults(results []Resulter, groups ...string) []Resulter {
	nr := []Resulter{}
	for _, res := range results {
		for _, grp := range groups {
			if strings.Contains(res.Token(), grp) {
				nr = append(nr, res)
			}
		}
	}
	return nr
}

func (m *Matcher) filterRules(groups []string) []Ruler {
	if len(groups) == 0 {
		return m.rules
	}
	rules := []Ruler{}
	for _, r := range m.rules {
		for _, g := range groups {
			if strings.Contains(r.Group(), g) {
				rules = append(rules, r)
			}
		}
	}
	return rules
}

func (m *Matcher) dependencies(results []Resulter) []Resulter {
	// TODO fix contained results
	return results
}

// LoadFileType will load a syntax file by file type
func (m *Matcher) LoadFileType(ft string) error {
	ft = strings.ToLower(ft) + ".json"
	for _, rt := range m.config.Syntax() {
		path := filepath.Join(rt, ft)
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
		r, err := jrs[i].asRuler(i)
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

// New returns a new syntax matcher
func New(config *conf.Configuration) Manager {
	return &Matcher{config: config}
}
