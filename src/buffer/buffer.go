package buffer

import (
	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/textstore"
)

// Buffer is a text buffer
type Buffer struct {
	id  string
	fn  string
	ft  string
	cur cursor.Cursorer
	txt textstore.TextStorer
}

// ID will return the identifier for the buffer
func (b *Buffer) ID() string { return b.id }

// Filename will return the buffer filename
func (b *Buffer) Filename() string { return b.fn }

// SetFilename will set the buffers file name
func (b *Buffer) SetFilename(n string) { b.fn = n }

// Filetype returns the buffer's file type
func (b *Buffer) Filetype() string { return b.ft }

// SetFiletype will set the buffer's file type
func (b *Buffer) SetFiletype(ft string) { b.ft = ft }

// TextStore will return the buffer's text store
func (b *Buffer) TextStore() textstore.TextStorer { return b.txt }

// Cursor will return the buffer's cursor
func (b *Buffer) Cursor() cursor.Cursorer { return b.cur }

// New will return a new Buffer
func New(txt textstore.TextStorer, cur cursor.Cursorer) Bufferer {
	return &Buffer{txt: txt, cur: cur}
}
