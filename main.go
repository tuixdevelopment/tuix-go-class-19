package main

import (
	"fmt"
	"github.com/tuixdevelopment/tuix-go-class-19/cards"
	"github.com/tuixdevelopment/tuix-go-class-19/players"
)

func main() {
	fmt.Println("Hi all!")

	card := cards.Card{Color: "blue", Number: 1}

	actionCard := cards.ActionCard{}
	actionCard.Color = "blue"
	actionCard.Number = 2
	actionCard.Action = "skip_next_player"

	player1 := players.Player{Name: "juan"}

	fmt.Println(card, actionCard, player1)
}
