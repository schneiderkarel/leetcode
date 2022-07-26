package climbing_stairs

import (
	"math"
	"unicode"
)

func isPalindrome(str string) bool {
	if !isEntryStringValid(str) {
		return false
	}

	normalizedStr := normalizeString(str)

	for i := 0; i < len(normalizedStr)/2; i++ {
		if normalizedStr[i] != normalizedStr[len(normalizedStr)-1-i] {
			return false
		}
	}

	return true
}

func normalizeString(str string) string {
	var normalizedStr string

	for i := 0; i < len(str); i++ {
		char := str[i]

		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') {
			if char >= 65 && char <= 90 {
				char += 32
			}

			normalizedStr += string(char)
		}
	}

	return normalizedStr
}

func isEntryStringValid(str string) bool {
	if len(str) < 1 || float64(len(str)) > (2*math.Pow10(5)) {
		return false
	}

	for _, char := range str {
		if char > unicode.MaxASCII {
			return false
		}
	}

	return true
}
