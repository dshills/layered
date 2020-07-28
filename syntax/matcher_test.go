package syntax

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/dshills/layered/textstore"
	"github.com/dshills/layered/undo"
)

func TestMatch(t *testing.T) {
	rt, err := filepath.Abs("./testdata")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open("testdata/scanner.go")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	store := textstore.New(undo.New)
	if _, err := store.ReadFrom(f); err != nil {
		t.Fatal(err)
	}

	mgr := New(rt)
	if err := mgr.LoadFileType("go"); err != nil {
		t.Error(err)
	}

	start := time.Now()
	results := mgr.Parse(store)
	fmt.Printf("Elapsed %v\n", time.Since(start))
	if len(results) == 0 {
		t.Errorf("Expected results got none")
	}
}
