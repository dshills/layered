package filetype

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// FTDetecter determines file types
type FTDetecter struct {
	exts     map[string]string
	m        sync.RWMutex
	patterns []ftEntry
	runtimes []string
}

type ftEntry struct {
	Pattern string `json:"pattern"`
	FT      string `json:"ft"`
	Ext     string `json:"ext"`
	regEx   *regexp.Regexp
}

// Load will load the ft detections
func (fd *FTDetecter) Load(path string) error {
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

func (fd *FTDetecter) loadAll() error {
	errs := []string{}
	for i := len(fd.runtimes) - 1; i >= 0; i-- {
		path := filepath.Join(fd.runtimes[i], "config", "ftdetect.json")
		if err := fd.Load(path); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("FTDetecter.Load: %v", strings.Join(errs, ", "))
	}
	return nil
}

// AddDirectory will add a directory to the list of search directories
func (fd *FTDetecter) AddDirectory(paths ...string) error {
	fd.runtimes = append(fd.runtimes, paths...)
	return fd.loadAll()
}

// RemoveDirectory will remove a directory from the runetime list
func (fd *FTDetecter) RemoveDirectory(path string) {
	dl := []string{}
	for i := range fd.runtimes {
		if fd.runtimes[i] != path {
			dl = append(dl, fd.runtimes[i])
		}
	}
	fd.runtimes = dl
}

// New returns a new file type detecter
func New(rtpaths ...string) (Detecter, error) {
	ft := &FTDetecter{exts: make(map[string]string)}
	err := ft.AddDirectory(rtpaths...)
	return ft, err
}
