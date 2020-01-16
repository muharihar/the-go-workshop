package main

import "testing"

func TestGetUser(t *testing.T) {
	if _, exists := getUser04_14(""); exists {
		t.Fail()
	}
	if name, exists := getUser04_14("305"); !exists || name != "Sue" {
		t.Fail()
	}
}
