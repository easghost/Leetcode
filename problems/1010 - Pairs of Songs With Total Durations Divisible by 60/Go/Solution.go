// Runtime: 43 ms
// Memory Usage: 6.9 MB

func numPairsDivisibleBy60(time []int) int {
	tmp := make([]int, 60)
	res := 0
	for _, t := range time {
		mod := t % 60
		if mod == 0 {
			res += tmp[mod]
		} else {
			res += tmp[60-mod]
		}
		tmp[mod]++
	}

	return res
}