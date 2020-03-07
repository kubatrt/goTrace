package trace

import (
	"fmt"
	"io"
)

// Tracer is an interface to trace events in the system at runtime
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New creates instance of Tracer parametrized with compatible
// io.Writer interface, for example: stdio, socket, bytes.Buffer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Struct for empty Tracer
type nilTracer struct{}

// Do nothing
func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates empty tracer
func Off() Tracer {
	return &nilTracer{}
}
