package syntax

// Result is a a syntax match result
type Result struct {
	tok string
	ln  int
	rg  [][]int
	pr  int
}

// IsEqual compares one result to another returning true if equal
func (r *Result) IsEqual(Resulter) bool {
	return false
}

// Token will return the rules token type
func (r *Result) Token() string { return r.tok }

// Line returns the line for the result
func (r *Result) Line() int { return r.ln }

// Range returns the range if matches
func (r *Result) Range() [][]int { return r.rg }

// Priority returns the results priority
func (r *Result) Priority() int { return r.pr }

type resultList []Resulter

func (rl resultList) Len() int           { return len(rl) }
func (rl resultList) Swap(i, j int)      { rl[i], rl[j] = rl[j], rl[i] }
func (rl resultList) Less(i, j int) bool { return rl[i].Priority() < rl[j].Priority() }
