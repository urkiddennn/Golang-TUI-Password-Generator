package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
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

// UI for the selection of the Numbers
func numberSelect() (int, bool) {
	var options []string

	target := "yes"
	var maxNumber int

	pterm.DefaultBasicText.Println("Add Number to the" + pterm.LightMagenta(" Password?"))

	options = append(options, fmt.Sprintf("yes"))
	options = append(options, fmt.Sprintf("no"))

	selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

	if contains(target, selectedOptions) {
		maxl, _ := pterm.DefaultInteractiveTextInput.Show("Enter Maximum Number: ")

		maxNumber = covertNumberToString(maxl)
	}
	fmt.Println()
	pterm.Info.Println("You've enter: ", maxNumber)

	return maxNumber, contains(target, selectedOptions)
}

// CHeck if it contains a certain value
func contains(target string, options string) bool {
	return target == options
}

// return maxNumber functions
func covertNumberToString(maxl string) int {
	maxNumber, err := strconv.Atoi(maxl)
	if err != nil {
		fmt.Println("Invalid Value, it should be a number")
	}

	return maxNumber
}

func letterOptions() (string, int) {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	pterm.DefaultBasicText.Println("Enter the Number of Strings you want to put in your" + pterm.LightMagenta(" Password?"))

	mString, _ := pterm.DefaultInteractiveTextInput.Show("Enter Maximum String: ")
	mxString := covertNumberToString(mString)
	return letterBytes, mxString
}

func main() {
	// tile bar
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Go", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("GENPASS", pterm.FgLightMagenta.ToStyle())).
		Render()
	// Go GENPASS details
	pterm.DefaultBasicText.Println("Create Password" + pterm.LightYellow(" Fast") + " as " + pterm.Yellow(" Lightning"))

	// call the number UI
	maxNumber, addNumbers := numberSelect()

	letterBytes, mxString := letterOptions()
	maxSymbols := 20
	symbols := "!@#$%^&*()"

	password := generatePassoword(maxNumber, addNumbers, letterBytes, mxString, maxSymbols, symbols)

	multi := pterm.DefaultMultiPrinter
	pb1, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Password Generating")

	multi.Start()

	for i := 1; i <= 100; i++ {
		pb1.Increment() // Increment the first progress bar at each iteration

		time.Sleep(time.Millisecond * 50) // Pause for 50 milliseconds at each iteration
	}
	multi.Stop()

	fmt.Println("password: ", password.combineValueResult)
}
