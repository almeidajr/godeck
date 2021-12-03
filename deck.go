package godeck

import (
	"math/rand"
	"time"
)

type Deck []Card

// Number of cards in a deck (without jokers)
const defaultDeckSize = 52

type DeckOption func(Deck) Deck

func NewDeck(opts ...DeckOption) Deck {
	deck := make(Deck, 0, defaultDeckSize)

	for s := Spade; s <= Heart; s++ {
		for r := Ace; r <= King; r++ {
			deck = append(deck, Card{s, r})
		}
	}

	for _, opt := range opts {
		deck = opt(deck)
	}

	return deck
}

func WithJokers(n int) DeckOption {
	return func(d Deck) Deck {
		return d.AddJokers(n)
	}
}

func WithRemove(r func(c Card) bool) DeckOption {
	return func(d Deck) Deck {
		return d.Remove(r)
	}
}

func WithMany(n int) DeckOption {
	return func(d Deck) Deck {
		ret := make(Deck, 0, n*len(d))
		for i := 0; i < n; i++ {
			ret = append(ret, d...)
		}
		return ret
	}
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

func (d Deck) Remove(r func(c Card) bool) Deck {
	var ret Deck

	for _, c := range d {
		if !r(c) {
			ret = append(ret, c)
		}
	}

	return ret
}
