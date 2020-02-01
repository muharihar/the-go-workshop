package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	ErrInvalidSSNNumbers    = errors.New("ssn has non-numeric digits")
	ErrInvalidSSNLength     = errors.New("ssn is not nine characters long")
	ErrInvalidSSNPrefix     = errors.New("ssn has three zeros as a prefix")
	ErrInvalidSSNDigitPlace = errors.New("ssn starts with a 9 requires 7 or 9 in the fourth place")
)

func main() {
	fmt.Println(os.Args[0])

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333"}
	log.Printf("Chenking data %#v", validateSSN)

	for idx, ssn := range validateSSN {
		log.Printf("Validate data %#v %d of %d", ssn, idx+1, len(validateSSN))

		ssn = strings.Replace(ssn, "-", "", -1)

		err := isNumber(ssn)
		if err != nil {
			log.Print(err)
		}

		err = validLength(ssn)
		if err != nil {
			log.Print(err)
		}

		err = isPrefixValid(ssn)
		if err != nil {
			log.Print(err)
		}

		err = validDigitPlace(ssn)
		if err != nil {
			log.Print(err)
		}
	}
}

func isNumber(ssn string) error {
	_, err := strconv.Atoi(ssn)
	if err != nil {
		return fmt.Errorf("the value of %s caused an error: %v\n", ssn, ErrInvalidSSNNumbers)
	}
	return nil
}

func validLength(ssn string) error {
	ssn = strings.TrimSpace(ssn)
	if len(ssn) != 9 {
		return fmt.Errorf("the value of %s caused an error: %v\n", ssn, ErrInvalidSSNLength)
	}
	return nil
}

func isPrefixValid(ssn string) error {
	if strings.HasPrefix(ssn, "000") {
		return fmt.Errorf("the value of %s caused an error: %v\n", ssn, ErrInvalidSSNPrefix)
	}
	return nil
}

func validDigitPlace(ssn string) error {
	if string(ssn[0]) == "9" && (string(ssn[3]) != "9" && string(ssn[3]) != "7") {
		return fmt.Errorf("the value of %s caused an error: %v\n", ssn, ErrInvalidSSNDigitPlace)
	}
	return nil
}
