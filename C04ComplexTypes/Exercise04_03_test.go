package main

import "testing"

func TestCompArrays04_03(t *testing.T) {
	if comp1, comp2, _ := compArrays04_03(); !comp1 || comp2 {
		t.Fail()
	}
}
