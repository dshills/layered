package textstore

import (
	"testing"
	"unicode/utf8"
)

const testStr = "This is a test of the emergency broadcast system. This is only a test."

func TestLineReset(t *testing.T) {
	l := newLine(testStr)
	s := l.String()
	if s != testStr {
		t.Errorf("Expected %v got %v", testStr, s)
	}
	l.reset(testStr)
	s = l.String()
	if s != testStr {
		t.Errorf("Expected %v got %v", testStr, s)
	}

}

func TestInsertRune(t *testing.T) {
	l := newLine(testStr)
	l.insertRune(0, '^')
	s := l.String()
	r, _ := utf8.DecodeRuneInString(s)
	if r != '^' {
		t.Errorf("Expected ^ got %v", r)
	}
}
func TestReplaceRune(t *testing.T) {
}
func TestDelete(t *testing.T) {
}
func TestInsertString(t *testing.T) {
}
func TestReplaceString(t *testing.T) {
}
func TestStartEdit(t *testing.T) {
}
func TestEndEdit(t *testing.T) {
}
func TestChangeSet(t *testing.T) {
}
func TestLenLine(t *testing.T) {
}
func TestString(t *testing.T) {
}
func TestRuneAt(t *testing.T) {
}
func TestDelimited(t *testing.T) {
}
