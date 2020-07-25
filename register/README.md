# register
--
    import "."


## Usage

#### type Factory

```go
type Factory func() Registerer
```

Factory will return a new registerer

#### type Marker

```go
type Marker struct {
	Line  int
	Match string
}
```

Marker is a document mark

#### type Register

```go
type Register struct {
}
```

Register is a register

#### func (*Register) AddDefYank

```go
func (r *Register) AddDefYank(s string)
```
AddDefYank will set the default yank

#### func (*Register) AddSearch

```go
func (r *Register) AddSearch(s string)
```
AddSearch will set the current search term

#### func (*Register) AddYank

```go
func (r *Register) AddYank(key, s string)
```
AddYank will add a yank to a register

#### func (*Register) AltFile

```go
func (r *Register) AltFile() string
```
AltFile will return the alternative file

#### func (*Register) Colon

```go
func (r *Register) Colon() string
```
Colon returns the last colon command

#### func (*Register) CurrentFile

```go
func (r *Register) CurrentFile() string
```
CurrentFile will return the current file

#### func (*Register) DefYank

```go
func (r *Register) DefYank() string
```
DefYank will return the default yank

#### func (*Register) Inserted

```go
func (r *Register) Inserted() string
```
Inserted is the last inserted text

#### func (*Register) Mark

```go
func (r *Register) Mark(key string) Marker
```
Mark wil;l return the mark associated with key

#### func (*Register) Reg

```go
func (r *Register) Reg(key string) string
```
Reg will return the register by key

#### func (*Register) Search

```go
func (r *Register) Search() string
```
Search will return the last search

#### func (*Register) SetAltFile

```go
func (r *Register) SetAltFile(s string)
```
SetAltFile will set the alternative file

#### func (*Register) SetColon

```go
func (r *Register) SetColon(s string)
```
SetColon will set the last colon command

#### func (*Register) SetCurrentFile

```go
func (r *Register) SetCurrentFile(s string)
```
SetCurrentFile will set the current file

#### func (*Register) SetInserted

```go
func (r *Register) SetInserted(s string)
```
SetInserted sets the last inserted text

#### func (*Register) SetMark

```go
func (r *Register) SetMark(key string, line int, match string)
```
SetMark will set a mark

#### func (*Register) SetReg

```go
func (r *Register) SetReg(key, s string)
```
SetReg will set a register

#### func (*Register) Yank

```go
func (r *Register) Yank(key string) string
```
Yank will return the yank register

#### type Registerer

```go
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
```

Registerer is a general interface for a register

#### func  New

```go
func New() Registerer
```
New is a register factory
