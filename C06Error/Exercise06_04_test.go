package main

import "testing"

func TestExercise06_04(t *testing.T) {
	if pay := payDay06_04(80, 40); pay != 3200 {
		t.Log(pay)
		t.Fail()
	}
}
