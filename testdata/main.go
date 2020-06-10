package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/zyedidia/tcell/terminfo"
)

const defLog = ".goed/error.log"

func main() {
	start := time.Now()
	f, err := startLogger(defLog)
	if err != nil {
		fmt.Printf("Opening log %v\n", err)
		os.Exit(1)
	}
	defer func() {
		/*
			if r := recover(); r != nil {
				logDebugf(r)
				logDebugf(getStackTrace())
				fmt.Println(r)
			}
		*/
		f.Close()
	}()
	term, err := NewTerminal()
	if err != nil {
		fmt.Printf("Failed creating terminal. Exiting... %v\n", err)
		os.Exit(1)
	}

	tm := os.Getenv("TERM")
	ti, err := terminfo.LookupTerminfo(tm)
	if err != nil {
		logErrorf("terminfo %v", err)
	} else {
		logDebugf("%v $TERM: %v Lines: %v Cols: %v Colors: %v", ti.Name, tm, ti.Lines, ti.Columns, ti.Colors)
	}

	var wg sync.WaitGroup
	go term.Start(&wg)
	ed, err := NewEditor(term)
	if err != nil {
		fmt.Printf("Failed creating editor. Exiting... %v\n", err)
		os.Exit(1)
	}
	fname := "testdata/main.go"
	if err := ed.FileBuffer(fname); err != nil {
		logErrorf("editor.Open %v %v", fname, err)
	}
	logDebugf("Total load elapsed %v", time.Since(start))
	wg.Wait()
	/*
		Multiline comment
	*/
}

func getStackTrace() string {
	stackBuf := make([]uintptr, 50)
	length := runtime.Callers(3, stackBuf[:])
	stack := stackBuf[:length]

	trace := ""
	frames := runtime.CallersFrames(stack)
	for {
		frame, more := frames.Next()
		trace = trace + fmt.Sprintf("\n\tFile: %s, Line: %d. Function: %s", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
	return trace
}
