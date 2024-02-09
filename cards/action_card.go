package cards

import "fmt"

// ActionCard is a card that allows us to change the rules of the game.
// sometimes action cards can make a player lose.
type ActionCard struct {
	Card
	Action string
}

func blah() {
	miniCard1 := miniCard{MiniColor: "mini-blue", MiniNumber: 1}

	fmt.Println(miniCard1)
}
