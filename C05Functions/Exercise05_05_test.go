package main

import "testing"

func TestSum05_05(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		wanted int
	}{
		{
			name:   "one int test",
			input:  []int{5},
			wanted: 5,
		},
		{
			name:   "three int test",
			input:  []int{5, 10, 15},
			wanted: 30,
		},
		{
			name:   "negative int test",
			input:  []int{5, -5, 10},
			wanted: 10,
		},
	}

	t.Log("Detail TestCases:")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sum(tc.input...)
			if got != tc.wanted {
				t.Errorf("Got %v, wanted %v", got, tc.wanted)
			}
		})
	}
}
