package main

import "testing"

func TestHoursWorked05_02(t *testing.T) {
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

	t.Log("Detail testCases:")
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			d := Developer05_02{}
			d.WorkWeek = tc.workWeek
			got := d.HoursWorked()
			if tc.wanted != got {
				t.Errorf("Hours worked not matching got: %v wanted: %v", got, tc.wanted)
			}
		})
	}
}

func TestLogHours05_02(t *testing.T) {
	testCases := []struct {
		name  string
		day   Weekday05_02
		hours int
	}{
		{
			name:  "Sunday worked",
			day:   Sunday05_02,
			hours: 12,
		},

		{
			name:  "Wednesday worked",
			day:   Wednesday05_02,
			hours: 2,
		},
		{
			name:  "Friday worked",
			day:   Friday05_02,
			hours: 8,
		},
	}

	t.Log("Detail testCases:")
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			d := Developer05_02{}
			d.LogHours(tc.day, tc.hours)

			if d.WorkWeek[tc.day] != tc.hours {
				t.Errorf("Hours for %v, got %v, expected %v ", tc.day, d.WorkWeek[tc.day], tc.hours)
			}
		})
	}

}

func TestPayDay05_02(t *testing.T) {
	testCases := []struct {
		name             string
		workWeek         [7]int
		hourlyRate       int
		wantedIsOverTime bool
		wantedPay        int
	}{
		{
			name:             "40 hour work week",
			workWeek:         [7]int{0, 8, 8, 8, 8, 8, 0},
			hourlyRate:       10,
			wantedIsOverTime: false,
			wantedPay:        400,
		},
		{
			name:             "Overtime week",
			workWeek:         [7]int{10, 10, 10, 10, 10, 10, 10},
			hourlyRate:       10,
			wantedIsOverTime: true,
			wantedPay:        760,
		},
		{
			name:             "20 hour work week",
			workWeek:         [7]int{0, 8, 8, 4, 0, 0, 0},
			hourlyRate:       10,
			wantedIsOverTime: false,
			wantedPay:        200,
		},
	}

	t.Log("Detail testCases:")
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			d := Developer05_02{}
			d.WorkWeek = tc.workWeek
			d.HourlyRate = tc.hourlyRate

			gotPay, gotIsOverTime := d.PayDay()
			if tc.wantedPay != gotPay || tc.wantedIsOverTime != gotIsOverTime {
				t.Errorf("Pay got %v, expected %v.  Is over time got %v, expected %v", gotPay, tc.wantedPay, gotIsOverTime, tc.wantedIsOverTime)
			}
		})
	}
}
