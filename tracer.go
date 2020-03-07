package trace

import "io"

// Tracer is an interface to trace events in the system at runtime
type Tracer interface {
	Trace(...interface{})
}

// New creates instance of Tracer parametrized with compatible
// io.Writer interface, for example: stdio, socket, bytes.Buffer
func New(w io.Writer) Tracer {
	return nil
}
