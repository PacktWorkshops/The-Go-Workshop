package exercise1

import "testing"

func TestMaths(t *testing.T) {
	r1 := 2*2
	if r1 != 4 {
		t.Errorf("we expected 4 but we received %d", r1)
	}

	r2 := 2*3
	if r2 != 6 {
		t.Errorf("we expected 6 but we received %d", r1)
	}

	r3 := "Hello world"
	if r3 != "Hello world" {
		t.Errorf("we expected 'Hello world' but we received %s", r3)
	}
}
