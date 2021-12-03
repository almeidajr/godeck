package godeck

import "testing"

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != defaultDeckSize {
		t.Errorf("Expected 52 cards, but got %d", len(d))
	}
}

func TestAddJokers(t *testing.T) {
	n := 2
	d := NewDeck().AddJokers(n)

	if len(d) != defaultDeckSize+n {
		t.Errorf("Expected %d cards, but got %d", defaultDeckSize+n, len(d))
	}
	for i := len(d) - n; i < len(d); i++ {
		if d[i].Suit != Joker {
			t.Errorf("Expected joker, but got %s", d[i])
		}
	}
}

func TestRemove(t *testing.T) {
	d := NewDeck()
	d = d.Remove(func(c Card) bool {
		switch c.Rank {
		case Eight, Nine, Ten:
			return true
		default:
			return false
		}
	})

	for _, c := range d {
		switch c.Rank {
		case Eight, Nine, Ten:
			t.Errorf("Expected to remove %s", c)
		}
	}
}
