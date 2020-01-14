package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	data := GetData()
	for i := 0; i < len(data); i++ {
		fmt.Printf("%v is %v\n", data[i], getTypeName(data[i]))
	}
}

func getTypeName(v interface{}) string {
	switch v.(type) {
	case int, int32, int64:
		return "int"
	case float32, float64:
		return "float"
	case bool:
		return "bool"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

func GetData() []interface{} {
	return []interface{}{
		1,
		3.14,
		"hello",
		true,
		struct{}{},
	}
}
