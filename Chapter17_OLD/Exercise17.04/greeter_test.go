package greeter

import (
	"testing"
)

func BenchmarkGreet(b *testing.B) {

	msg := "Hello, I'm John"

	in := make(chan string, 1)
	out := make(chan string, 1)
	er := make(chan error, 1)

	go Greet(in, out, er)

	for i:=0;i< b.N; i++ {

			in <- msg

			<- out
	}

}