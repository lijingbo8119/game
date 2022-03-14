package poker

type CardSuit uint8

const (
	CardSuitNone    CardSuit = 0
	CardSuitHeart   CardSuit = 1
	CardSuitDiamond CardSuit = 2
	CardSuitClub    CardSuit = 3
	CardSuitSpade   CardSuit = 4
)

type Card struct {
	suit  CardSuit
	value CardValue
}

func (r Card) Suit() CardSuit {
	return r.suit
}

func (r Card) Value() CardValue {
	return r.value
}

func (r Card) String() string {
	suitMap := map[CardSuit]string{
		CardSuitNone:    "none",
		CardSuitHeart:   "heart",
		CardSuitDiamond: "diamond",
		CardSuitClub:    "club",
		CardSuitSpade:   "spade",
	}
	valueMap := map[CardValue]string{
		CardValueAce:        "A",
		CardValueTwo:        "2",
		CardValueThree:      "3",
		CardValueFour:       "4",
		CardValueFive:       "5",
		CardValueSix:        "6",
		CardValueSeven:      "7",
		CardValueEight:      "8",
		CardValueNine:       "9",
		CardValueTen:        "10",
		CardValueJack:       "J",
		CardValueQueen:      "Q",
		CardValueKing:       "K",
		CardValueSmallJoker: "joker",
		CardValueBigJoker:   "JOKER",
	}
	return suitMap[r.suit] + valueMap[r.value]
}

func NewCard(suit CardSuit, value CardValue) *Card {
	return &Card{
		suit:  suit,
		value: value,
	}
}
