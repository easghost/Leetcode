// Runtime: 4 ms
// Memory Usage: 3 MB

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var mid int
	for start <= end {
		mid = (start + end + 1) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}