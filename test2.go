package main

import (
	"fmt"
)

func main() {

	bestP := bestPlayer(-10, 55, 10, 501, 501, 505, 3083)

	fmt.Println(bestP)

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
