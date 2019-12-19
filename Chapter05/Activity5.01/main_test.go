package main

import "testing"

func TestHoursWorked(t *testing.T) {
	testCases := []struct {
		name     string
		workWeek [7]int
		wanted   int
	}{
		{
			name:     "40 hour work week",
			workWeek: [7]int{0, 8, 8, 8, 8, 8, 0},
			wanted:   40,
		},

		{
			name:     "Weekend only work",
			workWeek: [7]int{8, 0, 0, 0, 0, 0, 8},
			wanted:   16,
		},

		{
			name:     "Work every day",
			workWeek: [7]int{10, 10, 10, 10, 10, 10, 10},
			wanted:   70,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			d := Developer{}
			d.WorkWeek = tc.workWeek
			got := d.HoursWorked()
			if tc.wanted != got {
				t.Errorf("Hours worked not matching got: %v wanted: %v", got, tc.wanted)
			}
		})

	}
}

func TestLogHours(t *testing.T) {
	testCases := []struct {
		name  string
		day   Weekday
		hours int
	}{
		{
			name:  "Sunday worked",
			day:   Sunday,
			hours: 12,
		},

		{
			name:  "Wednesday worked",
			day:   Wednesday,
			hours: 2,
		},
		{
			name:  "Friday worked",
			day:   Friday,
			hours: 8,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			d := Developer{}
			d.LogHours(tc.day, tc.hours)

			if d.WorkWeek[tc.day] != tc.hours {
				t.Errorf("Hours for %v, got %v, expected %v ", tc.day, d.WorkWeek[tc.day], tc.hours)
			}
		})
	}

}
