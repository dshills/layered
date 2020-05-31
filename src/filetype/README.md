# filetype
--
    import "."


## Usage

#### type Detecter

```go
type Detecter interface {
	Load(fts string) error
	Detect(path string) (string, error)
	AddDirectory(path ...string) error
	RemoveDirectory(path string)
}
```

Detecter is a file type detecter

#### func  New

```go
func New(rtpaths ...string) (Detecter, error)
```
New returns a new file type detecter

#### type FTDetecter

```go
type FTDetecter struct {
}
```

FTDetecter determines file types

#### func (*FTDetecter) AddDirectory

```go
func (fd *FTDetecter) AddDirectory(paths ...string) error
```
AddDirectory will add a directory to the list of search directories

#### func (*FTDetecter) Detect

```go
func (fd *FTDetecter) Detect(path string) (string, error)
```
Detect will return a file type or ""

#### func (*FTDetecter) Load

```go
func (fd *FTDetecter) Load(path string) error
```
Load will load the ft detections

#### func (*FTDetecter) RemoveDirectory

```go
func (fd *FTDetecter) RemoveDirectory(path string)
```
RemoveDirectory will remove a directory from the runetime list

#### type Factory

```go
type Factory func(runtime ...string) (Detecter, error)
```

Factory creates a
