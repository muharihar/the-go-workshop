package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println(os.Args[0])

	date := time.Now()
	day := strconv.Itoa(date.Day())

	month := strconv.Itoa(int(date.Month()))
	year := strconv.Itoa(date.Year())
	hour := strconv.Itoa(date.Hour())
	minute := strconv.Itoa(date.Minute())
	second := strconv.Itoa(date.Second())
	fmt.Println(hour + ":" + minute + ":" + second + " " + year + "/" + month + "/" + day)

}
