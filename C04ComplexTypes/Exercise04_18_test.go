package main

import "testing"

func TestCompare04_18(t *testing.T) {
	if a, b := compare04_18(); a || !b {
		t.Fail()
	}
}
