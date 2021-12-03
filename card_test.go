package godeck

import "testing"

func TestCardString(t *testing.T) {
	cards := []Card{
		{Suit: Spade, Rank: Ace},
		{Suit: Diamond, Rank: Five},
		{Suit: Club, Rank: Nine},
		{Suit: Heart, Rank: King},
		{Suit: Joker},
	}

	want := []string{
		"Ace of Spades",
		"Five of Diamonds",
		"Nine of Clubs",
		"King of Hearts",
		"Joker",
	}

	for i, card := range cards {
		if card.String() != want[i] {
			t.Errorf("Expected %s, but got %s", want[i], card)
		}
	}
}

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != defaultDeckSize {
		t.Errorf("Expected 52 cards, but got %d", len(d))
	}
}
