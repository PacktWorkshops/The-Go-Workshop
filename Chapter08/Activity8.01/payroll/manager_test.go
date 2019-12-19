package payroll

import (
	"fmt"
	"testing"
)

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
