package palindrome_number

import (
	"math"
	"strconv"
)

func isPalindrome(num int) bool {
	if !isEntryNumberValid(num) {
		return false
	}

	strNum := strconv.Itoa(num)

	for i, char := range strNum[:len(strNum)/2] {
		i++

		if uint8(char) != strNum[len(strNum)-i] {
			return false
		}
	}

	return true
}

func isEntryNumberValid(num int) bool {
	if float64(num) < -math.Pow(2, 31) || float64(num) > math.Pow(2, 31)-1 {
		return false
	}

	return true
}
