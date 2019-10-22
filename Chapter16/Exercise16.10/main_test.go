package main

import (
	"sync"
	"testing"
)

func TestSum(t *testing.T) {

	for c:=1; c<=100; c++ {
		mtx := &sync.Mutex{}
		in := make(chan int, 100)
		out := make(chan int)
		wrk := Worker{in: in, out:out, mtx:mtx}

		for w:=1; w<=c; w++ {
			wrk.readThem()
		}

		for i:=1;i<=100; i++ {
			in <- i
		}
		close(in)

		res := wrk.gatherResult()

		if res != 5050 {
			t.Errorf("Expected 5050 but received %d", res)
		}
	}

}
