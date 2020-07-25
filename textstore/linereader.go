package textstore

import (
	"errors"
	"io"
	"unicode/utf8"
)

// A Reader implements the io.Reader, io.ReaderAt, io.Seeker, io.WriterTo,
// io.ByteScanner, and io.RuneScanner interfaces by reading
// from a string.
// The zero value for Reader operates like a Reader of an empty string.
type Reader struct {
	i  int64 // current reading index
	pr int   // index of previous rune; or < 0
	ln *Line
}

func (r *Reader) String() string {
	return r.ln.String()
}

// Len returns the number of bytes of the unread portion of the
// string.
func (r *Reader) Len() int {
	if r.i >= int64(len(r.ln.text)) {
		return 0
	}
	return int(int64(len(r.ln.text)) - r.i)
}

// Size returns the original length of the underlying string.
// Size is the number of bytes available for reading via ReadAt.
// The returned value is always the same and is not affected by calls
// to any other method.
func (r *Reader) Size() int64 { return int64(len(r.ln.text)) }

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.ln.text)) {
		return 0, io.EOF
	}
	r.pr = -1
	n = copy(b, r.ln.text[r.i:])
	r.i += int64(n)
	return
}

// ReadAt will into p at position off
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) {
	// cannot modify state - see io.ReaderAt
	if off < 0 {
		return 0, errors.New("ReadAt: negative offset")
	}
	if off >= int64(len(r.ln.text)) {
		return 0, io.EOF
	}
	n = copy(b, r.ln.text[off:])
	if n < len(b) {
		err = io.EOF
	}
	return
}

// ReadRuneAt returns the rune at rune offset (not byte offset)
func (r *Reader) ReadRuneAt(offset int64) (rune, int, error) {
	if offset >= int64(len(r.ln.text)) {
		return 0, 0, io.EOF
	}
	if c := r.ln.text[offset]; c < utf8.RuneSelf {
		return rune(c), 1, nil
	}
	ch, size := utf8.DecodeRuneInString(r.ln.text[offset:])
	return ch, size, nil
}

// ReadByte will read the next byte from the seek position
func (r *Reader) ReadByte() (byte, error) {
	r.pr = -1
	if r.i >= int64(len(r.ln.text)) {
		return 0, io.EOF
	}
	b := r.ln.text[r.i]
	r.i++
	return b, nil
}

// UnreadByte will decrement the seek position by 1 byte
func (r *Reader) UnreadByte() error {
	if r.i <= 0 {
		return errors.New("UnreadByte: at beginning of string")
	}
	r.pr = -1
	r.i--
	return nil
}

// ReadRune will read the next rune from the seek position
func (r *Reader) ReadRune() (ch rune, size int, err error) {
	if r.i >= int64(len(r.ln.text)) {
		r.pr = -1
		return 0, 0, io.EOF
	}
	r.pr = int(r.i)
	if c := r.ln.text[r.i]; c < utf8.RuneSelf {
		r.i++
		return rune(c), 1, nil
	}
	ch, size = utf8.DecodeRuneInString(r.ln.text[r.i:])
	r.i += int64(size)
	return
}

// UnreadRune will decrement the seek position by size of last rune
func (r *Reader) UnreadRune() error {
	if r.i <= 0 {
		return errors.New("UnreadRune: at beginning of string")
	}
	if r.pr < 0 {
		return errors.New("UnreadRune: previous operation was not ReadRune")
	}
	r.i = int64(r.pr)
	r.pr = -1
	return nil
}

// Seek implements the io.Seeker interface.
func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	r.pr = -1
	var abs int64
	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = r.i + offset
	case io.SeekEnd:
		abs = int64(len(r.ln.text)) + offset
	default:
		return 0, errors.New("Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("Seek: negative position")
	}
	r.i = abs
	return abs, nil
}

// WriteTo implements the io.WriterTo interface.
func (r *Reader) WriteTo(w io.Writer) (n int64, err error) {
	r.pr = -1
	if r.i >= int64(len(r.ln.text)) {
		return 0, nil
	}
	s := r.ln.text[r.i:]
	m, err := io.WriteString(w, s)
	if m > len(s) {
		panic("WriteTo: invalid WriteString count")
	}
	r.i += int64(m)
	n = int64(m)
	if m != len(s) && err == nil {
		err = io.ErrShortWrite
	}
	return
}

// NewReader returns a new Reader reading from s.
// It is similar to bytes.NewBufferString but more efficient and read-only.
func NewReader(l *Line) *Reader { return &Reader{ln: l, pr: -1} }
