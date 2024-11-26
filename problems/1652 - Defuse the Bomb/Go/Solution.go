func decrypt(code []int, k int) []int {
	n := len(code)
	result := make([]int, n)
	if k == 0 {
		return result
	}

	absK := abs(k)
	for i := 0; i < n; i++ {
		sum := 0
		for step := 1; step <= absK; step++ {
			if k > 0 {
				sum += code[(i+step)%n]
			} else {
				sum += code[(i-step+n)%n]
			}
		}
		result[i] = sum
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}