func canBeIncreasing(nums []int) bool {
	isPossible := true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			if isPossible == false {
				return false
			}
			if i > 0 {
				if nums[i+1] <= nums[i-1] {
					nums[i+1] = nums[i]
				}
			}
			isPossible = false
		}
	}
	return true
}