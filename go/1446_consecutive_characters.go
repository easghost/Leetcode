// Runtime: 0 ms
// Memory Usage: 2.3 MB

func maxPower(s string) int {
	count := 1
	max := 0
	prev := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == prev {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 1
		}
		prev = s[i]
	}

	if count > max {
		return count
	}
	return max
}