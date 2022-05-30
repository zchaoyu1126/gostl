package math

import (
	"testing"
)

func TestMinInt32(t *testing.T) {
	var a, b int32 = 9, 15
	var expected int32 = 9
	output := Min(a, b)
	if expected != output {
		t.Errorf("expected:%d, but get %d", expected, output)
	}
}

func TestMinFloat64(t *testing.T) {
	var a, b float64 = 1.24, 3.95
	var expected float64 = 1.24
	output := Min(a, b)
	if expected != output {
		t.Errorf("expected:%g, but get %g", expected, output)
	}
}

func TestMaxInt32(t *testing.T) {
	var a, b int32 = 9, 15
	var expected int32 = 15
	output := Max(a, b)
	if expected != output {
		t.Errorf("expected:%d, but get %d", expected, output)
	}
}

func TestMaxFloat64(t *testing.T) {
	var a, b float64 = 1.24, 3.95
	var expected float64 = 3.95
	output := Max(a, b)
	if expected != output {
		t.Errorf("expected:%g, but get %g", expected, output)
	}
}
