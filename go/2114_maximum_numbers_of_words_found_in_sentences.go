func MostWordsFound(sentences []string) int {
	if len(sentences) == 0 {
		return 0
	}

	var ans int

	for i := range sentences {
		sentence := sentences[i]
		res := strings.Split(sentence, " ")
		numberOfWords := len(res)
		if numberOfWords > ans {
			ans = numberOfWords
		}
	}

	return ans
}