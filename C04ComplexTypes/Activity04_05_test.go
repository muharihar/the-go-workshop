package main

import (
	"reflect"
	"testing"
)

func TestRemoveBad04_05(t *testing.T) {
	s := []string{"Good", "Good", "Good", "Good"}
	if !reflect.DeepEqual(s, removeBad()) {
		t.Fail()
	}
}
