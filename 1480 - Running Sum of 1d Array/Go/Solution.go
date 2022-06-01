func runningSum(nums []int) []int {
    if len(nums) < 1 {
        return nums
    }
    
    res := []int{};
    sum := 0
    for _,v := range nums {
        sum += v
        res = append(res, sum)
    }
    
    return res
}