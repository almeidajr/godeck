package godeck

import (
	"math/rand"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != defaultDeckSize {
		t.Errorf("Expected 52 cards, but got %d", len(d))
	}
}

func TestWithAddJokers(t *testing.T) {
	n := 2
	d := NewDeck(WithJokers(n))

	if len(d) != defaultDeckSize+n {
		t.Errorf("Expected %d cards, but got %d", defaultDeckSize+n, len(d))
	}
	for i := len(d) - n; i < len(d); i++ {
		if d[i].Suit != Joker {
			t.Errorf("Expected joker, but got %s", d[i])
		}
	}
}

func TestWithRemove(t *testing.T) {
	d := NewDeck(
		WithRemove(func(c Card) bool {
			switch c.Rank {
			case Eight, Nine, Ten:
				return true
			default:
				return false
			}
		}),
	)

	for _, c := range d {
		switch c.Rank {
		case Eight, Nine, Ten:
			t.Errorf("Expected to remove %s", c)
		}
	}
}

func TestWithMany(t *testing.T) {
	n := 5
	d := NewDeck(WithMany(n))

	if len(d) != defaultDeckSize*n {
		t.Errorf("Expected %d cards, but got %d", defaultDeckSize*n, len(d))
	}

	m := 2
	d = NewDeck(WithMany(n), WithJokers(m))

	if len(d) != defaultDeckSize*n+m {
		t.Errorf("Expected %d cards, but got %d", defaultDeckSize*n+m, len(d))
	}
}

func TestShuffle(t *testing.T) {
	d := NewDeck()
	var prev []Card
	prev = append(prev, d...)

	shuffler = rand.New(rand.NewSource(0))
	d.Shuffle()

	mapping := []int{
		19, 36, 21, 44, 4, 45, 41, 31, 14, 51, 1, 40, 34, 24, 47, 22, 20,
		35, 30, 3, 28, 48, 25, 16, 42, 38, 18, 5, 6, 15, 46, 11, 43, 26,
		0, 27, 9, 23, 10, 33, 37, 50, 7, 39, 29, 8, 13, 17, 2, 32, 12, 49,
	}

	for i, j := range mapping {
		if d[i] != prev[j] {
			t.Errorf("Expected %s, but got %s", prev[j], d[i])
		}
	}
}
