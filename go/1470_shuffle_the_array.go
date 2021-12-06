// Runtime: 0 ms
// Memory Usage: 3.9 MB

func shuffle(nums []int, n int) []int {
	var res []int
	left, right := nums[0:n], nums[n:len(nums)]

	for i, v := range left {
		res = append(res, v)
		res = append(res, right[i])
	}
	return res
}