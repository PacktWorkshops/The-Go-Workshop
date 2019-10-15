package main

import (
	"bytes"
	"log"
	"testing"
)

func Test_Main(t *testing.T) {

	for i:=0;i < 10000; i++ {
		var s bytes.Buffer
		log.SetOutput(&s)
		log.SetFlags(0)
		main()

		if s.String() != "5050\n" {
			t.Error(s.String())
		}
	}
}
