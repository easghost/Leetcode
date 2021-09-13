// Runtime: 0 ms
// Memory Usage: 2 MB

func numberOfSteps(num int) int {
	if num == 0 || num < 0 {
		return num
	}

	counter := 0
	for num > 0 {
		if num%2 != 0 {
			num = num - 1
			counter++
		}
		if num > 0 {
			num = num / 2
			counter++
		}

	}
	return counter
}