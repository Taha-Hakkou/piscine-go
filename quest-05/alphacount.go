package piscine

func AlphaCount(s string) int {
	count := 0
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			count++
		}
	}
	return count
}
