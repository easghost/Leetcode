// Runtime: 32 ms
// Memory Usage: 8.3 MB

func moveZeroes(nums []int) {
	var zeroes int = 0
	if len(nums) > 1 {
		for idx, val := range nums {
			if val == 0 {
				if idx != len(nums)-1 {
					nums = append(nums[:idx], nums[idx+1:]...)
				}
				zeroes++
			}
		}
		for i := zeroes; i > 0; i-- {
			nums = append(nums, 0)
		}
	}
}