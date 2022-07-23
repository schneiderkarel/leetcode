package valid_parentheses

import (
	"math"
)

func isValid(str string) bool {
	if !isEntryStringValid(str) {
		return false
	}

	var (
		parenthesesPairs = map[string]string{
			")": "(",
			"]": "[",
			"}": "{",
		}
		parenthesesQueue []string
	)

	for _, char := range str {
		switch string(char) {
		case "(":
			parenthesesQueue = append(parenthesesQueue, "(")
		case ")", "]", "}":
			if len(parenthesesQueue) == 0 {
				return false
			}

			if parenthesesPairs[string(char)] != parenthesesQueue[len(parenthesesQueue)-1] {
				return false
			}

			parenthesesQueue = append(parenthesesQueue[len(parenthesesQueue):], parenthesesQueue[:len(parenthesesQueue)-1]...)
		case "[":
			parenthesesQueue = append(parenthesesQueue, "[")
		case "{":
			parenthesesQueue = append(parenthesesQueue, "{")
		}
	}

	if len(parenthesesQueue) != 0 {
		return false
	}

	return true
}

func isEntryStringValid(str string) bool {
	if len(str) < 1 || float64(len(str)) > math.Pow10(4) {
		return false
	}

	for _, char := range str {
		if char != 123 && char != 125 && char != 40 && char != 41 && char != 93 && char != 91 {
			return false
		}

	}

	return true
}
