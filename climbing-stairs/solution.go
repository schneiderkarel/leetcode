package climbing_stairs

func climbStairs(num int) int {
	if !isEntryNumberValid(num) {
		return 0
	}

	var fbSq = []int{0, 1, 1}

	for i := 2; i <= num; i++ {
		fbSq = append(fbSq, fbSq[i-1]+fbSq[i])
	}

	return fbSq[num+1]
}

func isEntryNumberValid(num int) bool {
	if num < 1 || num > 45 {
		return false
	}

	return true
}
