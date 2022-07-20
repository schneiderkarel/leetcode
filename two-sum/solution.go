package main

import "math"

func twoSum(nums []int, target int) []int {
	if !validEntries(nums, target) {
		return []int{}
	}

	for i, num := range nums {
		for j, restNum := range nums {
			if i == j {
				continue
			}

			if num+restNum == target {
				return []int{i, j}
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
