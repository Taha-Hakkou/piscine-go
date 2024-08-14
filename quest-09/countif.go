package piscine

func CountIf(f func(string) bool, tab []string) int {
	if len(tab) == 1 {
		if f(tab[0]) {
			return 1
		}
		return 0
	}
	if f(tab[0]) {
		return 1 + CountIf(f, tab[1:])
	}
	return CountIf(f, tab[1:])
}
