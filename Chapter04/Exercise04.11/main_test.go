package main

import "testing"

func TestGenSlices(t *testing.T) {
	s1, s2, s3 := genSlices()
	if len(s1) != 0 || cap(s1) != 0 {
		t.Fail()
	}
	if len(s2) != 10 || cap(s2) != 10 {
		t.Fail()
	}
	if len(s3) != 10 || cap(s3) != 50 {
		t.Fail()
	}
}
