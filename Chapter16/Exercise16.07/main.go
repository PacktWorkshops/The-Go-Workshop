package main

import (
	"log"
)

func push(from,to int, in chan bool, out chan int) {
	for i:=from;i<=to; i++ {
		<- in
		out <- i
	}
}

func main() {
	s1 := 0

	out := make(chan int, 100)
	in := make(chan bool,100)

	go push(1,25, in, out)
	go push(26,50,in, out)
	go push(51,75,in, out)
	go push(76,100, in, out)

	for c :=0; c< 100; c++ {
		in <- true
		i := <- out
		log.Println(i)
		s1 += i
	}

	log.Println(s1)

}
