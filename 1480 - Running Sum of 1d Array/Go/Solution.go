// Runtime: 10 ms
// Memory Usage: 3 MB

func runningSum(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}

	sum := 0
	for idx, val := range nums {
		sum += val
		nums[idx] = sum
	}
	return nums
}