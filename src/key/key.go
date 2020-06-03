package key

import "fmt"

// Key is a keyboard key press
type Key struct {
	r rune // char code, k == 0
	k int  // key code, r == 0
}

// Key will return the key value rune value == 0
func (k *Key) Key() int { return k.k }

// Rune returns the rune code, key == 0
func (k *Key) Rune() rune { return k.r }

func (k *Key) String() string {
	return fmt.Sprintf("Rune: %v (%v), %v", k.r, string(k.r), k.k)
}

// New will return a key
func New(r rune, k int) Keyer {
	return &Key{r: r, k: k}
}
