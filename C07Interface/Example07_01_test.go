package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestLoadPerson1(t *testing.T) {
	s := `{"Name": "Joe", "Age": 18}`
	pWanted := Person{
		Name: "Joe",
		Age:  18,
	}

	p, err := loadPerson(strings.NewReader(s))
	if !reflect.DeepEqual(p, pWanted) || err != nil {
		t.Fail()
	}
}

func TestLoadPerson2(t *testing.T) {
	s := `{"Name": "Jane", "Age": 21}`
	pWanted := Person{
		Name: "Jane",
		Age:  21,
	}

	p, err := loadPerson2(s)
	if !reflect.DeepEqual(p, pWanted) || err != nil {
		t.Fail()
	}
}
