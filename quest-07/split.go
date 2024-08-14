package piscine

func Split(s, sep string) []string {
	var strings []string
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			if s[0:i] != "" {
				strings = append(strings, s[0:i])
			}
			s = s[i+len(sep):]
			i = -1
		}
	}
	if s != "" {
		strings = append(strings, s)
	}
	return strings
}
