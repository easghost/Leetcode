// Runtime: 12 ms
// Memory Usage: 6.7 MB

func buildArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	var res []int
	for idx := range nums {
		res = append(res, nums[nums[idx]])
	}
	return res
}