package greeter

import (
	"testing"
)

func TestGreeter(t *testing.T) {

	tcs := []struct{valid bool; msg string}{
		{true, "Hello, I'm John"},
		{false, "Hi, My name is Ethan"},
		{true, "Hello, I'm David"},
	}

	for _, tc := range tcs {
		t.Run(tc.msg, func(t *testing.T) {
			in := make(chan string, 1)
			out := make(chan string, 1)
			er := make(chan error, 1)

			go Greet(in, out, er)

			in <- tc.msg

			select {
			case msg := <- out:
				if !tc.valid {
					t.Errorf("The message: '%s' is correct but we were expecting a wrong one", msg)
				}
			case  <- er:
				if tc.valid {
					t.Errorf("The message: '%s' is wrong but we were expecting a correct message and expected validity is %t", tc.msg, tc.valid)
				}
			}
		})
	}

}