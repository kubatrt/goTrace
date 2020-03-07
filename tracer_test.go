package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {

	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Function can not return nil!")
	} else {
		tracer.Trace("Trace package under test.")
		if buf.String() != "Trace package under test.\n" {
			t.Errorf("Method Trace sould not print '%s'.", buf.String())
		}
	}
}
