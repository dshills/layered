package filetype

// Factory creates a
type Factory func(runtime ...string) (Manager, error)

// Manager is a file type detecter
type Manager interface {
	AddRuntime(rtpath ...string)
	RemoveRuntime(rtpath string)
	Load() error
	Detect(path string) (string, error)
}
