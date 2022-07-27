package number_of_one_bits

import "fmt"

func hammingWeight(num uint32) int {
	var ones int

	for _, char := range fmt.Sprintf("%b", int64(num)) {
		if char == 49 {
			ones++
		}
	}

	return ones
}
