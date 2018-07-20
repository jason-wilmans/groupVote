package goExtensions

import "unicode"

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