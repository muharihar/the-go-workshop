package main

import "testing"

func TestGetDots(t *testing.T) {
	dots := getDots()
	if len(dots) != 4 {
		t.Fail()
	}
	if dots[0].name != "" || dots[1].name != "A" || dots[2].name != "B" || dots[3].name != "C" {
		t.Fail()
	}
}
