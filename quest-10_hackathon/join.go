package piscine

func Join(strs []string, sep string) string {
	s := strs[0]
	for i := 1; i < len(strs); i++ {
		s += sep + strs[i]
	}
	return s
}
