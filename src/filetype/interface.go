package filetype

// Detecter is a file type detecter
type Detecter interface {
	Load(fts string) error
	Detect(path string) (string, error)
}
