package goExtensions

import "unicode"

const UUIDRegEx = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89AB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}"

func IsNonWhitespaceWithMinimumLength(value string, minimumLength int) bool {
	var nonWhitespace = 0
	for _, rune := range value {
		if !unicode.IsSpace(rune) {
			nonWhitespace += 1
		}

		if nonWhitespace >= minimumLength {
			return true
		}
	}

	return false
}