package main

import "testing"

func TestEmailValidator_Validate(t *testing.T) {
	t.Log("checks if email is valid")
	{
		cases := []struct {
			email    Email
			expected bool
		}{
			{Email{"1"}, false},
			{Email{"manoj.pandey"}, false},
			{Email{"manoj.pandey@outlook.com"}, true},
			{Email{"manoj.pandey@@outlook.com"}, false},
			{Email{"manoj@13.com@"}, false},
		}
		ev := EmailValidator{}
		for _, tc := range cases {
			actual := ev.Validate(tc.email)
			if actual != tc.expected {
				t.Fatalf("Expected %t got %t for email value %s", tc.expected, actual, tc.email)
			}
		}
	}
}
