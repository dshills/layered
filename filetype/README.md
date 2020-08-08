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

#### type Factory

```go
type Factory func(*conf.Configuration) (Manager, error)
```

Factory creates a

#### type Manager

```go
type Manager interface {
	Load() error
	Detect(path string) (string, error)
}
```

Manager is a file type detecter

#### func  New

```go
func New(config *conf.Configuration) (Manager, error)
```
New returns a new file type detecter
