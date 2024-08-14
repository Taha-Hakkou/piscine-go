package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}
	if n%2 == 0 {
		if n%3 == 0 {
			return "rock and roll\n"
		}
		return "rock\n"
	} else if n%3 == 0 {
		return "roll\n"
	}
	return "error: non divisible\n"
}
