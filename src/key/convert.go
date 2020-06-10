package key

import "fmt"

// StrToKey converts a string representaion to a rune, key
func StrToKey(s string) (r rune, k int, err error) {
	var ok bool
	k, ok = convertKeyTable[s]
	if ok {
		return
	}
	r, ok = convertCharTable[s]
	if ok {
		return
	}
	if len(s) > 1 {
		return 0, 0, fmt.Errorf("StrToKey: Unknown key %v", s)
	}
	r = []rune(s)[0]
	return
}

// StrToKeyer will return a Keyer based on s
func StrToKeyer(s string) (Keyer, error) {
	r, k, err := StrToKey(s)
	if err != nil {
		return nil, err
	}
	return New(r, k), nil
}

// SpecialToString will convert a key.k value to a string
func SpecialToString(k int) string {
	switch k {
	case KeyInsert:
		return "<ins>"
	case KeyDelete:
		return "<del>"
	case KeyHome:
		return "<home>"
	case KeyEnd:
		return "<end>"
	case KeyPgup:
		return "<pgup>"
	case KeyPgdn:
		return "<pgdown>"
	case KeyArrowUp:
		return "<up>"
	case KeyArrowDown:
		return "<down>"
	case KeyArrowLeft:
		return "<left>"
	case KeyArrowRight:
		return "<right>"
	case KeyEnter:
		return "<cr>"
	case KeyEsc:
		return "<esc>"
	case KeySpace:
		return "<space>"
	case KeyTab:
		return "<tab>"
	case KeyBackspace:
		return "<bs>"
	}
	for key, val := range convertKeyTable {
		if val == k {
			return key
		}
	}
	return "<unknown>"
}
