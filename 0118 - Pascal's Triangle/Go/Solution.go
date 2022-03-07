func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		arr := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || i == j {
				arr[j] = 1
			} else {
				arr[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res[i] = arr
	}

	return res
}