package single_number

import "math"

func singleNumber(nums []int) int {
	if !areEntryNumbersValid(nums) {
		return 0
	}

	var res int

	for _, number := range nums {
		res ^= number
	}

	return res
}

func areEntryNumbersValid(nums []int) bool {
	if len(nums) < 1 || float64(len(nums)) > 3*math.Pow10(4) {
		return false
	}

	var numsMap = make(map[int]int)

	for _, num := range nums {
		if float64(num) < 3*-math.Pow10(4) || float64(num) > 3*math.Pow10(4) {
			return false
		}

		numsMap[num]++
	}

	var oneHit bool

	for _, amountOfAppearances := range numsMap {
		if amountOfAppearances == 1 && oneHit {
			return false
		}

		if amountOfAppearances > 2 {
			return false
		}

		if amountOfAppearances == 1 {
			oneHit = true
		}
	}

	return true
}
