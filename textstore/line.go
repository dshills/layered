package textstore

import "github.com/dshills/layered/undo"

type line struct {
	text     []rune
	changes  []undo.Changer
	original []rune
}

func (l *line) reset(s string) {
	l.startEdit()
	defer l.endEdit()
	l.text = []rune(s)
}

func (l *line) insertRune(col int, r rune) {
	l.startEdit()
	defer l.endEdit()
	if col < 0 {
		l.text = append([]rune{r}, l.text...)
		return
	}
	if col >= len(l.text) {
		l.text = append(l.text, r)
		return
	}
	l.text = append(l.text, 0)
	copy(l.text[col+1:], l.text[col:])
	l.text[col] = r
}

func (l *line) replaceRune(col int, r rune) {
	l.startEdit()
	defer l.endEdit()
	if col < 0 {
		l.text = append([]rune{r}, l.text...)
		return
	}
	if col >= len(l.text) {
		l.text = append(l.text, r)
		return
	}
	l.text[col] = r
}

func (l *line) delete(col, cnt int) {
	l.startEdit()
	defer l.endEdit()
	if col < 0 {
		col = len(l.text) - cnt
		if col < 0 {
			col = 0
		}
	}
	if col >= len(l.text) {
		return
	}
	if col+cnt >= len(l.text) {
		l.text = l.text[col:]
		return
	}
}

func (l *line) insertString(col int, s string) {
	l.startEdit()
	defer l.endEdit()
	str := []rune(s)
	if col < 0 {
		l.text = append(str, l.text...)
		return
	}
	if col >= len(l.text) {
		l.text = append(l.text, str...)
		return
	}
	tmp := append(l.text[:col], str...)
	tmp = append(tmp, l.text[col:]...)
	l.text = tmp
}

func (l *line) replaceString(from, to int, s string) {
	l.startEdit()
	defer l.endEdit()
	str := []rune(s)
	if from < 0 {
		l.text = append(str, l.text...)
		return
	}
	if from >= len(l.text) {
		l.text = append(l.text, str...)
		return
	}
	tmp := append(l.text[:from], str...)
	if to < len(l.text) {
		tmp = append(tmp, l.text[to:]...)
	}
	l.text = tmp
}

func (l *line) startEdit() {
	l.original = l.text
}

func (l *line) endEdit() {
	l.changes = append(l.changes, undo.NewChange(string(l.original), string(l.text)))
}

func (l *line) changeSet(ln int, fac undo.Factory) undo.ChangeSetter {
	cs := fac()
	for _, c := range l.changes {
		c.SetLine(ln)
		cs.AddChanges(c)
	}
	l.changes = []undo.Changer{}
	return cs
}

func (l *line) Len() int {
	return len(string(l.text))
}

func (l *line) String() string {
	return string(l.text)
}

func (l *line) runeAt(col int) rune {
	return l.text[col]
}

func (l *line) delimited(del string) string {
	return string(l.text) + del
}

func newLine(s string) *line {
	return &line{text: []rune(s)}
}
