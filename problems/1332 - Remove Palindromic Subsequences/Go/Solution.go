func removePalindromeSub(s string) int {
    if len(s) == 0 { return 0 }
    l, r := 0, len(s) - 1
    for l < r {
        if s[l] == s[r] {
            l++
            r--
        } else {
            break
        }
    }
    if l == r || l > r { return 1 }
    return 2
}