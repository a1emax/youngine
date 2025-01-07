package fault

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// Trace returns given error wrapped to attach calling goroutine's stack trace, or nil if error is nil.
//
// Due to many various implementations of error handling in Golang ecosystem, it is strongly recommended to
// wrap only errors which format you know, because any other errors can already contain stack trace in one
// form or another. Also do not "rewrap" errors except if you really want to get multiple stack trace.
func Trace(err error) error {
	if err == nil {
		return nil
	}

	// There we limit recorded frames number up to 16 that is probably sufficient stack depth in most cases.
	pc := make([]uintptr, 16)
	n := runtime.Callers(2, pc)

	return tracedError{err, pc[:n]}
}

// tracedError is error that wraps base error by attaching goroutine's stack trace.
type tracedError struct {
	err error
	pc  []uintptr
}

func (e tracedError) Error() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("%+v", e.err))

	// This marker allows to detect empty or multiple stack trace.
	b.WriteString("\n\t[stack trace]")

	frames := runtime.CallersFrames(e.pc)
	for {
		frame, more := frames.Next()

		b.WriteString("\n\t")
		if frame.Function != "" {
			b.WriteString(frame.Function)
		} else {
			b.WriteString("<unknown>")
		}
		if frame.File != "" {
			b.WriteString("\n\t\t")
			b.WriteString(frame.File)
			b.WriteRune(':')
			b.WriteString(strconv.Itoa(frame.Line))
		}

		if !more {
			break
		}
	}

	return b.String()
}

// Unwrap implements the internal interface of the [errors.Unwrap] function.
func (e tracedError) Unwrap() error {
	return e.err
}
