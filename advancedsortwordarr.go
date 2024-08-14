package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := range a {
		for j := range a[i+1:] {
			if f(a[i], a[i+j+1]) > 0 {
				a[i], a[i+j+1] = a[i+j+1], a[i]
			}
		}
	}
}
