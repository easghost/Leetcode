func minBitFlips(start int, goal int) int {
	x := start ^ goal
	flipCounter := 0
	for ; x != 0; flipCounter++ {
		x &= x - 1
	}
	return flipCounter
}