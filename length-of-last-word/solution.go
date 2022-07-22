package main

import "math"

func lengthOfLastWord(s string) int {
	if !validEntries(s) {
		return 0
	}

	var res string

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if i == len(s)-1 {
				continue
			}

			if s[i+1] != ' ' && len(res) > 0 {
				return len(res)
			}

			continue
		}

		res += string(s[i])
	}

	return len(res)
}

func validEntries(s string) bool {
	if len(s) < 1 || float64(len(s)) > math.Pow10(4) {
		return false
	}

	return true
}
