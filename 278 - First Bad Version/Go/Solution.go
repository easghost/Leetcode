// Runtime: 900 ms
// Memory Usage: 2.1 MB

func firstBadVersion(n int) int {
	if n < 2 {
		return n
	}
	for i := n; i > 0; i-- {
		if isBadVersion(i) == false {
			return i + 1
		}
	}
	return 1
}