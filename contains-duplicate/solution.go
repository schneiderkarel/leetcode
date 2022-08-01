package contains_duplicate

import "math"

func containsDuplicate(nums []int) bool {
	if !areEntryNumbersValid(nums) {
		return false
	}

	var values = make(map[int]struct{})

	for _, num := range nums {
		if _, ex := values[num]; ex {
			return true
		}

		values[num] = struct{}{}
	}

	return false
}

func areEntryNumbersValid(nums []int) bool {
	if len(nums) < 1 || float64(len(nums)) > math.Pow10(5) {
		return false
	}

	for _, num := range nums {
		if float64(num) < -math.Pow10(9) || float64(num) > math.Pow10(9) {
			return false
		}
	}

	return true
}
