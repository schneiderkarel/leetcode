package longest_substring_without_repeating_characters

import (
	"math"
	"unicode"
)

func lengthOfLongestSubstring(str string) int {
	if !isEntryStringValid(str) {
		return 0
	}

	var (
		usedChars = make(map[string]struct{})
		res       int
	)

	for leftCharIndex, leftChar := range str {
		var newCount int

		if _, ex := usedChars[string(leftChar)]; !ex {
			usedChars[string(leftChar)] = struct{}{}
			newCount++
		}

		for _, rightChar := range str[leftCharIndex+1:] {
			if _, ex := usedChars[string(rightChar)]; !ex {
				usedChars[string(rightChar)] = struct{}{}
				newCount++
				continue
			}

			break
		}

		usedChars = make(map[string]struct{})
		if newCount > res {
			res = newCount
		}
	}

	return res
}

func isEntryStringValid(str string) bool {
	if len(str) < 1 || float64(len(str)) > math.Pow(5*10, 4) {
		return false
	}

	for _, char := range str {
		if char > unicode.MaxASCII {
			return false
		}
	}

	return true
}
