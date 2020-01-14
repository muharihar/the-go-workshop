package main

import "testing"

func TestGetTypeName(t *testing.T) {
	if getTypeName(1) != "int" {
		t.Fail()
	}

	if getTypeName(1.123) != "float" {
		t.Fail()
	}

	if getTypeName("string") != "string" {
		t.Fail()
	}

	if getTypeName(true) != "bool" {
		t.Fail()
	}

	if getTypeName(struct{}{}) != "unknown" {
		t.Fail()
	}
}
