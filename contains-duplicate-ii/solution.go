package contains_duplicate_ii

import "math"

func containsNearbyDuplicate(nums []int, offset int) bool {
	if !areEntriesValid(nums, offset) {
		return false
	}

	var values = make(map[int]int)

	for i, num := range nums {
		if value, ok := values[num]; ok {
			if int(math.Abs(float64(value-i))) <= offset {
				return true
			}
		}

		values[num] = i
	}

	return false
}

func areEntriesValid(nums []int, offset int) bool {
	if !areEntryNumbersValid(nums) {
		return false
	}

	if !isEntryOffsetValid(offset) {
		return false
	}

	return true
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

func isEntryOffsetValid(offset int) bool {
	if offset < 0 || float64(offset) > math.Pow10(5) {
		return false
	}

	return true
}
