package action

// Actioner represents an editor action
type Actioner interface {
	Name() string
	Target() string
	SetTarget(string)
	Param() string
	SetParam(string)
	Column() int
	SetColumn(int)
	Line() int
	SetLine(int)
	Object() string
	SetObject(string)
	Count() int
	SetCount(int)
	Valid(bufid string) error
	NeedBuffer() bool
}

// Transactioner is group of actions for with a a buffer id
type Transactioner interface {
	Buffer() string
	SetBuffer(string)
	Actions() []Actioner
	Add(acts ...Actioner)
	Set(acts ...Actioner)
	Valid() error
	NeedBuffer() bool
}
