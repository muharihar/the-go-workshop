package main

import "testing"

func TestExercise06_05(t *testing.T) {
	if pay := payDay06_05(80, 40); pay != 3280 {
		t.Log(pay)
		t.Fail()
	}
}
