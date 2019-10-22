package main

import (
	"log"
)

func greet(ch chan string) {
	ch <- "Hello"
}
func main() {
	ch := make(chan string)

	go greet(ch)

	log.Println(<-ch)

}
