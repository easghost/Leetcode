func findDuplicates(nums []int) []int {
    var res = make([]int, 0)
    if len(nums) == 1 {
        return res
    }
    
    for _, v := range nums {
		v = int(math.Abs(float64(v)))
		if nums[v-1] > 0 {
			nums[v-1] = nums[v-1] * -1
		} else {
			res = append(res, v)
		}
	}
    
    return res
}