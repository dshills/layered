package filetype

import "github.com/dshills/layered/conf"

// Factory creates a
type Factory func(*conf.Configuration) (Manager, error)

// Manager is a file type detecter
type Manager interface {
	Load() error
	Detect(path string) (string, error)
}
