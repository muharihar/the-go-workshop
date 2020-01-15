package main

import "testing"

func TestCompareArrays(t *testing.T) {
	if comp1, comp2, comp3 := compArrays(); !comp1 || !comp2 || comp3 {
		t.Fail()
	}
}
