// Runtime: 20 ms
// Memory Usage: 5.1 MB

func isPalindrome(number int) bool {
	if number < 0 {
		return false
	}

	var remainder, temp int
	var reverse int = 0
	temp = number

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break
		}
	}

	if temp == reverse {
		return true
	} else {
		return false
	}
}