package main

import "testing"

func TestValidateRoutingNumber(t *testing.T) {
	testCases := []struct {
		name          string
		routingNumber int
		wantedError   func(t *testing.T, err error) bool
	}{
		{
			name:          "Error occurs",
			routingNumber: 99,
			wantedError:   func(t *testing.T, err error) bool { return err == ErrInvalidRoutingNum },
		},
		{
			name:          "No Error",
			routingNumber: 101,
			wantedError:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := directDeposit{
				routingNumber: tc.routingNumber,
			}
			err := d.validateRoutingNumber()
			if err != nil {
				if tc.wantedError != nil && tc.wantedError(t, err) {
					return
				}
				t.Errorf("d.ValidateRoutingNumber: %v", err)
			}
		})
	}
}
func TestValidateLastName(t *testing.T) {
	testCases := []struct {
		name        string
		lastName    string
		wantedError func(t *testing.T, err error) bool
	}{
		{
			name:        "Error occurs with no name",
			lastName:    "",
			wantedError: func(t *testing.T, err error) bool { return err == ErrInvalidLastName },
		},
		{
			name:        "Error occurs with spaces only provided",
			lastName:    "      ",
			wantedError: func(t *testing.T, err error) bool { return err == ErrInvalidLastName },
		},

		{
			name:        "No error",
			lastName:    "smith",
			wantedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := directDeposit{
				lastName: tc.lastName,
			}
			err := d.validateLastName()
			if err != nil {
				if tc.wantedError != nil && tc.wantedError(t, err) {
					return
				}
				t.Errorf("d.ValidateRoutingNumber: %v", err)
			}
		})
	}
}
