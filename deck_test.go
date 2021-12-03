package godeck

import "testing"

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
