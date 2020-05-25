package key

import (
	"strings"
)

// Key is a keyboard key press
type Key struct {
	s, a, c bool
	r       rune
	sp      string
}

// Special returns true if a special key press
func (k *Key) Special() bool { return k.s }

// SpecialKey returns true if a special key press
func (k *Key) SpecialKey() string { return k.sp }

// Alt returns true if an alt key press
func (k *Key) Alt() bool { return k.a }

// Ctrl returns true if an ctrl key press
func (k *Key) Ctrl() bool { return k.c }

// Rune returns the key rune
func (k *Key) Rune() rune { return k.r }

// IsEqual returns true if same key
func (k *Key) IsEqual(o Keyer) bool {
	if k.a != o.Alt() || k.c != o.Ctrl() || k.s != o.Special() {
		return false
	}
	if k.s && k.sp != o.SpecialKey() {
		return false
	}
	if k.r != o.Rune() {
		return false
	}
	return true
}

// NewKey will return a key
func NewKey(a, c, s bool, r rune, sp string) Keyer {
	return &Key{a: a, c: c, s: s, r: r, sp: sp}
}

// StrToKey converts a string representaion to a Keyer
func StrToKey(s string) Keyer {
	if len([]rune(s)) == 1 {
		return NewKey(false, false, false, []rune(s)[0], "")
	}
	ls := strings.ToLower(s)
	switch {
	case strings.HasPrefix(ls, "<ctrl-") && strings.HasSuffix(ls, ">"):
		r := []rune(s)[6]
		return NewKey(false, true, false, r, "")
	case strings.HasPrefix(ls, "<c-") && strings.HasSuffix(ls, ">"):
		r := []rune(s)[3]
		return NewKey(false, true, false, r, "")
	case strings.HasPrefix(ls, "<alt-") && strings.HasSuffix(ls, ">"):
		r := []rune(s)[5]
		return NewKey(true, false, false, r, "")
	case strings.HasPrefix(ls, "<a-") && strings.HasSuffix(ls, ">"):
		r := []rune(s)[3]
		return NewKey(true, false, false, r, "")
	}
	return NewKey(false, false, true, 0, ls)
}
