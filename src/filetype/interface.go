package filetype

// Factory creates a
type Factory func(runtime ...string) (Detecter, error)

// Detecter is a file type detecter
type Detecter interface {
	Load(fts string) error
	Detect(path string) (string, error)
	AddDirectory(path ...string) error
	RemoveDirectory(path string)
}
