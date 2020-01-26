package main

import "testing"

func TestExercise06_03(t *testing.T) {
	pay, err := payDay(80, 40)
	if pay != 3280 || err != nil {
		t.Fail()
	}
}
