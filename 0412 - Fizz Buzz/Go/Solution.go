func fizzBuzz(n int) []string {
	res := make([]string, n)
	i := 1

	for i <= n {
		if i%15 == 0 {
			res[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			res[i-1] = "Fizz"
		} else if i%5 == 0 {
			res[i-1] = "Buzz"
		} else {
			res[i-1] = strconv.Itoa(i)
		}

		i += 1
	}

	return res

}