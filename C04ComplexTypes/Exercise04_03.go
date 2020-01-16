package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	comp1, comp2, comp3 := compArrays04_03()

	fmt.Println("[10]int == [...]int{9: 0}          :", comp1)
	fmt.Println("[10]int == [10]int{1, 9: 10, 4: 3} :", comp2)
	fmt.Println("arr3                               :", comp3)
}

func compArrays04_03() (bool, bool, [10]int) {
	var arr1 [10]int
	arr2 := [...]int{9: 0}
	arr3 := [10]int{1, 9: 10, 4: 3}
	return arr1 == arr2, arr1 == arr3, arr3
}
