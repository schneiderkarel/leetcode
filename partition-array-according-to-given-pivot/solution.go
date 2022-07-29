package partition_array_according_to_given_pivot

import "math"

func pivotArray(nums []int, pivot int) []int {
	if !areEntriesValid(nums, pivot) {
		return []int{}
	}

	var (
		smaller []int
		equal   []int
		larger  []int
	)

	for _, num := range nums {
		switch {
		case num > pivot:
			larger = append(larger, num)
			continue
		case num < pivot:
			smaller = append(smaller, num)
		default:
			equal = append(equal, num)
		}
	}

	return append(append(smaller, equal...), larger...)
}

func areEntriesValid(nums []int, pivot int) bool {
	if len(nums) < 1 || float64(len(nums)) > math.Pow10(5) {
		return false
	}

	var pivotHit bool

	for _, num := range nums {
		if pivot == num {
			pivotHit = true
		}

		if float64(num) < -math.Pow10(6) || float64(num) > math.Pow10(6) {
			return false
		}
	}

	if !pivotHit {
		return false
	}

	return true
}
