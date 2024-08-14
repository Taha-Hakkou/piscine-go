package piscine

func ShoppingListSort(slice []string) []string {
	if len(slice) == 1 {
		return slice
	}
	for i, item := range slice {
		if len(item) < len(slice[0]) {
			slice[0], slice[i] = slice[i], slice[0]
		}
	}
	ShoppingListSort(slice[1:])
	return slice
}
