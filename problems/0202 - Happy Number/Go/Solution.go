func isHappy(n int) bool {
    visited := make(map[int]struct{})
    for {
        if n == 1 {
            return true
        }
        if _, ok := visited[n]; ok {
            return false
        }
        visited[n] = struct{}{}
        tmp := 0
        for i := n; i > 0; i = i / 10 {
            j := i % 10
            tmp += j * j
        }
        n = tmp
    }
}