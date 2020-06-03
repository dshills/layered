# key
--
    import "."


## Usage

#### type Key

```go
type Key struct {
}
```

Key is a keyboard key press

#### func (*Key) Key

```go
func (k *Key) Key() int
```
Key will return the key value rune value == 0

#### func (*Key) Rune

```go
func (k *Key) Rune() rune
```
Rune returns the rune code, key == 0

#### func (*Key) String

```go
func (k *Key) String() string
```

#### type Keyer

```go
type Keyer interface {
	Rune() rune
	Key() int
}
```

Keyer represents a keyboard item

#### func  New

```go
func New(r rune, k int) Keyer
```
New will return a key
