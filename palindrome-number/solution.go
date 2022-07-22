package main

import (
	"math"
	"strconv"
)

func isPalindrome(x int) bool {
	if !validEntry(x) {
		return false
	}

	str := strconv.Itoa(x)

	for i, char := range str[:len(str)/2] {
		i++

		if uint8(char) != str[len(str)-i] {
			return false
		}
	}

	return true
}

func validEntry(x int) bool {
	if float64(x) < math.Pow(-2, 31) || float64(x) > math.Pow(2, 31)-1 {
		return false
	}

	return true
}
