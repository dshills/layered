package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// GetStackTrace returns a stack trace
func GetStackTrace(max int) string {
	stackBuf := make([]uintptr, 50)
	length := runtime.Callers(3, stackBuf[:])
	stack := stackBuf[:length]

	builder := strings.Builder{}
	frames := runtime.CallersFrames(stack)
	for i := 0; i < max; i++ {
		frame, more := frames.Next()
		rpath := filepath.Base(frame.File)
		fpath := filepath.Base(frame.Function)
		builder.WriteString(fmt.Sprintf("%s:%v %s\n", rpath, frame.Line, fpath))
		if !more {
			break
		}
	}
	return builder.String()
}
