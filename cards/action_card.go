package cards

import "fmt"

type ActionCard struct {
	Card
	Action string
}

func blah() {
	miniCard1 := miniCard{MiniColor: "mini-blue", MiniNumber: 1}

	fmt.Println(miniCard1)
}
