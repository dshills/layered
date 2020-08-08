package screen

type noticebar struct {
	hidden bool
	pos    BarPos
	height int
	notice string
}

func (nb *noticebar) SetHidden(h bool)      { nb.hidden = h }
func (nb *noticebar) Hidden() bool          { return nb.hidden }
func (nb *noticebar) SetPosition(bp BarPos) { nb.pos = bp }
func (nb *noticebar) Position() BarPos      { return nb.pos }
func (nb *noticebar) SetHeight(h int)       { nb.height = h }
func (nb *noticebar) Height() int           { return nb.height }
func (nb *noticebar) Clear()                { nb.notice = "" }
func (nb *noticebar) Set(notice string)     { nb.notice = notice }
func (nb *noticebar) Output(width int) string {
	width *= nb.height
	if len(nb.notice) >= width {
		return string([]rune(nb.notice)[:width])
	}
	return nb.notice
}

// NewNoticebar will return a noticebar
func NewNoticebar() Noticebar {
	return &noticebar{pos: Bottom, height: 1}
}
