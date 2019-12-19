package main

import (
	"fmt"
	"testing"
)

func TestTriangleArea(t *testing.T) {

	testCases := []struct {
		name   string
		base   float64
		height float64
		wanted float64
	}{
		{
			name:   "Triangle Area Scenario 1",
			base:   10,
			height: 2,
			wanted: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := triangle{base: tc.base, height: tc.height}
			got := s.Area()
			if got != tc.wanted {
				fmt.Errorf("Got: %v wandted %v", got, tc.wanted)
			}

		})
	}
}

func TestTriangleName(t *testing.T) {
	testCases := []struct {
		name   string
		wanted string
	}{
		{
			name:   "Triangle Name Scenario 1",
			wanted: "triangle",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := triangle{}
			got := s.Name()
			if got != tc.wanted {
				fmt.Errorf("Got: %v wandted %v", got, tc.wanted)
			}

		})
	}
}

func TestRectangleArea(t *testing.T) {

	testCases := []struct {
		name   string
		length float64
		width  float64
		wanted float64
	}{
		{
			name:   "Rectangle Area Scenario 1",
			length: 10,
			width:  2,
			wanted: 20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := rectangle{length: tc.length, width: tc.width}
			got := s.Area()
			if got != tc.wanted {
				fmt.Errorf("Got: %v wandted %v", got, tc.wanted)
			}

		})
	}
}

func TestRectangleName(t *testing.T) {
	testCases := []struct {
		name   string
		wanted string
	}{
		{
			name:   "Rectangle Name Scenario 1",
			wanted: "rectangle",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := rectangle{}
			got := s.Name()
			if got != tc.wanted {
				fmt.Errorf("Got: %v wandted %v", got, tc.wanted)
			}

		})
	}
}
func TestSquareArea(t *testing.T) {

	testCases := []struct {
		name   string
		side   float64
		wanted float64
	}{
		{
			name:   "Square Area Scenario 1",
			side:   10,
			wanted: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := square{side: tc.side}
			got := s.Area()
			if got != tc.wanted {
				fmt.Errorf("Got: %v wandted %v", got, tc.wanted)
			}
		})
	}
}

func TestSquareName(t *testing.T) {
	testCases := []struct {
		name   string
		wanted string
	}{
		{
			name:   "Square Name Scenario 1",
			wanted: "square",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := square{}
			got := s.Name()
			if got != tc.wanted {
				fmt.Errorf("Got: %v wandted %v", got, tc.wanted)
			}

		})
	}
}
