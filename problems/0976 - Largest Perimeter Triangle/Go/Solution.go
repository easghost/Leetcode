func largestPerimeter(nums []int) int {
    sort.Ints(nums)
    for i := len(nums) - 1; i >= 2; i-- {
        if(isTriangle(nums[i], nums[i-1], nums[i-2]) == true) {
            return nums[i] + nums[i-1] + nums[i-2]
        }
    }
    return 0
}

func isTriangle(a, b, c int) bool {
    return (a + b) > c && (a + c) > b && (b + c) > a
}