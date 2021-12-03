package godeck

type Suit uint

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]string{
	"Spade",
	"Diamond",
	"Club",
	"Heart",
	"Joker",
}

func (s Suit) String() string {
	return suits[s]
}
