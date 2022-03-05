// Runtime: 0 ms
// Memory Usage: 1.9 MB

func climbStairs(n int) int {
	if n == 0 {
		return n
	}

	pre, res := 1, 1
	for i := 1; i < n; i++ {
		pre, res = res, pre+res
	}

	return res
}