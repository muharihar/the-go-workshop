package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Print(convert04_20())
}

func convert04_20() string {
	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14
	m := fmt.Sprintf("int8    = %v  > int64   = %v\n", i8, int64(i8))
	m += fmt.Sprintf("int     = %v  > int8    = %v\n", i, int8(i))
	m += fmt.Sprintf("int8    = %v  > float32 = %v\n", i8, float32(i8))
	m += fmt.Sprintf("float64 = %v > int     = %v\n", f64, int(f64))

	return m
}
