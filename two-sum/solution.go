package two_sum

import "math"

func twoSum(nums []int, targetNum int) []int {
	if !areEntriesValid(nums, targetNum) {
		return []int{}
	}

	for num1i, num1 := range nums {
		for num2i, num2 := range nums {
			if num1i == num2i {
				continue
			}

			if num1+num2 == targetNum {
				return append([]int{num1i, num2i})
			}
		}
	}

	return []int{}
}

func areEntriesValid(nums []int, targetNum int) bool {
	if !areEntryNumbersValid(nums) {
		return false
	}

	if !isEntryTargetNumberValid(targetNum) {
		return false
	}

	return true
}

func areEntryNumbersValid(nums []int) bool {
	if len(nums) < 2 || float64(len(nums)) > math.Pow10(4) {
		return false
	}

	for _, num := range nums {
		if float64(num) < -math.Pow10(9) || float64(num) > math.Pow10(9) {
			return false
		}
	}

	return true
}

func isEntryTargetNumberValid(targetNum int) bool {
	if float64(targetNum) < -math.Pow10(9) || float64(targetNum) > math.Pow10(9) {
		return false
	}

	return true
}
