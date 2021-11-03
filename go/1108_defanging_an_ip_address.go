// Runtime: 0 ms
// Memory Usage: 2 MB

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}