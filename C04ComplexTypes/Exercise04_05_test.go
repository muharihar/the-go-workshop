package main

import "testing"

func TestMessage04_05(t *testing.T) {
	if message04_05() != "It's time to Go\n" {
		t.Fail()
	}
}
