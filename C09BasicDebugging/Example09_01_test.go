package main

import "testing"

func TestExample09_01Random(t *testing.T) {
	r := random(1, 20)
	if (r - r) != 0 {
		t.Fail()
	}
}
