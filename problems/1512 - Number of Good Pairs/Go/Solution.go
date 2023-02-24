func numIdenticalPairs(nums []int) int {
    hMap := make(map[int]int)
    res := 0
    
    for _, num := range nums {
		res += hMap[num]
        hMap[num]++
    }
    
    return res
}