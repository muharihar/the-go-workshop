package main

import (
	"fmt"
	"os"
)

type animal07_03 struct {
	name     string
	category string
}

type person07_03 struct {
	lastName string
	isMaried bool
}

type record07_03 struct {
	key       string
	valueType string
	data      interface{}
}

func main() {
	fmt.Println(os.Args[0])

	m := make(map[string]interface{})
	a := animal07_03{name: "oreo", category: "cat"}
	p := person07_03{lastName: "Doe", isMaried: false}

	m["person"] = p
	m["animal"] = a
	m["age"] = 54
	m["isMarried"] = true
	m["lastName"] = "Smith"

	rs := []record07_03{}
	for k, v := range m {
		r := newRecord(k, v)
		rs = append(rs, r)
	}

	for _, v := range rs {
		fmt.Println("Key: ", v.key)
		fmt.Println("Data: ", v.data)
		fmt.Println("Type: ", v.valueType)
		fmt.Println()
	}
}

func newRecord(key string, i interface{}) record07_03 {
	r := record07_03{}
	r.key = key
	switch v := i.(type) {
	case int:
		r.valueType = "int"
		r.data = v
		return r
	case bool:
		r.valueType = "bool"
		r.data = v
		return r
	case string:
		r.valueType = "string"
		r.data = v
		return r
	case person07_03:
		r.valueType = "person07_03"
		r.data = v
		return r
	default:
		r.valueType = "unknown"
		r.data = v
		return r
	}
}
