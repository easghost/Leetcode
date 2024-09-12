func countConsistentStrings(allowed string, words []string) int {
	count, wordSet := len(words), make([]int, 26)

	for i := 0; i < len(allowed); i++ {
		wordSet[allowed[i]-'a']++
	}

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if wordSet[word[i]-'a'] == 0 {
				count--
				break
			}
		}
	}

	return count
}