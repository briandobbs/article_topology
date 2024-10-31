package service

import "unicode"

// Define structs for each category with additional properties
type AlphabetCategory struct {
	Count    int
	FirstPos int
	LastPos  int
}

type DigitCategory struct {
	Count    int
	FirstPos int
	LastPos  int
}

type WhitespaceCategory struct {
	Count    int
	FirstPos int
	LastPos  int
}

type ControlCategory struct {
	Count    int
	FirstPos int
	LastPos  int
}

type SpecialCategory struct {
	Count    int
	FirstPos int
	LastPos  int
}

// Main struct that holds each category struct
type CharCategoryCounter struct {
	Alphabet   AlphabetCategory
	Digit      DigitCategory
	Whitespace WhitespaceCategory
	Control    ControlCategory
	Special    SpecialCategory
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

// Method to parse a text body and categorize each character, tracking first and last positions
func (counter *CharCategoryCounter) ParseText(text []byte) {
	for i, b := range text {
		switch {
		case IsAlphabetChar(b):
			if counter.Alphabet.Count == 0 {
				counter.Alphabet.FirstPos = i
			}
			counter.Alphabet.Count++
			counter.Alphabet.LastPos = i

		case IsDigit(b):
			if counter.Digit.Count == 0 {
				counter.Digit.FirstPos = i
			}
			counter.Digit.Count++
			counter.Digit.LastPos = i

		case IsWhitespace(b):
			if counter.Whitespace.Count == 0 {
				counter.Whitespace.FirstPos = i
			}
			counter.Whitespace.Count++
			counter.Whitespace.LastPos = i

		case IsControlChar(b):
			if counter.Control.Count == 0 {
				counter.Control.FirstPos = i
			}
			counter.Control.Count++
			counter.Control.LastPos = i

		case IsSpecialChar(b):
			if counter.Special.Count == 0 {
				counter.Special.FirstPos = i
			}
			counter.Special.Count++
			counter.Special.LastPos = i
		}
	}
}
