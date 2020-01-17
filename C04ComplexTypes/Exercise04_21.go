package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	resInt, _ := doubler(5)
	fmt.Println("5    :", resInt)

	resStr, _ := doubler("string-")
	fmt.Println("string-    :", resStr)

	if resOthers, err := doubler(math.MaxFloat32); err != nil {
		fmt.Println(math.MaxFloat32, "    :", err.Error())
	} else {
		fmt.Println(math.MaxFloat32, "    :", resOthers)
	}
}

func doubler(v interface{}) (string, error) {
	if i, ok := v.(int); ok {
		return fmt.Sprint(i * 2), nil
	}

	if s, ok := v.(string); ok {
		return s + s, nil
	}

	return "", errors.New("unsupported type passed")
}
