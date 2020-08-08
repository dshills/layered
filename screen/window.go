package screen

type window struct {
	id            string
	hidden        bool
	x, y          int
	width, height int
}

func (w *window) SetHidden(h bool) { w.hidden = h }
func (w *window) Hidden() bool     { return w.hidden }
func (w *window) SetLocation(x, y, width, height int) {
	w.x = x
	w.width = width
	w.y = y
	w.height = height
}
func (w *window) Location() (int, int, int, int) {
	return w.x, w.y, w.width, w.height
}
func (w *window) SplitVertical(id string)   {}
func (w *window) SplitHorizontal(id string) {}
func (w *window) Move(x, y int) {
	w.x = x
	w.y = y
}
func (w *window) ID() string      { return w.id }
func (w *window) SetID(id string) { w.id = id }

// NewWindow returns a new Window
func NewWindow(id string) Window {
	return &window{id: id}
}
