package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true
	for i := range a[1:] {
		if f(a[i], a[i+1]) > 0 {
			ascending = false
		} else if f(a[i], a[i+1]) < 0 {
			descending = false
		}
	}
	return ascending || descending
}
