package action

// Action is an editor action
type Action struct {
	n, t, p string
}

// Name returns the action name
func (a *Action) Name() string { return a.n }

// Target is the target of the action
func (a *Action) Target() string { return a.t }

// Param is a required parameter
func (a *Action) Param() string { return a.p }

// New will return a new Actioner
func New(name, target, param string) Actioner {
	return &Action{n: name, t: target, p: param}
}

// Transaction is a group of actions with a buffer identifier
type Transaction struct {
	acts []Actioner
	buf  string
}

// Buffer will return the buffer id
func (t *Transaction) Buffer() string { return t.buf }

// Actions returns the set of actions
func (t *Transaction) Actions() []Actioner { return t.acts }
