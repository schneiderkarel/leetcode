package main

import "math"

func twoSum(nums []int, target int) []int {
	if !validEntries(nums, target) {
		return []int{}
	}

	for i1, num1 := range nums {
		for i2, num2 := range nums {
			if i1 == i2 {
				continue
			}

			if num1+num2 == target {
				return []int{i1, i2}
			}
		}
	}

	return []int{}
}

func validEntries(nums []int, target int) bool {
	if len(nums) < 2 || float64(len(nums)) > math.Pow(10, 4) {
		return false
	}

	for _, num := range nums {
		if float64(num) < math.Pow(-10, 9) || float64(num) > math.Pow(10, 9) {
			return false
		}
	}

	if float64(target) < math.Pow(-10, 9) || float64(target) > math.Pow(10, 9) {
		return false
	}

	return true
}
