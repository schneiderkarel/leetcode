package pascals_triangle

func generate(numRows int) [][]int {
	if !isEntryNumberValid(numRows) {
		return [][]int{}
	}

	var rows = [][]int{
		{
			1,
		},
	}

	for i := 0; i < numRows-1; i++ {
		var newRow = []int{1}

		for j := 0; j < len(rows[i])-1; j++ {
			newRow = append(newRow, rows[i][j]+rows[i][j+1])
		}

		rows = append(rows, append(newRow, 1))
	}

	return rows
}

func isEntryNumberValid(numRows int) bool {
	if numRows < 1 || numRows > 30 {
		return false
	}

	return true
}
