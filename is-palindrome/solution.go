package climbing_stairs

import (
	"math"
	"unicode"
)

func isPalindrome(str string) bool {
	if !isEntryStringValid(str) {
		return false
	}

	normStr := normalizeString(str)

	for i := 0; i < len(normStr)/2; i++ {
		if normStr[i] != normStr[len(normStr)-1-i] {
			return false
		}
	}

	return true
}

func normalizeString(str string) string {
	var normStr string

	for i := 0; i < len(str); i++ {
		char := str[i]

		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') {
			if char >= 65 && char <= 90 {
				char += 32
			}

			normStr += string(char)
		}
	}

	return normStr
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
