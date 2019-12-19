package exercise3

import (
	"fmt"
	"testing"
)

func TestQuadratic(t *testing.T) {

	tcs := []struct{expected, tested float32} {
		{3,0},
		{10, 1},
		{21,2},
		{45, 3.5},
		{0, -1},
	}

	for _, k := range tcs {
		t.Run(fmt.Sprintf("Testing %f", k.tested), func(t *testing.T) {
			computed := quadraticFit(k.tested)
if k.expected != computed {
t.Errorf("expected %f but received %f", k.expected, computed)
}
})
}
}
