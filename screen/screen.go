package screen

import "fmt"

type screen struct {
	height, width int
	status        Statusbar
	notice        Noticebar
	windows       []Window
	active        Window
}

func (s *screen) SetSize(width, height int) {
	s.height = height
	s.width = width
}
func (s *screen) Size() (int, int)     { return s.width, s.height }
func (s *screen) Statusbar() Statusbar { return s.status }
func (s *screen) Noticebar() Noticebar { return s.notice }
func (s *screen) NewWindow(id string) {
	s.windows = append(s.windows, NewWindow(id))
}
func (s *screen) Window(id string) (Window, error) {
	for _, w := range s.windows {
		if w.ID() == id {
			return w, nil
		}
	}
	return nil, fmt.Errorf("Not found")
}
func (s *screen) Windows() []Window { return s.windows }
func (s *screen) SetActive(id string) {
	s.active, _ = s.Window(id)
}
func (s *screen) Active() Window { return s.active }
func (s *screen) ActivateRight() {}
func (s *screen) ActivateLeft()  {}
func (s *screen) ActivateUp()    {}
func (s *screen) ActivateDown()  {}

// New will return a new screen
func New() Screen {
	return &screen{
		status: NewStatusbar(),
		notice: NewNoticebar(),
	}
}
