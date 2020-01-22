package main

import "testing"

func TestCsvHeaderCol05_04(t *testing.T) {
	testCases := []struct {
		name   string
		hdr    []string
		wanted string
	}{
		{
			name:   "employee test",
			hdr:    []string{"employee", "empid"},
			wanted: "employee",
		},

		{
			name:   "hours worked test",
			hdr:    []string{"employee", "hours worked"},
			wanted: "hours worked",
		},

		{
			name:   "hourly rate test",
			hdr:    []string{"employee", "hourly rate"},
			wanted: "hourly rate",
		},
	}

	t.Log("Detail TestCases:")
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := csvHdrCol05_04(tc.hdr)
			if got == nil {
				t.Errorf("got was nil %v", got)
			}
			for idx := range tc.hdr {
				if v, ok := got[idx]; ok {
					if v != tc.wanted {
						t.Errorf("got %v, wanted %v", got, tc.wanted)
					}

				}
			}
		})
	}
}
