package main

import (
	"context"
	"log"
	"time"
)

func countNumbers(c context.Context, r chan int) {
	v := 0
	for {
		select {
		case <-c.Done():
			r <- v
			break
		default:
			time.Sleep(time.Millisecond * 100)
			v++
		}
	}
}


func main() {
	r := make(chan int)
	c := context.TODO()
	cl, stop := context.WithCancel(c)
	go countNumbers(cl, r)

	go func() {
		time.Sleep(time.Millisecond*100*3)
		stop()
	}()

	v := <- r

	log.Println(v)
}