package main

import (
	"io/ioutil"
	"strings"
	"sync"
	"testing"
)

func Test_Main(t *testing.T) {
	main()
	bts, err := ioutil.ReadFile("result.txt")
	if err != nil {
		t.Error(err)
	}

	if len(string(bts)) != 14 {
		t.Error("Wrong string", string(bts))
	}

	evens := strings.ReplaceAll(string(bts), "Odd 9\n", "")
	if len(evens) >= 14 {
		t.Error("No odds removed")
	}

	empty := strings.ReplaceAll(evens, "Even 12\n", "")
	if len(empty) !=0 {
		t.Error("non empty string", empty)
	}
}

func TestSource(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	out := make(chan int)

	go source("./input1.dat", out, wg)

	nums := []int{}
	for i:=0;i<3;i++ {
		nums = append(nums, <-out)
	}

	for k,v:= range []int{1,2,5} {
		if nums[k] != v {
			t.Errorf("Wrong inputs, expected on index %d value %d but received %d", k,v,nums[k])
		}
	}
}

func TestSplitter(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	in := make(chan int, 10)
	odd := make(chan int)
	even := make(chan int)
	go splitter(in, odd, even, wg)

	for i:=1;i<=10; i++ {
		in <- i
	}
	close(in)

	oddSum, evenSum := 0,0

	for i:=1;i<=10; i++ {
		select {
		case i := <- odd:
			oddSum += i
		case i := <- even:
			evenSum += i
		}
	}

	wg.Wait()


	if oddSum != 1+3+5+7+9 {
		t.Error("Odds shoudl not be ", oddSum)
	}

	if evenSum != 2+4+6+8+10 {
		t.Error("Odds shoudl not be ", evenSum)
	}
}

func TestSum(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	in := make(chan int, 10)
	out := make(chan int, 1)

	go sum(in,out,wg)

	for i:=0;i<10;i++ {
		in <- 1
	}
	close(in)

	sm  := <- out
	if sm != 10 {
		t.Errorf("Expected 10 but received %d", sm)
	}
}

func TestMerger(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	even := make(chan int, 10)
	odd := make(chan int, 1)
	outFl := "test_output.txt"

	go merger(even, odd,  wg , outFl)

	even <- 10
	odd <- 20
	wg.Wait()

	bts, err := ioutil.ReadFile(outFl)
	if err != nil {
		t.Error(err)
	}

	if len(string(bts)) != 15 {
		t.Error("Wrong string", string(bts), len(string(bts)))
	}

	evens := strings.ReplaceAll(string(bts), "Odd 20\n", "")
	if len(evens) >= 14 {
		t.Error("No odds removed")
	}

	empty := strings.ReplaceAll(evens, "Even 10\n", "")
	if len(empty) !=0 {
		t.Error("non empty string", empty)
	}

}