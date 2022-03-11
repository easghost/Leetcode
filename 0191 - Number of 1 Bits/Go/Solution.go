func hammingWeight(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		res += int((num >> 31) & 1)
		num <<= 1
	}
	return res
}