package main

import (
	"reflect"
	"testing"
)

func TestGetLocales(t *testing.T) {
	if !reflect.DeepEqual(getLocals([]string{"fr_CN", "en_AU"}), []string{"en_US", "fr_FR", "fr_CN", "en_AU"}) {
		t.Fail()
	}
}
