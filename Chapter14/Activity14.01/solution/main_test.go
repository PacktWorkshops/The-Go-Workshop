package main

import (
	"testing"
)

// TestGetDataAndParseResponse requires the server to be running to succeed
func TestGetDataAndParseResponse(t *testing.T) {
	electricCount, boogalooCount := getDataAndParseResponse()
	if electricCount < 1 {
		t.Errorf("expected more than one name 'Electric', recieved: %d", electricCount)
	}
	if boogalooCount < 1 {
		t.Errorf("expected more than one name 'Boogaloo', recieved: %d", boogalooCount)
	}
}
