package register

import "sync"

// Register is a register
type Register struct {
	marks    map[string]Marker
	mm       sync.RWMutex
	search   string
	colon    string
	altfile  string
	file     string
	inserted string
	yank     map[string]string
	ym       sync.RWMutex
	reg      map[string]string
	rm       sync.RWMutex
	defYank  string
}

// Mark wil;l return the mark associated with key
func (r *Register) Mark(key string) Marker {
	r.mm.RLock()
	defer r.mm.RUnlock()
	return r.marks[key]
}

// SetMark will set a mark
func (r *Register) SetMark(key string, line int, match string) {
	r.mm.Lock()
	defer r.mm.Unlock()
	r.marks[key] = Marker{Line: line, Match: match}
}

// Search will return the last search
func (r *Register) Search() string { return r.search }

// AddSearch will set the current search term
func (r *Register) AddSearch(s string) { r.search = s }

// Inserted is the last inserted text
func (r *Register) Inserted() string { return r.inserted }

// SetInserted sets the last inserted text
func (r *Register) SetInserted(s string) { r.inserted = s }

// Colon returns the last colon command
func (r *Register) Colon() string { return r.colon }

// SetColon will set the last colon command
func (r *Register) SetColon(s string) { r.colon = s }

// CurrentFile will return the current file
func (r *Register) CurrentFile() string { return r.file }

// SetCurrentFile will set the current file
func (r *Register) SetCurrentFile(s string) { r.file = s }

// AltFile will return the alternative file
func (r *Register) AltFile() string { return r.altfile }

// SetAltFile will set the alternative file
func (r *Register) SetAltFile(s string) { r.altfile = s }

// Yank will return the yank register
func (r *Register) Yank(key string) string {
	r.ym.RLock()
	defer r.ym.RUnlock()
	return r.yank[key]
}

// DefYank will return the default yank
func (r *Register) DefYank() string { return r.defYank }

// AddYank will add a yank to a register
func (r *Register) AddYank(key, s string) {
	r.ym.Lock()
	defer r.ym.Unlock()
	r.yank[key] = s
}

// AddDefYank will set the default yank
func (r *Register) AddDefYank(s string) { r.defYank = s }

// Reg will return the register by key
func (r *Register) Reg(key string) string {
	r.rm.RLock()
	defer r.rm.RUnlock()
	return r.reg[key]
}

// SetReg will set a register
func (r *Register) SetReg(key, s string) {
	r.rm.Lock()
	defer r.rm.Unlock()
	r.reg[key] = s
}

// New is a register factory
func New() Registerer {
	return &Register{marks: make(map[string]Marker), yank: make(map[string]string), reg: make(map[string]string)}
}
