package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		name      string
		input     int
		wantedInt int
		wantedStr string
	}{

		{
			name:      "Fizz",
			input:     6,
			wantedInt: 6,
			wantedStr: "Fizz",
		},

		{
			name:      "Buzz",
			input:     10,
			wantedInt: 10,
			wantedStr: "Buzz",
		},

		{
			name:      "FizzBuzz",
			input:     15,
			wantedInt: 15,
			wantedStr: "FizzBuzz",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotInt, gotStr := fizzBuzz(tc.input)

			if gotInt != tc.wantedInt || gotStr != tc.wantedStr {
				t.Errorf("Wanted integer: %v, got %v, wanted string %v got %v ", tc.wantedInt, gotInt, tc.wantedStr, gotStr)
			}

		})

	}

}
