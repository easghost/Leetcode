// Runtime: 32 ms
// Memory Usage: 6.9 MB

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for idx, val := range nums {
		if val == target {
			return idx
		}
	}
	return -1
}