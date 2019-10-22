package main

import (
	"fmt"
	"testing"
)

func Test_Sum(t *testing.T) {
	for i :=4;i<20; i++ {
		res := sum(i, 1,100)
		if res != 5050 {
			t.Errorf("We were expecting 5050 with %d workers but we received %d", i, res)
		}
	}
}

func Benchmark_Sum(b *testing.B) {
	for i :=4;i<20; i++ {
		b.Run(fmt.Sprintf("%d wrokers", i), func(b *testing.B) {
			for c:=0;c<b.N; c++{
				res := sum(i, 1,100)
				if res != 5050 {
					b.Errorf("We were expecting 5050 with %d workers but we received %d", i, res)
				}
			}
		})
	}
}