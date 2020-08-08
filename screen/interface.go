package screen

// BarPos is the position for a bar
type BarPos int

// Bar positions
const (
	Top BarPos = iota
	Bottom
)

// Screen represents the application screen area
type Screen interface {
	SetSize(height, width int)
	Size() (int, int)
	Statusbar() Statusbar
	Noticebar() Noticebar
	NewWindow(id string)
	Window(id string) (Window, error)
	Windows() []Window
	SetActive(id string)
	Active() Window
	ActivateRight()
	ActivateLeft()
	ActivateUp()
	ActivateDown()
}

// Window is a portion of the screen
// usually displaying a text buffer
type Window interface {
	ID() string
	SetID(id string)
	SetHidden(bool)
	Hidden() bool
	SetLocation(int, int, int, int)
	Location() (int, int, int, int)
	SplitVertical(id string)
	SplitHorizontal(id string)
	Move(x, y int)
}

// Noticebar provides user feedback
type Noticebar interface {
	SetHidden(bool)
	Hidden() bool
	SetPosition(BarPos)
	Position() BarPos
	SetHeight(int)
	Height() int
	Clear()
	Set(notice string)
	Output(width int) string
}

// Statusbar is the status bar
type Statusbar interface {
	SetHidden(bool)
	Hidden() bool
	SetPosition(BarPos)
	Position() BarPos
	Components() []SBComponent
	SetComponentValue(key string, val string)
	ComponentValue(key string) string
	Add(...SBComponent)
	Remove(key string)
	Output(width int) string
}

// SBComponent is a statusbar item
type SBComponent struct {
	Key      string
	Right    bool
	Pre      string
	Post     string
	maxWidth int
}
