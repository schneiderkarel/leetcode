package main

import "math"

func lengthOfLastWord(s string) int {
	if !validEntries(s) {
		return 0
	}

	var count int

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if i == len(s)-1 {
				continue
			}

			if s[i+1] != ' ' && count > 0 {
				return count
			}

			continue
		}

		count++
	}

	return count
}

func validEntries(s string) bool {
	if len(s) < 1 || float64(len(s)) > math.Pow10(4) {
		return false
	}

	return true
}
