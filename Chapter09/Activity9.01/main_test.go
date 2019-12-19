package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestValidateLength(t *testing.T) {
	testCases := []struct {
		name   string
		ssn    string
		wanted error
	}{
		{
			name:   "Test-Basic",
			ssn:    "123456789",
			wanted: nil,
		},
		{
			name:   "Test SSN spaces",
			ssn:    "123456789   ",
			wanted: nil,
		},
		{
			name:   "Test SSN less than 9",
			ssn:    "1234",
			wanted: fmt.Errorf("the value of %s caused an error: %v\n", "1234", ErrInvalidSSNLength),
		},
		{
			name:   "Test SSN more than 9",
			ssn:    "12345467890",
			wanted: fmt.Errorf("the value of %s caused an error: %v\n", "12345467890", ErrInvalidSSNLength),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := validLength(tc.ssn)

			if got != tc.wanted {
				if !strings.Contains(tc.wanted.Error(), got.Error()) {
					t.Errorf("\nGot    :%v\n wanted:%v", got, tc.wanted)
				}
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	testCases := []struct {
		name   string
		ssn    string
		wanted error
	}{
		{
			name:   "Test-Basic",
			ssn:    "123456789",
			wanted: nil,
		},
		{
			name:   "Test SSN is not a number",
			ssn:    "1234w",
			wanted: fmt.Errorf("the value of %s caused an error: %v\n", "1234w", ErrInvalidSSNNumbers),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isNumber(tc.ssn)

			if got != tc.wanted {
				if !strings.Contains(tc.wanted.Error(), got.Error()) {
					t.Errorf("\nGot    :%v\n wanted:%v", got, tc.wanted)
				}
			}
		})
	}
}

func TestIsPrefixValid(t *testing.T) {
	testCases := []struct {
		name   string
		ssn    string
		wanted error
	}{
		{
			name:   "Test-Basic",
			ssn:    "123456789",
			wanted: nil,
		},
		{
			name:   "Test SSN invalid prefix",
			ssn:    "00012",
			wanted: fmt.Errorf("the value of %s caused an error: %v\n", "00012", ErrInvalidSSNPrefix),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isPrefixValid(tc.ssn)

			if got != tc.wanted {
				if !strings.Contains(tc.wanted.Error(), got.Error()) {
					t.Errorf("\nGot    :%v\n wanted:%v", got, tc.wanted)
				}
			}
		})
	}
}

func TestIsDigitPlace(t *testing.T) {
	testCases := []struct {
		name   string
		ssn    string
		wanted error
	}{
		{
			name:   "Test-Basic",
			ssn:    "123456789",
			wanted: nil,
		},
		{
			name:   "Test SSN invalid digit placement",
			ssn:    "923656789",
			wanted: fmt.Errorf("the value of %s caused an error: %v\n", "923656789", ErrInvalidSSNDigitPlace),
		},
		{
			name:   "Test SSN valid valid didgit placement of 9",
			ssn:    "923956789",
			wanted: nil,
		},
		{
			name:   "Test SSN valid valid didgit placement of 7",
			ssn:    "923756789",
			wanted: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := validDigitPlace(tc.ssn)
			//fmt.Println("got variable: ", got.Error())
			if got != tc.wanted {
				if !strings.Contains(tc.wanted.Error(), got.Error()) {
					t.Errorf("\nGot    :%v\n wanted:%v", got, tc.wanted)
				}
			}
		})
	}
}
