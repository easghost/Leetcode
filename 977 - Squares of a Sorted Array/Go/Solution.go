// Runtime: 32 ms
// Memory Usage: 6.9 MB

func sortedSquares(nums []int) []int {
	for idx, val := range nums {
		nums[idx] = val * val
	}
	sort.Ints(nums)
	return nums
}