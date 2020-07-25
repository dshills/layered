# filetype
--
    import "."


## Usage

#### type FTDetecter

```go
type FTDetecter struct {
}
```

FTDetecter determines file types

#### func (*FTDetecter) AddRuntime

```go
func (fd *FTDetecter) AddRuntime(paths ...string)
```
AddRuntime will add a directory to the list of search directories

#### func (*FTDetecter) Detect

```go
func (fd *FTDetecter) Detect(path string) (string, error)
```
Detect will return a file type or ""

#### func (*FTDetecter) Load

```go
func (fd *FTDetecter) Load() error
```
Load will load file type detecters

#### func (*FTDetecter) RemoveRuntime

```go
func (fd *FTDetecter) RemoveRuntime(path string)
```
RemoveRuntime will remove a directory from the runetime list

#### type Factory

```go
type Factory func(runtime ...string) (Manager, error)
```

Factory creates a

#### type Manager

```go
type Manager interface {
	AddRuntime(rtpath ...string)
	RemoveRuntime(rtpath string)
	Load() error
	Detect(path string) (string, error)
}
```

Manager is a file type detecter

#### func  New

```go
func New(rtpaths ...string) (Manager, error)
```
New returns a new file type detecter
