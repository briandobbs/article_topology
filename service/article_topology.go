package service

import "unicode"

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
