package action

// Actioner represents an editor action
type Actioner interface {
	Name() string
	Target() string
	Param() string
}

// Transactioner is group of actions for with a a buffer id
type Transactioner interface {
	Buffer() string
	Actions() []Actioner
}
