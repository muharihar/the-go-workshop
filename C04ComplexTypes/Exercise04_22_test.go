package main

import (
	"testing"
	"time"
)

func TestDoubler04_22(t *testing.T) {
	if res, err := doubler04_22(1); res != "2" || err != nil {
		t.Fail()
	}
	if res, err := doubler04_22(-1); res != "-2" || err != nil {
		t.Fail()
	}
	if res, err := doubler04_22(true); res != "truetrue" || err != nil {
		t.Fail()
	}
	if _, err := doubler04_22(time.Now()); err == nil {
		t.Fail()
	}
}
