package best_time_to_buy_and_sell_stock

func plusOne(digits []int) []int {
	if !areEntryDigitsValid(digits) {
		return []int{}
	}

	return recursiveCallsOverDigits(len(digits)-1, digits[len(digits)-1], digits)
}

func recursiveCallsOverDigits(currentDigitIndex int, currentDigit int, digits []int) []int {
	if currentDigit == 9 {
		digits[currentDigitIndex] = 0

		if currentDigitIndex-1 >= 0 {
			digits = recursiveCallsOverDigits(currentDigitIndex-1, digits[currentDigitIndex-1], digits)
		} else {
			digits = append([]int{1}, digits...)
		}
	} else {
		digits[currentDigitIndex]++
	}

	return digits
}

func areEntryDigitsValid(digits []int) bool {
	if len(digits) < 1 || len(digits) > 100 {
		return false
	}

	for _, digit := range digits {
		if digit < 0 || digit > 9 {
			return false
		}
	}

	return true
}
