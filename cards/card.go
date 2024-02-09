package cards

// ActionCard represents a card in the one game,
// use it to create any card.
type Card struct {
	Color  string
	Number uint
}

type miniCard struct {
	MiniColor  string
	MiniNumber uint
}
