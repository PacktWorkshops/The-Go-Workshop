package greeter

import (
	"errors"
	"fmt"
)

func Greet(in,out chan string, er chan error) {
	for {
		msg := <- in

		prefix := "Hello, I'm "
		if len(msg) <= len(prefix) || msg[:len(prefix)] != prefix {
			fmt.Println( msg[:len(prefix)])
			er <- errors.New("wrong string")
		}

		name := msg[len(prefix):]
		out <- fmt.Sprintf("Hi %s, nice to meet you", name)
	}
}
