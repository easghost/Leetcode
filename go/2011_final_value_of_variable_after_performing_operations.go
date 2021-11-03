// Runtime: 4 ms
// Memory Usage: 3,6 MB

func finalValueAfterOperations(operations []string) int {
	res := 0
	for _, v := range operations {
		if strings.Contains(v, "--") {
			res--
		} else {
			res++
		}
	}
}