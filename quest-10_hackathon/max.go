package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	max := a[0]
	for _, n := range a {
		if n > max {
			max = n
		}
	}
	return max
}
