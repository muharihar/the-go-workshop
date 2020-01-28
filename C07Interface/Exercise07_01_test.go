package main

import "testing"

func TestPerson07_01_String(t *testing.T) {
	testCases := []struct {
		name      string
		age       int
		isMarried bool
		wanted    string
	}{
		{
			name:      "Cayden Smash",
			age:       4,
			isMarried: false,
			wanted:    "Cayden Smash (4 years old).\nMarried status: false ",
		},
	}

	t.Log("Detailed TestCases:")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := Person07_01{name: tc.name, age: tc.age, isMarried: tc.isMarried}
			got := p.String()
			if got != tc.wanted {
				t.Errorf("Got   : %v\n Wanted: %v\n", got, tc.wanted)
			}
		})
	}
}
