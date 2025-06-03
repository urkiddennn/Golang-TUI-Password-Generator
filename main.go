package main

import (
	"fmt"
	"math/rand"
)

type Password struct {
	randomeInteger int
	strString      string
	addNumbers     bool
	maxString      int
}

// generate Password
func generatePassoword(maxLen int64, toggleNumber bool, strString string, maxstr int) Password {
	var randomInt int // Declare randomInt at function scope
	// assign value
	if !toggleNumber {
		fmt.Println("Turn on the add number to add Number to the password")
	} else {
		randomInt = rand.Intn(int(maxLen))
	}
	maxStringValue := maxstr
	tgNumber := toggleNumber
	randomeLetter := RandString(maxstr, strString)

	// return value as the Password Struct
	return Password{
		randomeInteger: randomInt,
		strString:      randomeLetter,
		addNumbers:     tgNumber,
		maxString:      maxStringValue,
	}
}

// Randomize String
func RandString(n int, lrBytes string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = lrBytes[rand.Intn(len(lrBytes))]
	}
	return string(b)
}

func RandInteger(maxInt int, toogle bool) int {
	if toogle {
		return rand.Intn(maxInt)
	}
	return 0 // Added return for false case
}

func main() {
	var maxl int64 = 21
	addNumbers := true
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	mxString := 10

	password := generatePassoword(maxl, addNumbers, letterBytes, mxString)

	fmt.Println("password: ", password.randomeInteger, password.strString)
}
