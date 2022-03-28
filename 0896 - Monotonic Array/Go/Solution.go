func isMonotonic(nums []int) bool {
    inc, dec := true, true
    
    for i := 1; i <= len(nums)-1; i++ {
        if nums[i-1] < nums[i] {
            dec = false
        }
        
        if nums[i-1] > nums[i] {
            inc = false
        }
    }
    
    return inc || dec
}