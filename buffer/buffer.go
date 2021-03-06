package buffer

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/filetype"
	"github.com/dshills/layered/logger"
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
	cur           cursor.Cursor
	txt           textstore.TextStorer
	mat           syntax.Manager
	ftd           filetype.Manager
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

// Content will return the content within a range
func (b *Buffer) Content(start, cnt int) ([]string, error) {
	return b.txt.LineRange(start, cnt)
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
	if err := b.mat.LoadFileType(ft); err != nil {
		logger.Errorf("Buffer.SetFiletype: %v", err)
	}
	b.matchSyntax()
}

// TextStore will return the buffer's text store
func (b *Buffer) TextStore() textstore.TextStorer { return b.txt }

// Cursor will return the buffer's cursor
func (b *Buffer) Cursor() cursor.Cursor { return b.cur }

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

func (b *Buffer) matchSyntax(fgrps ...string) {
	b.syntaxResults = b.mat.Parse(b.txt, fgrps...)
}

// SyntaxResults returns the syntax scanning results
func (b *Buffer) SyntaxResults(update bool, fgrps ...string) []syntax.Resulter {
	if len(fgrps) > 0 || update {
		b.matchSyntax(fgrps...)
	}
	return b.syntaxResults
}

// SyntaxResultsRange returns the syntax scanning results
func (b *Buffer) SyntaxResultsRange(ln, cnt int, update bool, fgrps ...string) []syntax.Resulter {
	if len(fgrps) > 0 || update {
		b.matchSyntax(fgrps...)
	}
	mx := ln + cnt
	res := []syntax.Resulter{}
	for _, r := range b.syntaxResults {
		tln := r.Line()
		if tln >= ln && tln < mx {
			res = append(res, r)
		}
	}
	return res
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
func New(txt textstore.TextStorer, cur cursor.Cursor, m syntax.Manager, ftd filetype.Manager, reg register.Registerer) Bufferer {
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
