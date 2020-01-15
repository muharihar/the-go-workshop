package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	comp1, comp2, comp3 := compArrays()

	fmt.Println("[5]int == [5]int{0}               : ", comp1)
	fmt.Println("[5]int == [...]int{0, 0, 0, 0, 0} : ", comp2)
	fmt.Println("[5]int == [5]int{0, 0, 0, 0, 9}   : ", comp3)
}

func compArrays() (bool, bool, bool) {
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0, 0, 0, 0, 0}
	arr4 := [5]int{0, 0, 0, 0, 9}
	return arr1 == arr2, arr1 == arr3, arr1 == arr4
}
