package math

import (
	"testing"
)

func TestAbsInt32(t *testing.T) {
	var n int32 = -4
	var expected int32 = 4
	output := Abs(n)
	if expected != output {
		t.Errorf("expected:%d, but get %d", expected, output)
	}
}

func TestAbsFloat64(t *testing.T) {
	var n float64 = -3.1415926
	var expected float64 = 3.1415926
	output := Abs(n)
	if expected != output {
		t.Errorf("expected:%g, but get %g", expected, output)
	}
}
