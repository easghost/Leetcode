// Runtime: 4 ms
// Memory Usage: 3.1 MB

func maximumWealth(accounts [][]int) int {
	res, tmp := 0, 0
	if accounts == nil {
		return res
	}
	for _, v1 := range accounts {
		for _, v2 := range v1 {
			tmp += v2
		}
		if tmp > res {
			res = tmp
		}
		tmp = 0
	}

	return res
}