package piscine

func AtoiBase(s string, base string) int {
	nbr := 0
	if isValidBase(base) {
		for i := 0; i < len(s); i++ {
			j := 0
			for j = 0; j < len(base); j++ {
				if rune(s[i]) == rune(base[j]) {
					nbr *= len(base)
					nbr += j
					break
				}
			}
			if j == len(base) {
				return 0
			}
		}
	}
	return nbr
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		c := rune(base[i])
		if c == '+' || c == '-' {
			return false
		}
		for j := 0; j < i; j++ {
			if c == rune(base[j]) {
				return false
			}
		}
	}
	return true
}
