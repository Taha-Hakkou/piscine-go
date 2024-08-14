package piscine

type food struct {
	name     string
	preptime int
}

func FoodDeliveryTime(order string) int {
	menu := []food{
		{name: "burger", preptime: 15},
		{name: "chips", preptime: 10},
		{name: "nuggets", preptime: 12},
	}
	for _, f := range menu {
		if f.name == order {
			return f.preptime
		}
	}
	return 404
}
