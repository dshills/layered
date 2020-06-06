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
