package textstore

import (
	"os"
	"testing"

	"github.com/dshills/layered/undo"
)

var store = New(undo.New)

func TestReset(t *testing.T) {
	str := "This is a test"
	store.Reset(str)
}
func TestReadFrom(t *testing.T) {
	f, err := os.Open("/Users/dshills/Development/projects/layered/testdata/scanner.go")
	if err != nil {
		t.Fatal(err)
	}
	store.ReadFrom(f)
}
func TestNewLine(t *testing.T) {
	store.NewLine("", 0)
}

func TestDeleteLine(t *testing.T) {
	store.DeleteLine(1)
}

func TestResetLine(t *testing.T) {
	store.ResetLine("Ta-da", 1)
}

func TestLineString(t *testing.T) {
	store.LineString(0)
}
func TestLineRangeString(t *testing.T) {
	store.LineRangeString(0, 10)
}
func TestNumLines(t *testing.T) {
	store.NumLines()
}
func TestLineLen(t *testing.T) {
	store.LineLen(0)
}
func TestLen(t *testing.T) {
	store.Len()
}
func TestReadRuneAt(t *testing.T) {
	store.ReadRuneAt(0, 0)
}
func TestLineAt(t *testing.T) {
	store.LineAt(0)
}
func TestLineWriterAt(t *testing.T) {
	store.LineWriterAt(1)
}
func TestSetLineDelim(t *testing.T) {
	store.SetLineDelim("\n")
}
func TestLineDelim(t *testing.T) {
	store.LineDelim()
}
func TestUndo(t *testing.T) {
	store.Undo()
}
func TestRedo(t *testing.T) {
	store.Redo()
}
func TestStartGroupUndo(t *testing.T) {
	store.StartGroupUndo()
}
func TestStopGroupUndo(t *testing.T) {
	store.StopGroupUndo()
}
