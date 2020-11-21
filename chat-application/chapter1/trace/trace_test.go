package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	// chúng ta cần có thể capture output của tracer vào kiểu bytes.Buffer
	// để có thể đảm bảo rằng string trong buffer khớp với giá trị mong muốn
	// nếu nó không khớp method t.Errorf sẽ được gọi để fail the test
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("error could not create tracer")
	} else {
		// trước đó chúng ta kiểm tra để đảm bảo rằng sự trả về từ hàm New sẽ không
		// là nil, nếu là nil the test sẽ fail bằng cách gọi t.Error
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}
