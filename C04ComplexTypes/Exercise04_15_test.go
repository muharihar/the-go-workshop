package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetUser04_15(t *testing.T) {
	deletedUser("305")
	fmt.Println(users04_15)
	if !reflect.DeepEqual(users04_15, map[string]string{"204": "Bob", "631": "Jake", "073": "Tracy"}) {
		t.Fail()
	}
}
