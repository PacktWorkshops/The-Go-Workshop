package main

import (
	"log"
	"time"
)

func push(from,to int, out chan int) {
	for i:=from;i<=to; i++ {
		out <- i
		time.Sleep(time.Microsecond)
	}
}

func main() {
	s1 := 0

	ch := make(chan int, 100)

	go push(1,25, ch)
	go push(26,50,ch)
	go push(51,75,ch)
	go push(76,100, ch)

	for c :=0; c< 100; c++ {
		i := <- ch
		log.Println(i)
		s1 += i
	}

	log.Println(s1)

}
