package register

// Factory will return a new registerer
type Factory func() Registerer

// Marker is a document mark
type Marker struct {
	Line  int
	Match string
}

// Registerer is a general interface for a register
type Registerer interface {
	Mark(key string) Marker
	SetMark(key string, line int, match string)
	Search() string
	AddSearch(string)
	Inserted() string
	SetInserted(string)
	Colon() string
	SetColon(string)
	CurrentFile() string
	SetCurrentFile(string)
	AltFile() string
	SetAltFile(string)
	Yank(key string) string
	DefYank() string
	AddYank(key, s string)
	AddDefYank(s string)
	Reg(string) string
	SetReg(key, s string)
}
