package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Password struct {
	randomeInteger     int
	strString          string
	addNumbers         bool
	maxString          int
	maxSymbols         int
	symbols            string
	combineValueResult string
}

// generate Password
func generatePassoword(maxLen int, toggleNumber bool, strString string, maxstr int, maxsymbols int, symbol string) Password {
	// Declare randomInt at function scope
	// assign value

	// call the Randomize Integer
	randomInt := RandInteger(maxLen, toggleNumber)

	// Call the RandString
	randomeLetter := RandString(maxstr, strString)

	// call the Randomize Symbols
	randomSymbol := RandSymbols(maxsymbols, symbol)

	combineValue := combineAndRandom(randomInt, randomeLetter, randomSymbol)
	// return value as the Password Struct
	return Password{
		randomeInteger:     randomInt,
		strString:          randomeLetter,
		symbols:            randomSymbol,
		combineValueResult: combineValue,
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

// Randomize Number in password
func RandInteger(maxInt int, toogle bool) int {
	if toogle {
		return rand.Intn(maxInt)
	}
	return 0 // Added return for false case
}

// Randomize Symbols
func RandSymbols(n int, symbolBytes string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = symbolBytes[rand.Intn(len(symbolBytes))]
	}

	return string(b)
}

// combine random and randomize
func combineAndRandom(randomInt int, randomString string, randomChar string) string {
	combined := strconv.Itoa(randomInt) + randomString + randomChar

	charValue := make([]byte, len(combined))
	for i := range len(combined) {
		charValue[i] = combined[rand.Intn(len(combined))]
	}

	// fmt.Println(string(charValue), "Combined lenght:", len(combined), randomInt)

	return string(charValue)
}

func main() {
	maxl := 64
	addNumbers := false
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	mxString := 10
	maxSymbols := 20
	symbols := "!@#$%^&*()"

	password := generatePassoword(maxl, addNumbers, letterBytes, mxString, maxSymbols, symbols)

	fmt.Println("password: ", password.combineValueResult)
}
