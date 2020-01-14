package main

import (
	"reflect"
	"testing"
)

func TestGetWeek(t *testing.T) {
	s := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if !reflect.DeepEqual(s, getWeek()) {
		t.Fail()
	}
}
