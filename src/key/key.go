package key

import (
	"fmt"
	"strings"
	"unicode"
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

// Matches returns the number of key matches from 0
func (k *Key) Matches(keys ...Keyer) int {
	kk := []Keyer{}
	for i := range keys {
		kk = append(kk, keys[i])
		if !k.IsEqual(kk...) {
			return len(kk) - 1
		}
	}
	return len(kk)
}

// IsEqual returns true if key(s) match
func (k *Key) IsEqual(keys ...Keyer) bool {
	switch k.sp {
	case Any:
		return true
	case Printable:
		for i := range keys {
			if keys[i].Special() || keys[i].Alt() || keys[i].Ctrl() {
				return false
			}
			if !unicode.IsGraphic(keys[i].Rune()) {
				return false
			}
		}
		return true
	case Control:
		for i := range keys {
			if !unicode.IsControl(keys[i].Rune()) {
				return false
			}
		}
		return true
	case Digit:
		for i := range keys {
			if keys[i].Special() || keys[i].Alt() || keys[i].Ctrl() {
				return false
			}
			if !unicode.IsDigit(keys[i].Rune()) {
				return false
			}
		}
		return true
	case Letter:
		for i := range keys {
			if keys[i].Special() || keys[i].Alt() || keys[i].Ctrl() {
				return false
			}
			if !unicode.IsLetter(keys[i].Rune()) {
				return false
			}
		}
		return true
	case Lower:
		for i := range keys {
			if keys[i].Special() || keys[i].Alt() || keys[i].Ctrl() {
				return false
			}
			if !unicode.IsLower(keys[i].Rune()) {
				return false
			}
		}
		return true
	case Upper:
		for i := range keys {
			if keys[i].Special() || keys[i].Alt() || keys[i].Ctrl() {
				return false
			}
			if !unicode.IsUpper(keys[i].Rune()) {
				return false
			}
		}
		return true
	case NonBlank:
		for i := range keys {
			if keys[i].Special() || keys[i].Alt() || keys[i].Ctrl() {
				return false
			}
			if !unicode.IsGraphic(keys[i].Rune()) || unicode.IsSpace(keys[i].Rune()) {
				return false
			}
		}
		return true
	case Pattern:
		// TODO
		return false
	}
	if len(keys) > 1 {
		return false
	}

	if k.a != keys[0].Alt() || k.c != keys[0].Ctrl() || k.s != keys[0].Special() {
		return false
	}
	if k.s && k.sp != keys[0].SpecialKey() {
		return false
	}
	if k.r != keys[0].Rune() {
		return false
	}
	return true
}

// IsMatchMultiple returns true if the key pattern matches multiple keys
func (k *Key) IsMatchMultiple() bool {
	switch k.sp {
	case Any:
		return true
	case Printable:
		return true
	case Control:
		return true
	case Digit:
		return true
	case Letter:
		return true
	case Lower:
		return true
	case Upper:
		return true
	case NonBlank:
		return true
	case Pattern:
		return true
	}
	return false
}

// NewKey will return a key
func NewKey(a, c, s bool, r rune, sp string) Keyer {
	return &Key{a: a, c: c, s: s, r: r, sp: sp}
}

func (k *Key) String() string {
	if k.c {
		return fmt.Sprintf("<ctrl-%v>", string(k.r))
	}
	if k.a {
		return fmt.Sprintf("<alt-%v>", string(k.r))
	}
	if k.s {
		return k.sp
	}
	return string(k.r)
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
