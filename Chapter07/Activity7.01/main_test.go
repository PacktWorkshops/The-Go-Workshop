package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestPay(t *testing.T) {
	testCases := []struct {
		name                   string
		inputFirst             string
		inputLast              string
		inputHourlyRate        float64
		inputHoursWorkedInYear float64
		wantedFullName         string
		wantedPay              float64
	}{
		{
			name:                   "Test Basic Pay",
			inputFirst:             "Celina",
			inputLast:              "Smith",
			inputHourlyRate:        40.00,
			inputHoursWorkedInYear: 2000,
			wantedFullName:         "Celina Smith",
			wantedPay:              80_000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := Developer{Individual: Employee{Id: 1, FirstName: tc.inputFirst, LastName: tc.inputLast}, HourlyRate: tc.inputHourlyRate, HoursWorkedInYear: tc.inputHoursWorkedInYear, Review: nil}
			gotName, gotSalary := d.Pay()
			if gotName != tc.wantedFullName || gotSalary != tc.wantedPay {
				fmt.Errorf("Got name: %v wanted name: %v Got salary: %v wanted salary %v", gotName, tc.wantedFullName, gotSalary, tc.wantedPay)
			}
		})
	}
}

func TestFullName(t *testing.T) {
	testCases := []struct {
		name   string
		fname  string
		lname  string
		wanted string
	}{
		{
			name:   "First test",
			fname:  "Cailyn",
			lname:  "Jones",
			wanted: "Cailyn Jones",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := Developer{Individual: Employee{Id: 1, FirstName: tc.fname, LastName: tc.lname}}
			gotName := d.FullName()
			if gotName != tc.wanted {
				fmt.Errorf("Got name: %v wanted name: %v ", gotName, tc.wanted)
			}
		})
	}
}

func TestManagerPay(t *testing.T) {
	testCases := []struct {
		name                string
		inputFirst          string
		inputLast           string
		inputSalary         float64
		inputCommissionRate float64
		wantedFullName      string
		wantedPay           float64
	}{
		{
			name:                "Test Basic Pay",
			inputFirst:          "Celina",
			inputLast:           "Smith",
			inputSalary:         150000,
			inputCommissionRate: .10,
			wantedFullName:      "Cayden Jackson",
			wantedPay:           165000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := Manager{Individual: Employee{Id: 2, FirstName: tc.inputFirst, LastName: tc.inputLast}, Salary: tc.inputSalary, CommissionRate: tc.inputCommissionRate}
			gotName, gotSalary := m.Pay()
			if gotName != tc.wantedFullName || gotSalary != tc.wantedPay {
				fmt.Errorf("Got name: %v wanted name: %v Got salary: %v wanted salary %v", gotName, tc.wantedFullName, gotSalary, tc.wantedPay)
			}
		})
	}
}

func TestConvertReviewToInt(t *testing.T) {
	testCases := []struct {
		name      string
		review    string
		wantedInt int
		wantedErr error
	}{
		{
			name:      "Test-Excellent",
			review:    "Excellent",
			wantedInt: 5,
			wantedErr: nil,
		},
		{
			name:      "Test-Good",
			review:    "Good",
			wantedInt: 4,
			wantedErr: nil,
		},
		{
			name:      "Test-Fair",
			review:    "Fair",
			wantedInt: 3,
			wantedErr: nil,
		},
		{
			name:      "Test-Poor",
			review:    "Poor",
			wantedInt: 2,
			wantedErr: nil,
		},
		{
			name:      "Test-Unsatisfactory",
			review:    "Unsatisfactory",
			wantedInt: 1,
			wantedErr: nil,
		},
		{
			name:      "Test-Error",
			review:    "great",
			wantedInt: 0,
			wantedErr: errors.New("invalid rating: great"),
		},
	}

	for _, tc := range testCases {
		gotrating, goterr := convertReviewToInt(tc.review)

		if gotrating != tc.wantedInt || goterr != tc.wantedErr {
			fmt.Errorf("Got rating: %v wanted: %v Got error: %v wanted %v", gotrating, tc.wantedInt, goterr, tc.wantedErr)
		}
	}
}
