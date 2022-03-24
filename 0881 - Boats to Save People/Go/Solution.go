func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    numBoats, start := 0, 0
    
    for end := len(people)-1; start <= end; end-- {
        if end > 0 && people[end] + people[end-1] <= limit {
            end--
        } else if start < end && people[end] + people[start] <= limit {
            start++
        }
        
        numBoats++
    }
    
    return numBoats
}