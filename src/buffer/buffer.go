package buffer

import (
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/syntax"
	"github.com/dshills/layered/textstore"
	"github.com/google/uuid"
)

// Buffer is a text buffer
type Buffer struct {
	id            string
	filename      string
	filetype      string
	cur           cursor.Cursorer
	txt           textstore.TextStorer
	mat           syntax.Matcherer
	ftd           filetype.Detecter
	dirty         bool
	updates       chan bool
	syntaxResults []syntax.Resulter
}

// ID will return the identifier for the buffer
func (b *Buffer) ID() string { return b.id }

// Filename will return the buffer filename
func (b *Buffer) Filename() string { return b.filename }

// SetFilename will set the buffers file name
func (b *Buffer) SetFilename(n string) {
	b.filename = n
	ft, err := b.ftd.Detect(n)
	if err != nil {
		return
	}
	b.SetFiletype(ft)
}

// Filetype returns the buffer's file type
func (b *Buffer) Filetype() string { return b.filetype }

// SetFiletype will set the buffer's file type
func (b *Buffer) SetFiletype(ft string) {
	b.filetype = ft
	b.mat.LoadFileType(ft)
	b.matchSyntax()
}

// TextStore will return the buffer's text store
func (b *Buffer) TextStore() textstore.TextStorer { return b.txt }

// Cursor will return the buffer's cursor
func (b *Buffer) Cursor() cursor.Cursorer { return b.cur }

// Dirty will return true if the buffer is unsaved
func (b *Buffer) Dirty() bool { return b.dirty }

func (b *Buffer) listenUpdates() {
	for {
		up := <-b.updates
		if up {
			b.matchSyntax()
		}
	}
}

func (b *Buffer) matchSyntax() {
	b.syntaxResults = b.mat.Parse(b.txt)
}

// SyntaxResults returns the syntax scanning results
func (b *Buffer) SyntaxResults() []syntax.Resulter {
	return b.syntaxResults
}

// New will return a new Buffer
func New(txt textstore.TextStorer, cur cursor.Cursorer, m syntax.Matcherer, ftd filetype.Detecter) Bufferer {
	up := make(chan bool)
	id := uuid.New().String()
	b := &Buffer{
		ftd:     ftd,
		txt:     txt,
		cur:     cur,
		mat:     m,
		updates: up,
		id:      id,
	}
	txt.Subscribe(b.id, b.updates)
	go b.listenUpdates()
	return b
}
