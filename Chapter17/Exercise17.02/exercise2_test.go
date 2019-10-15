package exercise2

import "testing"

func TestLinearFit(t *testing.T) {
	r1 := linearFit(0)
	if r1 != float32(5) {
		t.Errorf("Expected 5 but received %f", r1)
	}

	r2 := linearFit(1)
	if r2 != float32(8) {
		t.Errorf("Expected 8 but received %f", r2)
	}

	r3 := linearFit(2)
	if r3 != float32(11) {
		t.Errorf("Expected 11 but received %f", r3)
	}

	r4 := linearFit(3)
	if r4 != float32(11) {
		t.Errorf("Expected 11 but received %f", r4)
	}

	t.Run("Subtest name", func(t *testing.T) {

	})
}
