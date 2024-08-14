package piscine

func SortIntegerTable(table []int) {
	for i := range table {
		for j := range table[i+1:] {
			if table[i+j+1] < table[i] {
				table[i], table[i+j+1] = table[i+j+1], table[i]
			}
		}
	}
}
