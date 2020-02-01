package main

import (
	"testing"
	"time"
)

func TestElapsed(t *testing.T) {
	Start := time.Now()
	time.Sleep(2 * time.Second)
	End := time.Now()
	result := elapsedTime(Start, End)
	t.Log(result)

}
