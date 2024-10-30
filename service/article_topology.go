package service

import "unicode"

// Define a struct to hold character category counts
type CharCategoryCounter struct {
	AlphabetCount   int
	DigitCount      int
	WhitespaceCount int
	ControlCount    int
	SpecialCount    int
}

// Function to check if a byte is an alphabet character (letter)
func IsAlphabetChar(b byte) bool {
	return unicode.IsLetter(rune(b))
}

// Function to check if a byte is a digit
func IsDigit(b byte) bool {
	return unicode.IsDigit(rune(b))
}

// Function to check if a byte is a whitespace
func IsWhitespace(b byte) bool {
	return unicode.IsSpace(rune(b)) && !IsControlChar(b)
}

// Function to check if a byte is a control character
func IsControlChar(b byte) bool {
	return unicode.IsControl(rune(b))
}

// Function to check if a byte is a special character
func IsSpecialChar(b byte) bool {
	return !IsAlphabetChar(b) && !IsDigit(b) && !IsWhitespace(b)
}

// Method to parse a text body and categorize each character
func (counter *CharCategoryCounter) ParseText(text []byte) {
	for _, b := range text {
		switch {
		case IsAlphabetChar(b):
			counter.AlphabetCount++
		case IsDigit(b):
			counter.DigitCount++
		case IsWhitespace(b):
			counter.WhitespaceCount++
		case IsControlChar(b):
			counter.ControlCount++
		case IsSpecialChar(b):
			counter.SpecialCount++
		}
	}
}
