package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func source(filename string, out chan int, wg *sync.WaitGroup)  {
	f, err :=  os.Open(filename)
	if err != nil {
		panic(err)
	}

	rd := bufio.NewReader(f)
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				wg.Done()
				return
			} else {
				panic(err)
			}
		}
		iStr := strings.ReplaceAll(str, "\n", "")
		i, err := strconv.Atoi(iStr)
		if err != nil {
			panic(err)
		}
		out <- i
	}
}

func splitter(in, odd, even chan int, wg *sync.WaitGroup)  {
	for i := range in {
		switch i%2 {
		case 0:
			even <- i
		case 1:
			odd <- i
		}
	}

	close(even)
	close(odd)
	wg.Done()
}

func sum(in, out chan int, wg *sync.WaitGroup) {
	sum := 0
	for i := range in {
		sum += i
	}
	out <- sum
	wg.Done()
}

func merger(even, odd chan int, wg *sync.WaitGroup, resultFile string) {
	rs, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}
	for i:= 0; i< 2; i++{
		select {
		case i:= <- even:
			rs.Write([]byte(fmt.Sprintf("Even %d\n", i)))
		case i:= <- odd:
			rs.Write([]byte(fmt.Sprintf("Odd %d\n", i)))
		}

	}
	rs.Close()
	wg.Done()
}


func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	wg2 := &sync.WaitGroup{}
	wg2.Add(4)
	odd := make(chan int)
	even := make(chan int)
	out := make(chan int)
	sumodd := make(chan int)
	sumeven := make(chan int)
	go source("./input1.dat", out, wg)
	go source("./input2.dat", out, wg)
	go splitter(out, odd, even, wg2)
	go sum(even, sumeven, wg2)
	go sum(odd, sumodd,wg2)
	go merger(sumeven, sumodd, wg2, "./result.txt")
	wg.Wait()
	close(out)
	wg2.Wait()
}
