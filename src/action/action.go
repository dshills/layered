package action

// Action is an editor action
type Action struct {
	n, t, p string
	ln, col int
	obj     string
	cnt     int
}

// Name returns the action name
func (a *Action) Name() string { return a.n }

// Target is the target of the action
func (a *Action) Target() string { return a.t }

// Param is a required parameter
func (a *Action) Param() string { return a.p }

// Line will return the line -1 represents not set
func (a *Action) Line() int { return a.ln }

// Column returns the column -1 represents not set
func (a *Action) Column() int { return a.col }

// SetLine will set the line associated with the action
func (a *Action) SetLine(l int) { a.ln = l }

// SetColumn will set the column associated with the action
func (a *Action) SetColumn(c int) { a.col = c }

// Object returns the object associated with the action
func (a *Action) Object() string { return a.obj }

// SetObject will set the object for the action
func (a *Action) SetObject(obj string) { a.obj = obj }

// Count will return the action count
func (a *Action) Count() int { return a.cnt }

// SetCount will set the action count
func (a *Action) SetCount(c int) { a.cnt = c }

// New will return a new Actioner
func New(name, target, param string) Actioner {
	return &Action{n: name, t: target, p: param, ln: -1, col: -1, cnt: 1}
}

// Transaction is a group of actions with a buffer identifier
type Transaction struct {
	acts []Actioner
	id   string
}

// Buffer will return the buffer id
func (t *Transaction) Buffer() string { return t.id }

// SetBuffer will set the transaction buffer
func (t *Transaction) SetBuffer(id string) {
	t.id = id
}

// Actions returns the set of actions
func (t *Transaction) Actions() []Actioner { return t.acts }

// Add will add actions to the transactions
func (t *Transaction) Add(acts ...Actioner) {
	t.acts = append(t.acts, acts...)
}

// Set will set actions to the transactions
func (t *Transaction) Set(acts ...Actioner) {
	t.acts = acts
}

// NewTransaction will return an we transactioner
func NewTransaction(id string, actions ...Actioner) Transactioner {
	return &Transaction{id: id, acts: actions}
}
