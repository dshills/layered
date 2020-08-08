package filetype

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/dshills/layered/conf"
)

// FTDetecter determines file types
type FTDetecter struct {
	config   *conf.Configuration
	exts     map[string]string
	m        sync.RWMutex
	patterns []ftEntry
}

type ftEntry struct {
	Pattern string `json:"pattern"`
	FT      string `json:"ft"`
	Ext     string `json:"ext"`
	regEx   *regexp.Regexp
}

// Load will load file type detecters
func (fd *FTDetecter) Load() error {
	errs := []string{}
	for _, p := range fd.config.FTDetect() {
		if err := fd._load(p); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("%v", strings.Join(errs, ", "))
	}
	return nil
}

// Load will load the ft detections
func (fd *FTDetecter) _load(path string) error {
	//logger.Debugf("Load file type detection %v", path)
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	pats := []ftEntry{}
	err = json.NewDecoder(f).Decode(&pats)
	if err != nil {
		return err
	}
	fd.m.Lock()
	defer fd.m.Unlock()

	errs := []string{}
	for _, en := range pats {
		if en.Ext != "" {
			splits := strings.Split(en.Ext, ",")
			for _, ext := range splits {
				fd.exts[strings.TrimSpace(ext)] = en.FT
			}
		} else {
			en.regEx, err = regexp.Compile(en.Pattern)
			if err != nil {
				err = fmt.Errorf("%v", en.FT)
				errs = append(errs, err.Error())
			}
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ","))
	}
	return nil
}

// Detect will return a file type or ""
func (fd *FTDetecter) Detect(path string) (string, error) {
	fd.m.RLock()
	defer fd.m.RUnlock()
	ext := filepath.Ext(path)
	if len(ext) > 0 && ext[0] == '.' {
		ext = ext[1:]
	}
	ft, ok := fd.exts[ext]
	if ok {
		return ft, nil
	}
	for _, pat := range fd.patterns {
		if pat.regEx != nil {
			if pat.regEx.MatchString(path) {
				return pat.FT, nil
			}
		}
	}
	return "", fmt.Errorf("Not found")
}

// New returns a new file type detecter
func New(config *conf.Configuration) (Manager, error) {
	ft := &FTDetecter{exts: make(map[string]string), config: config}
	err := ft.Load()
	return ft, err
}
