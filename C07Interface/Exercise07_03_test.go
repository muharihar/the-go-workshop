package main

import (
	"fmt"
	"testing"
)

func TestExercise07_03_NewRecord(t *testing.T) {
	testCases := []struct {
		name           string
		inputKey       string
		inputInterface interface{}
		wanted         record07_03
	}{
		{
			name:           "Int Scenario",
			inputKey:       "intValue",
			inputInterface: 100,
			wanted:         record07_03{key: "intValue", valueType: "int", data: 100},
		},
		{
			name:           "Bool Scenario",
			inputKey:       "intValue",
			inputInterface: true,
			wanted:         record07_03{key: "boolValue", valueType: "bool", data: 100},
		},
		{
			name:           "String Scenario",
			inputKey:       "stringValue",
			inputInterface: 100,
			wanted:         record07_03{key: "stringValue", valueType: "string", data: 100},
		},
	}

	t.Log("Detailed TestCases:")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := newRecord(tc.inputKey, tc.inputInterface)
			if got.key != tc.wanted.key {
				fmt.Errorf("Got: %v wanted %v", got, tc.wanted.key)
			}
		})
	}
}
