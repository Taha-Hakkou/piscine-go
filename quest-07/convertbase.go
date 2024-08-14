package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := 0
	for _, r := range nbr {
		for i := range baseFrom {
			if r == rune(baseFrom[i]) {
				n *= len(baseFrom)
				n += i
			}
		}
	}
	converted := ""
	for n > 0 {
		converted = string(baseTo[n%len(baseTo)]) + converted
		n /= len(baseTo)
	}
	return converted
}
