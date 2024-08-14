package piscine

func JumpOver(str string) string {
	if str == "" || len(str) < 2 {
		return "\n"
	}
	s := ""
	for i := 2; i < len(str); i += 3 {
		s += string(str[i])
	}
	return s + "\n"
}
