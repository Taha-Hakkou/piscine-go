package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := 1; i <= 4; i++ {
		cards := deck[(i-1)*3 : i*3]
		fmt.Printf("Player %d: %d, %d, %d\n", i, cards[0], cards[1], cards[2])
	}
}
