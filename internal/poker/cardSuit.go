package poker

type Suit uint8

const (
	SuitNone Suit = iota
	SuitHeart
	SuitDiamond
	SuitClub
	SuitSpade
)

var (
	suitMap = map[Suit]string{
		SuitNone:    "",
		SuitHeart:   "Heart",
		SuitDiamond: "Diamond",
		SuitClub:    "Club",
		SuitSpade:   "Spade",
	}
)
