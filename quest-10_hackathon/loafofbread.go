package piscine

func LoafOfBread(str string) string {
	loaf := ""
	piece := ""
	for i := 0; i < len(str); i++ {
		if rune(str[i]) != ' ' {
			piece += string(str[i])
			if len(piece) == 5 {
				i++
			}
		}
		if len(piece) == 5 {
			loaf += piece + " "
			piece = ""
		}
	}
	if len(loaf) == 0 {
		return "\n"
	} else if len(loaf) >= 5 {
		loaf += piece
		if rune(loaf[len(loaf)-1]) == ' ' {
			loaf = loaf[:len(loaf)-1]
		}
		return loaf + "\n"
	} else {
		return "Invalid Output\n"
	}
}
