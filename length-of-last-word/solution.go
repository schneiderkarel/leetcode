package length_of_last_word

import (
	"math"
)

func lengthOfLastWord(str string) int {
	if !isEntryStringValid(str) {
		return 0
	}

	var lastWordLen int

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			if i == len(str)-1 {
				continue
			}

			if str[i+1] != ' ' && lastWordLen > 0 {
				return lastWordLen
			}

			continue
		}

		lastWordLen++
	}

	return lastWordLen
}

func isEntryStringValid(str string) bool {
	if len(str) < 1 || float64(len(str)) > math.Pow10(4) {
		return false
	}

	return true
}
