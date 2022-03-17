func mergeAlternately(word1 string, word2 string) string {
	res := ""
    max := int(math.Max(float64(len(word1)), float64(len(word2))))

	for i := 0; i < max; i++ {
		if i < len(word1) {
			res += string(word1[i])
		}
		if i < len(word2) {
			res += string(word2[i])
		}
	}

	return res
}
