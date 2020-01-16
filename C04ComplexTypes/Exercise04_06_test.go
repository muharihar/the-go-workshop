package main

import "testing"

func TestMessage04_06(t *testing.T) {
	if message04_06() != `0: 1
1: 4
2: 9
3: 16
` {
		t.Fail()
	}
}
