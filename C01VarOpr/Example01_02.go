package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Taken from: https://en.wiktionary.org/wiki/Hello_World#Translations
var helloList0102 = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"سلام دنیا‎",
	"Halo, dunia",
	"Wilujeng, dunyo",
}

func main() {
	fmt.Println(os.Args[0])

	// Seed random number generator using the current time
	rand.Seed(time.Now().UnixNano())
	// Generate a random number in the range of out list
	index := rand.Intn(len(helloList0102))
	// Call a function and receive multiple return values
	msg, err := hello0102(index)
	// Handle any errors
	if err != nil {
		log.Fatal(err)
	}
	// Print our message to the console
	fmt.Println(msg)
}

func hello0102(index int) (string, error) {
	if index < 0 || index > len(helloList0102)-1 {
		// Create an error, convert the int type to a string
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList0102[index], nil
}
