package main

import (
	"bytes"
	"log"
	"testing"
)

func Test_Main(t *testing.T) {
	var s bytes.Buffer
	log.SetOutput(&s)

	main()
	expected := "a\nb\nc\nd\ne\nf\n"
	if s.String() != expected {
		t.Errorf("Expected %s but received %s", expected, s.String())
	}
}
