package main

import (
	"testing"
)

func TestConversion(t *testing.T) {
	if res, err := doubler(5); res != "10" && err != nil {
		t.Fail()
	}
	if res, err := doubler("yum"); res != "yumyum" && err != nil {
		t.Fail()
	}
	if _, err := doubler(3.14); err == nil {
		t.Fail()
	}
}
