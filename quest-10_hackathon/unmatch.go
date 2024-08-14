package piscine

func Unmatch(a []int) int {
	for _, n := range a {
		paired := true
		for _, m := range a {
			if n == m {
				paired = !paired
			}
		}
		if !paired {
			return n
		}
	}
	return -1
}
