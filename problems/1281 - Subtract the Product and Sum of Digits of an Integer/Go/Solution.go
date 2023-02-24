func subtractProductAndSum(n int) int {
	sum := 0
	prod := 1
	for n > 0 {
		sum += n % 10
		prod *= n % 10
		n = n / 10
	}

	return prod - sum
}