package main

import "testing"

func TestSalary(t *testing.T) {
	testCases := []struct {
		name string
		x int
		y int
		f func(int,int) int
		wanted int
	}{
		{
			name: "dev salary",
			x:50,
			y:40,
			f: developerSalary,
			wanted:2000,
		},
	}

	for _, tc :=range testCases {
	 t.Run(tc.name, func(t *testing.T){
		 got := salary(tc.x,tc.y,tc.f)
		 if got != tc.wanted {
			t.Errorf("Got %v wanted %v", got, tc.wanted)
		 }
	 })
	}
}

func TestDeveloperSalary(t *testing.T){
	testCases := []struct {
		name string
		hourlyRate int
		hoursWorked int
		wanted int
	}{
		{
			name: "Normal work week",
			hourlyRate: 50,
			hoursWorked:40,
			wanted: 2000,
		},
		{
			name: "60 hours",
			hourlyRate: 50,
			hoursWorked:60,
			wanted: 3000,
		},
	}


	for _,tc := range testCases {
	t.Run(tc.name, func(t *testing.T){
		got :=developerSalary(tc.hourlyRate,tc.hoursWorked)
		if got != tc.wanted {
			t.Errorf("Got %v, wanted %v",got,tc.wanted)

		}
	})
	}
}

func TestManagerSalary(t *testing.T){
	testCases := []struct {
		name string
		salary int
		bonus int
		wanted int
	}{
		{
			name: "Junior Manager",
			salary: 100_000,
			bonus: 10_000,
			wanted: 110_000,
		},
		{
			name: "Executive Manager",
			salary: 250_000,
			bonus: 50_000,
			wanted: 300_000,
		},

	}

	for _,tc := range testCases {
	t.Run(tc.name, func(t *testing.T){
		got := managerSalary(tc.salary,tc.bonus)
		if got != tc.wanted {
			t.Errorf("Got %v, wanted %v",got,tc.wanted)

		}
	})
	}
}

