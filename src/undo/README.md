# undo
--
    import "."


## Usage

#### type Change

```go
type Change struct {
}
```

Change is a change to a text store

#### func (*Change) Cursor

```go
func (c *Change) Cursor() []int
```
Cursor returns the cursor position before the change

#### func (*Change) Dirty

```go
func (c *Change) Dirty() bool
```
Dirty returns the dirty status before the change

#### func (*Change) Line

```go
func (c *Change) Line() int
```
Line will return the line the change was made

#### func (*Change) Type

```go
func (c *Change) Type() ChangeType
```
Type will return the change type

#### func (*Change) Undo

```go
func (c *Change) Undo(after string) string
```
Undo will return the text before the change

#### type ChangeSet

```go
type ChangeSet struct {
}
```

ChangeSet is a set of editor changes

#### func (*ChangeSet) AddChanges

```go
func (cs *ChangeSet) AddChanges(c ...Changer)
```
AddChanges will add changes to the change set

#### func (*ChangeSet) AddLine

```go
func (cs *ChangeSet) AddLine(ln int)
```
AddLine will add a line add to the set

#### func (*ChangeSet) ChangeLine

```go
func (cs *ChangeSet) ChangeLine(ln int, before, after string)
```
ChangeLine will add a line change to the set

#### func (*ChangeSet) Changes

```go
func (cs *ChangeSet) Changes() []Changer
```
Changes returns the list of changes

#### func (*ChangeSet) RemoveLine

```go
func (cs *ChangeSet) RemoveLine(ln int)
```
RemoveLine will add a line deletion to the set

#### type ChangeSetter

```go
type ChangeSetter interface {
	AddChanges(...Changer)
	Changes() []Changer
	RemoveLine(ln int)
	AddLine(ln int)
	ChangeLine(ln int, before, after string)
}
```

ChangeSetter is a set of changes

#### func  New

```go
func New() ChangeSetter
```
New will return a new ChangeSet

#### type ChangeType

```go
type ChangeType int
```

ChangeType is a the type of change made

```go
const (
	DeleteLine ChangeType = iota
	AddLine
	ChangeLine
)
```
Change types

#### func (ChangeType) String

```go
func (c ChangeType) String() string
```

#### type Changer

```go
type Changer interface {
	Cursor() []int
	Dirty() bool
	Type() ChangeType
	Line() int
	Undo(after string) string
}
```

Changer is a change to a text store

#### type Factory

```go
type Factory func() ChangeSetter
```

Factory is a function that returns a new change set
