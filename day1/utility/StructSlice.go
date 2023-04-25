package main

import "fmt"

type Player struct {
	name       string
	occupation string
	country    string
}

func main() {

	players := []Player{}
	players = append(players, Player{"John Doe", "gardener", "USA"})
	players = append(players, Player{"Roger Roe", "driver", "UK"})
	players = append(players, Player{"Paul Smith", "programmer", "Canada"})
	players = append(players, Player{"Lucia Mala", "teacher", "Slovakia"})
	players = append(players, Player{"Patrick Connor", "shopkeeper", "USA"})
	players = append(players, Player{"Tim Welson", "programmer", "Canada"})
	players = append(players, Player{"Tomas Smutny", "programmer", "Slovakia"})

	for _, player := range players {

		fmt.Println(player)
	}
	//filter programmers
	var programmers []Player

	for _, player := range players {

		if isProgrammer(player) {
			programmers = append(programmers, player)
		}
	}

	fmt.Println("Programmers:")
	for _, u := range programmers {

		fmt.Println(u)
	}
}
func isProgrammer(player Player) bool {

	return player.occupation == "programmer"
}
