package climbing_stairs

func climbStairs(n int) int {
	if !validEntry(n) {
		return 0
	}

	var fbSq = []int{0, 1, 1}

	for i := 2; i <= n; i++ {
		fbSq = append(fbSq, fbSq[i-1]+fbSq[i])
	}

	return fbSq[n+1]
}

func validEntry(n int) bool {
	if n < 1 || n > 45 {
		return false
	}

	return true
}
