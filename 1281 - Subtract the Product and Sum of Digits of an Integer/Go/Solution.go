// Runtime: 5 ms
// Memory Usage: 2.1 MB

func subtractProductAndSum(n int) int {
	res, sum := 0, 0
	prod := 1
	for n > 0 {
		sum += n % 10
		prod *= n % 10
		n = n / 10
	}

	res = prod - sum
	return res
}