package key

// Keyer represents a keyboard item
type Keyer interface {
	Rune() rune
	Key() int
	Eq(Keyer) bool
}
