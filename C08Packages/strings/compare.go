package strings

// https://golang.org/src/strings/compare.go
func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
