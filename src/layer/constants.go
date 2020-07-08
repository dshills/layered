package layer

// MatchStatus is the status of a parser operation
type MatchStatus int

// ParseStatus constants
const (
	NoMatch MatchStatus = iota
	PartialMatch
	Match
)

func (s MatchStatus) String() string {
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
