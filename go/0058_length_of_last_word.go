// runtime: 0ms
// memory usage 2.2 MB

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	return len(words[len(words)-1])
}