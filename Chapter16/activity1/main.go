package main

import (
	"fmt"
	"log"
	"sync"
)

func sum(from,to int, wg *sync.WaitGroup, res *string, mtx *sync.Mutex) {
	for i:=from;i<=to; i++ {
		mtx.Lock()
		*res = fmt.Sprintf("%s|%d|",*res, i)
		mtx.Unlock()
	}

	wg.Done()

	return
}

func main() {
	s1 := ""
	mtx := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go sum(1,25, wg,&s1, mtx)
	go sum(26,50, wg, &s1, mtx)
	go sum(51,75, wg, &s1, mtx)
	go sum(76,100, wg, &s1, mtx)
	wg.Wait()

	log.Println(s1)

}
