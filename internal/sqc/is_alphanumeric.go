package sqc

import "unicode"

func IsAlphanumeric(seq string) bool {
	for _, char := range seq {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return len(seq) > 0
}
