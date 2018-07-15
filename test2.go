package main

import (
	"fmt"
)

func main() {

	// bestP := bestPlayer(-10, 55, 10, 501, 501, 505, 3083)
	cards := []string{newCard(), "aces of dimonds", newCard()}
	cards = append(cards, "six of dimond")

	for a, c := range cards {
		fmt.Println(a, c)
	}

	// fmt.Println(cards)
	// _ = bestP

}

func newCard() string {
	return "ace of heart"

}
func bestPlayer(players ...int) int {
	bestS := players[0]
	for a, s := range players {
		fmt.Println(s, a)
		if s > bestS {
			bestS = s

		}

	}

	return bestS

}
