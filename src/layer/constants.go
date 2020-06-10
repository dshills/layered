package layer

// ParseStatus is the status of a parser operation
type ParseStatus int

// ParseStatus constants
const (
	NoMatch ParseStatus = iota
	PartialMatch
	Match
)

func (s ParseStatus) String() string {
	switch s {
	case NoMatch:
		return "No match"
	case PartialMatch:
		return "Partial match"
	case Match:
		return "Match"
	}
	return "Unknown status"
}

// Key matcher constants
const (
	Any       = "<any>"       // any key
	Printable = "<printable>" // Any printable character
	Control   = "<control>"   // any control character
	Digit     = "<digit>"     // 0-9
	Letter    = "<letter>"    // Any letter
	Lower     = "<lower>"     // Any lower case
	Upper     = "<upper>"     // Any upper case
	NonBlank  = "<non-blank>" // Any non space printable character
	Pattern   = "<pattern=>"  // regex pattern
)
