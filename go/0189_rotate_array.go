// Runtime: 32 ms
// Memory Usage: 8,3 MB

func rotate(nums []int, k int) {
	k = k % len(nums)
	temp := append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
	copy(nums, temp)
}