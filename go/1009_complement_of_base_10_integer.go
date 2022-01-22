// Runtime: 43 ms
// Memory Usage: 6.9 MB

func bitwiseComplement(n int) int {
	mask := 1
	for i := n; i > 1; i >>= 1 {
		mask <<= 1
		mask += 1
	}
	return n ^ mask
}