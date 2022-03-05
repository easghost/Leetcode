// Runtime: 1 ms
// Memory Usage: 4.1 MB

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	var j int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}