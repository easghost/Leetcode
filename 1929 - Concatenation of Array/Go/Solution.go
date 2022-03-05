// Runtime: 8 ms
// Memory Usage: 6.4 MB

func getConcatenation(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}

	nums = append(nums, nums...)
	return nums
}