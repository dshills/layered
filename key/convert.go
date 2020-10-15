package key

import (
	"fmt"
	"strings"
)

// StrToKey converts a string representaion to a rune, key
func StrToKey(s string) (r rune, k int, err error) {
	if len(s) == 0 {
		err = fmt.Errorf("StrToKey: Blank string")
		return
	}
	s = parseKeyString(s)
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

func parseKeyString(s string) string {
	may := false
	spe := ""
	norm := ""
	for _, c := range s {
		switch {
		case c == '<' && may:
			norm += spe
			spe = ""
			may = false
		case c == '<':
			norm += spe
			spe = ""
			may = true
		case c == '>' && may:
			norm += convertSpecial(spe)
			spe = ""
			may = false
		case may:
			spe += string(c)
		default:
			norm += string(c)
		}
	}
	norm += spe
	return norm
}

func convertSpecial(s string) string {
	if len(s) < 3 {
		return "<" + s + ">"
	}

	rs := strings.ToLower(string([]rune(s)[:2]))
	switch rs {
	case "c-":
		return "<ctrl-" + string([]rune(s)[3:]) + ">"
	case "a-":
		return "<alt-" + string([]rune(s)[3:]) + ">"
	}

	if s == "ins" {
		return "<insert>"
	}
	if s == "del" {
		return "<delete>"
	}

	// Check special list
	lc := "<" + strings.ToLower(s) + ">"
	for _, sp := range specialKeys {
		if lc == sp {
			return sp
		}
	}
	return "<" + s + ">"
}
