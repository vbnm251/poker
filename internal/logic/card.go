package logic

type Card struct {
	Value int
	Suit  string
}

func GenerateDeck() []Card {
	deck := make([]Card, 52)
	i := 0

	for _, suit := range []string{Spades, Diamonds, Clubs, Hearts} {
		for val := 2; val <= 14; val++ {
			deck[i] = Card{Value: val, Suit: suit}
			i++
		}
	}

	return deck
}
