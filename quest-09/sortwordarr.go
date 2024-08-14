package piscine

func SortWordArr(a []string) {
	for i := range a {
		for j := range a[i+1:] {
			if a[i] > a[i+j+1] {
				a[i], a[i+j+1] = a[i+j+1], a[i]
			}
		}
	}
}
