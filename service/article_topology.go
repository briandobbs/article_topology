package service

import "unicode"

type CharacterSet struct {
	Character string
	Size      int
	FirstPos  int
	LastPos   int
}

type CharacterTypeSet struct {
	CharacterType string
	Size          int
	FirstPos      int
	LastPos       int
	SubSets       []CharacterSet
}

// Main struct that holds each category struct
type TextBodyTopology struct {
	CharacterTypeSets []CharacterTypeSet
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
func (topology *TextBodyTopology) ParseText(text []byte) {
	for i, b := range text {
		var charTypeSet *CharacterTypeSet
		character := string(b)

		switch {
		case IsAlphabetChar(b):
			// Find or create the Alphabet CharacterTypeSet
			charTypeSet = findOrCreateCharacterTypeSet(topology, "Alphabet")
		case IsDigit(b):
			// Find or create the Digit CharacterTypeSet
			charTypeSet = findOrCreateCharacterTypeSet(topology, "Digit")
		case IsWhitespace(b):
			// Find or create the Whitespace CharacterTypeSet
			charTypeSet = findOrCreateCharacterTypeSet(topology, "Whitespace")
		case IsControlChar(b):
			// Find or create the Control CharacterTypeSet
			charTypeSet = findOrCreateCharacterTypeSet(topology, "Control")
		case IsSpecialChar(b):
			// Find or create the Special CharacterTypeSet
			charTypeSet = findOrCreateCharacterTypeSet(topology, "Special")
		}

		if charTypeSet != nil {
			// Update overall CharacterTypeSet information
			if charTypeSet.Size == 0 {
				charTypeSet.FirstPos = i
			}
			charTypeSet.Size++
			charTypeSet.LastPos = i

			// Update or create the CharacterSet for the specific character
			updateOrCreateCharacterSet(charTypeSet, character, i)
		}
	}
}

// Helper function to find or create a CharacterTypeSet by type name
func findOrCreateCharacterTypeSet(topology *TextBodyTopology, characterType string) *CharacterTypeSet {
	for i := range topology.CharacterTypeSets {
		if topology.CharacterTypeSets[i].CharacterType == characterType {
			return &topology.CharacterTypeSets[i]
		}
	}

	// If not found, create a new CharacterTypeSet
	newTypeSet := CharacterTypeSet{
		CharacterType: characterType,
		FirstPos:      -1,
		LastPos:       -1,
		SubSets:       []CharacterSet{},
	}
	topology.CharacterTypeSets = append(topology.CharacterTypeSets, newTypeSet)
	return &topology.CharacterTypeSets[len(topology.CharacterTypeSets)-1]
}

// Helper function to update or create a CharacterSet within a CharacterTypeSet
func updateOrCreateCharacterSet(charTypeSet *CharacterTypeSet, character string, position int) {
	for i := range charTypeSet.SubSets {
		if charTypeSet.SubSets[i].Character == character {
			// Update existing CharacterSet for this character
			if charTypeSet.SubSets[i].Size == 0 {
				charTypeSet.SubSets[i].FirstPos = position
			}
			charTypeSet.SubSets[i].Size++
			charTypeSet.SubSets[i].LastPos = position
			return
		}
	}

	// If character not found in SubSets, create a new CharacterSet
	newCharSet := CharacterSet{
		Character: character,
		Size:      1,
		FirstPos:  position,
		LastPos:   position,
	}
	charTypeSet.SubSets = append(charTypeSet.SubSets, newCharSet)
}
