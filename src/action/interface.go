package action

// Actioner represents an editor action
type Actioner interface {
	Name() string
	Target() string
	Param() string
	Line() int
	Column() int
	SetLine(int)
	SetColumn(int)
	Object() string
	SetObject(string)
	Count() int
	SetCount(int)
}

// Transactioner is group of actions for with a a buffer id
type Transactioner interface {
	Buffer() string
	Actions() []Actioner
}
