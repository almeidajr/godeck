package godeck

import (
	"math/rand"
	"time"
)

type Deck []Card

// Number of cards in a deck (without jokers)
const defaultDeckSize = 52

func NewDeck() Deck {
	deck := make(Deck, 0, defaultDeckSize)

	for s := Spade; s <= Heart; s++ {
		for r := Ace; r <= King; r++ {
			deck = append(deck, Card{s, r})
		}
	}

	return deck
}

func (d Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d Deck) AddJokers(n int) Deck {
	for i := 0; i < n; i++ {
		d = append(d, Card{Suit: Joker})
	}

	return d
}
