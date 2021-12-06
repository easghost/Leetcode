// Runtime: 4 ms
// Memory Usage: 2.8 MB

func targetIndices(nums []int, target int) []int {
	if len(nums) == 0 {
		return nums
	}

	sort.Ints(nums)
	var res []int
	for i, v := range nums {
		if v == target {
			res = append(res, i)
		}
	}

	return res
}