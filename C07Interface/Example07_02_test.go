package main

import (
	"fmt"
	"testing"
)

func TestExample07_02(t *testing.T) {
	c := cat{name: "oreo"}

	if fmt.Sprintf("%T", c) != "main.cat" {
		t.Fail()
	}
}
