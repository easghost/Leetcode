func lengthOfLongestSubstring(s string) int {
    var idx, res int = 0, 0
    slice := make(map[byte]int)
    
    for i := 0; i < len(s); i++ {
        slice[s[i]]++
        for slice[s[i]] == 2 && idx < i {
            slice[s[idx]]--
            idx++
        }
        
        if res < i - idx + 1 {
            res = i - idx + 1
        }
    }
    
    return res
}