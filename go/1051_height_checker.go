func heightChecker(heights []int) int {
	arr := make([]int, len(heights))
	for i := range arr {
		arr[i] = heights[i]
	}
	sort.Ints(heights)
	c := 0
	for i := range arr {
		if arr[i] != heights[i] {
			c++
		}
	}
	return c
}