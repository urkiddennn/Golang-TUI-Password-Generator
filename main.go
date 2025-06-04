package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

// Password represents the generated password and its components.
type Password struct {
	Value         string // The final combined and shuffled password
	RandomInteger int    // The random number component
	RandomString  string // The random string component
	RandomSymbols string // The random symbols component
}

// PasswordGenerator holds parameters for password generation.
type PasswordGenerator struct {
	IncludeNumbers bool
	MaxNumber      int
	LetterCharset  string
	MaxLetters     int
	SymbolCharset  string
	MaxSymbols     int
}

// NewPasswordGenerator creates a new PasswordGenerator with default or provided values.
func NewPasswordGenerator() *PasswordGenerator {
	return &PasswordGenerator{
		LetterCharset: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		SymbolCharset: "~!@#$%^&*()_+",
	}
}

// GeneratePassword generates a password based on the generator's settings.
func (pg *PasswordGenerator) GeneratePassword() Password {
	var randomInt int
	if pg.IncludeNumbers {
		randomInt = rand.Intn(pg.MaxNumber + 1) // +1 to make MaxNumber inclusive
	}

	randomLetters := pg.RandString(pg.MaxLetters, pg.LetterCharset)
	randomSymbols := pg.RandString(pg.MaxSymbols, pg.SymbolCharset) // Reusing RandString for symbols

	combinedValue := pg.combineAndShuffle(randomInt, randomLetters, randomSymbols)

	return Password{
		Value:         combinedValue,
		RandomInteger: randomInt,
		RandomString:  randomLetters,
		RandomSymbols: randomSymbols,
	}
}

// RandString generates a random string of length n from a given charset.
func (pg *PasswordGenerator) RandString(n int, charset string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// combineAndShuffle combines the components and shuffles them.
func (pg *PasswordGenerator) combineAndShuffle(randomInt int, randomString, randomChar string) string {
	var combined string
	if pg.IncludeNumbers {
		combined += strconv.Itoa(randomInt)
	}
	combined += randomString
	combined += randomChar

	// Convert to a slice of runes to handle Unicode characters correctly (though not strictly needed for current charsets)
	runes := []rune(combined)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})

	return string(runes)
}

// promptForNumbers prompts the user for number inclusion and max number.
func promptForNumbers() (int, bool) {
	pterm.DefaultBasicText.Println("Add Number to the " + pterm.LightMagenta("Password?"))

	options := []string{"yes", "no"}
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

	if selectedOption == "yes" {
		maxl, _ := pterm.DefaultInteractiveTextInput.Show("Enter Maximum Number (e.g., 100 for numbers up to 99): ")
		maxNumber, err := strconv.Atoi(maxl)
		if err != nil {
			pterm.Error.Println("Invalid value. Please enter a number. Defaulting to 0.")
			return 0, false
		}
		pterm.Info.Println("You've entered: ", maxNumber)
		return maxNumber, true
	}
	fmt.Println()
	return 0, false
}

// promptForLength prompts the user for the desired length of a component.
func promptForLength(componentName string) int {
	pterm.DefaultBasicText.Println(fmt.Sprintf("Enter the Number of %s you want to put in your ", componentName) + pterm.LightMagenta("Password?"))
	input, _ := pterm.DefaultInteractiveTextInput.Show(fmt.Sprintf("Enter Maximum %s: ", componentName))
	length, err := strconv.Atoi(input)
	if err != nil {
		pterm.Error.Println("Invalid value. Please enter a number. Defaulting to 0.")
		return 0
	}
	return length
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Title bar
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Go", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("GENPASS", pterm.FgLightMagenta.ToStyle())).
		Render()
	pterm.DefaultBasicText.Println("Create Password" + pterm.LightYellow(" Fast") + " as " + pterm.Yellow("Lightning"))

	pg := NewPasswordGenerator()

	// Prompt for number options
	pg.MaxNumber, pg.IncludeNumbers = promptForNumbers()

	// Prompt for letter options
	pg.MaxLetters = promptForLength("Strings")

	// Prompt for symbol options
	pg.MaxSymbols = promptForLength("Symbols")

	multi := pterm.DefaultMultiPrinter
	pb1, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Password Generating")

	multi.Start()
	for i := 1; i <= 100; i++ {
		pb1.Increment()
		time.Sleep(time.Millisecond * 50)
	}
	multi.Stop()

	password := pg.GeneratePassword()
	fmt.Println("Generated Password: ", password.Value)
}
