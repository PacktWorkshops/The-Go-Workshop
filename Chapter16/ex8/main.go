package main

import (
	"log"
)

func worker(in <- chan int, out chan int) {
	sum := 0
	for i := range in {
		sum += i
	}

	out <- sum
}

func sum(workers, from, to int) int {

	out := make(chan int, workers)
	in := make(chan int,4)

	for i:=0;i<workers;i++ {
		go worker(in, out)
	}

	for i:=from;i<=to; i++ {
		in <- i
	}

	close(in)

	sum := 0
	for i:=0;i<workers; i++ {
		sum += <-out
	}

	close(out)

	return sum
}

func main() {

	res := sum(100,1,100)
	log.Println(res)
}
