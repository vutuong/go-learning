package trace

import (
	"fmt"
	"io"
)

// Tracer là interface mô tả một object có thể tra vết các event xuyên qua code
type Tracer interface {
	Trace(...interface{})
}

// New tạo mới một Tracer mà nó sẽ ghi output vào một io.Write cụ thể
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// tracer là một Tracer mà ghi đến một io.Writer
type tracer struct {
	out io.Writer
}

// Trace writes the arguments to this Tracers io.Writer.
func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates a Tracer that will ignore calls to Trace.
func Off() Tracer {
	return &nilTracer{}
}
