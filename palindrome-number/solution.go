package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if !validEntry(x) {
		return false
	}

	str := fmt.Sprintf("%d", x)

	for i, v := range str {
		i++

		if i > len(str)/2 {
			break
		}

		if uint8(v) != str[len(str)-i] {
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
