// Runtime: 0 ms
// Memory Usage: 2.2 MB

func minCostToMoveChips(chips []int) int {
	even, odd := 0, 0
	for _, v := range chips {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if odd < even {
		return odd
	}
	return even
}