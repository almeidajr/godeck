package godeck

import "fmt"

type Card struct {
	Suit
	Rank
}

// Number of cards in a deck (without jokers)
const defaultDeckSize = 52

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

func NewDeck() []Card {
	cards := make([]Card, 0, defaultDeckSize)

	for s := Spade; s <= Heart; s++ {
		for r := Ace; r <= King; r++ {
			cards = append(cards, Card{s, r})
		}
	}

	return cards
}
