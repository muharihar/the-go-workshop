package strings

import (
	"strings"
	"sync"
)

// https://golang.org/src/strings/replace.go
type Replacer struct {
	once   sync.Once // guards buildOnce method
	r      strings.Replacer
	oldnew []string
}
