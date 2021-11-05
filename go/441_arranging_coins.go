// Runtime: 4 ms
// Memory Usage: 2.3 MB

func arrangeCoins(n int) int {
	res := 1
	for res <= n {
		n -= res
		res++
	}

	return res - 1
}