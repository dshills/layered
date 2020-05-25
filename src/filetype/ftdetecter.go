package filetype

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/dshills/layered/logger"
)

// FTDetecter determines file types
type FTDetecter struct {
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

	failCnt := 0
	for _, en := range pats {
		if en.Ext != "" {
			splits := strings.Split(en.Ext, ",")
			for _, ext := range splits {
				fd.exts[strings.TrimSpace(ext)] = en.FT
			}
		} else {
			en.regEx, err = regexp.Compile(en.Pattern)
			if err != nil {
				failCnt++
				//logger.Errorf("FTDetecter %v %v", en.Pattern, err)
			}
		}
	}
	if failCnt > 0 {
		logger.Errorf("FTDetecter %v patterns failed to compile", failCnt)
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

// NewFTDetecter returns a new file type detecter
func NewFTDetecter() Detecter {
	return &FTDetecter{exts: make(map[string]string)}
}
