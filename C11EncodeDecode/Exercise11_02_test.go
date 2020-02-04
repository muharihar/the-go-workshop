package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewStudent(t *testing.T) {
	str := []byte(`{"id": 1, "lname":"John", "mname": "S", "fname": "Smith", "IsMarried":false, "IsEnrolled":true}`)
	var s1 student11_02
	err := json.Unmarshal(str, &s1)
	if err != nil {
		t.Errorf("Failed: %v", err)
		t.Fail()
	}

	s := newStudent(1, "John", "S", "Smith", false, true)
	if !reflect.DeepEqual(s1, s) {
		t.Log("S1:", s1)
		t.Log("not equal with S: ", s)
		t.Fail()
	}
}
