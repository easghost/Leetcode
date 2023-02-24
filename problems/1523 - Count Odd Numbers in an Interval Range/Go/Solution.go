func countOdds(low int, high int) int {
	diff := high - low
	if high%2 != 0 || low%2 != 0 {
		return diff/2 + 1
	}
	return diff / 2
}