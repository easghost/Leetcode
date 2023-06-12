func summaryRanges(nums []int) []string {
    var ans []string
    
	for i:=0;i<len(nums);i++ {
        j:=i
        for j<len(nums)-1 && nums[j] == nums[j+1]-1 {
            j++
        }
        
        if i==j {
            ans = append(ans, fmt.Sprint(nums[i]))
        } else {
            ans = append(ans, fmt.Sprintf("%d->%d", nums[i], nums[j]))
            i = j
        }
    }

    return ans
}