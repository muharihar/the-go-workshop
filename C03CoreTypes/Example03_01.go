package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Println(10 > 5)  // Warning:(11, 14) Expression '10 >5' is always 'true'
	fmt.Println(10 == 5) // Warning:(12, 14) Expression '10 == 5' is always 'false'
}
