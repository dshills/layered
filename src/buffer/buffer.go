package buffer

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/register"
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
	updates       chan uint64
	syntaxResults []syntax.Resulter
	searchResults []SearchResult
	reg           register.Registerer
	txthash       uint64
}

// SearchResult is the results for a line
type SearchResult struct {
	Line    int
	Matches [][]int
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
		h := <-b.updates
		if h != b.txthash {
			b.dirty = true
			b.txthash = h
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

// SearchResults will return the current search results
func (b *Buffer) SearchResults() []SearchResult {
	return b.searchResults
}

// Search will search the current textstore
func (b *Buffer) Search(s string) ([]SearchResult, error) {
	rex, err := regexp.Compile(s)
	if err != nil {
		return nil, err
	}
	lns := b.txt.NumLines()
	wg := sync.WaitGroup{}
	for i := 0; i < lns; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, ln int) {
			str, _ := b.txt.LineString(ln)
			matches := rex.FindAllStringIndex(str, -1)
			if len(matches) > 0 {
				b.searchResults = append(b.searchResults, SearchResult{Line: ln, Matches: matches})
			}
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	if len(b.searchResults) == 0 {
		return nil, fmt.Errorf("Not found")
	}
	return b.searchResults, nil
}

// New will return a new Buffer
func New(txt textstore.TextStorer, cur cursor.Cursorer, m syntax.Matcherer, ftd filetype.Detecter, reg register.Registerer) Bufferer {
	up := make(chan uint64)
	id := uuid.New().String()
	b := &Buffer{
		ftd:     ftd,
		txt:     txt,
		cur:     cur,
		mat:     m,
		updates: up,
		id:      id,
		reg:     reg,
	}
	txt.Subscribe(b.id, b.updates)
	go b.listenUpdates()
	return b
}
