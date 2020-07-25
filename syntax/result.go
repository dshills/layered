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

// SetToken will set the result token
func (r *Result) SetToken(tok string) { r.tok = tok }

// SetLine will set the result line
func (r *Result) SetLine(ln int) { r.ln = ln }

// SetRanges will set the result ranges
func (r *Result) SetRanges(rg [][]int) { r.rg = rg }

// AddRanges will append result ranges
func (r *Result) AddRanges(rg [][]int) {
	r.rg = append(r.rg, rg...)
}

// SetPriority will set the priority of the result
func (r *Result) SetPriority(p int) { r.pr = p }

type resultList []Resulter

func (rl resultList) Len() int           { return len(rl) }
func (rl resultList) Swap(i, j int)      { rl[i], rl[j] = rl[j], rl[i] }
func (rl resultList) Less(i, j int) bool { return rl[i].Priority() > rl[j].Priority() }
