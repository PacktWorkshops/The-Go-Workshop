package main

import "testing"

func TestString(t *testing.T) {
	testCases := []struct {
		name      string
		age       int
		isMarried bool
		wanted    string
	}{

		{
			name:      "Cayden Smash",
			age:       4,
			isMarried: false,
			wanted:    "Cayden Smash (4 years old).\nMarried status: false ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			p := person{name: tc.name, age: tc.age, isMarried: tc.isMarried}
			got := p.String()

			if got != tc.wanted {
				t.Errorf("Got   : %v\n Wanted: %v\n", got, tc.wanted)
			}
		})
	}

}

func TestSpeak(t *testing.T) {
	testCases := []struct {
		name   string
		wanted string
	}{

		{
			name:   "Cayden Smash",
			wanted: "Hi my name is: Cayden Smash",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			p := person{name: tc.name}
			got := p.Speak()

			if got != tc.wanted {
				t.Errorf("Got   : %v\n Wanted: %v\n", got, tc.wanted)
			}
		})
	}

}
