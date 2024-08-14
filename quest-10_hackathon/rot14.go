package piscine

func Rot14(s string) string {
	str := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			if r+'O'-'A' > 'Z' {
				str += string(r + 'O' - 'Z' - rune(1))
			} else {
				str += string(r + 'O' - 'A')
			}
		} else if r >= 'a' && r <= 'z' {
			if r+'o'-'a' > 'z' {
				str += string(r + 'o' - 'z' - rune(1))
			} else {
				str += string(r + 'o' - 'a')
			}
		} else {
			str += string(r)
		}
	}
	return str
}
