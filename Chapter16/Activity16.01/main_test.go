package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func Test_Main(t *testing.T) {

		var s bytes.Buffer
		log.SetOutput(&s)
		log.SetFlags(0)
		main()

		// 90 numbers of two digits |xx| and 9 numbers of 1 digit |9| plus newline \n
		expectedLength := (3*9)+ (4*91)+2

		if len(s.String()) != expectedLength {
			t.Errorf("Exppecting length of %d but received %d",  expectedLength, len(s.String()) )
		}

		res := s.String()
		for i:=1; i<=100; i++ {
			str := fmt.Sprintf("|%d|", i)
			expectedLength -= len(str)
			res = strings.ReplaceAll(res, fmt.Sprintf("|%d|", i), "")

			if len(res) !=expectedLength {
				t.Errorf("Expected %d but received %d", expectedLength, len(res))
			}
		}
}
