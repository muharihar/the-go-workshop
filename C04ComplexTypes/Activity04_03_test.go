package main

import "testing"

func TestRemoveBad(t *testing.T) {
	if exists := localeExists(locale{language: "", territory: ""}); exists {
		t.Fail()
	}

	if exists := localeExists(locale{language: "en", territory: "US"}); !exists {
		t.Fail()
	}
}
