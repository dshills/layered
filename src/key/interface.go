package key

// Keyer represents a keyboard item
type Keyer interface {
	Special() bool
	Alt() bool
	Ctrl() bool
	Rune() rune
	SpecialKey() string
	IsEqual(o Keyer) bool
}
