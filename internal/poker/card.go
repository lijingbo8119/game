package poker

import "fmt"

type Suit uint8

const (
	SuitNone Suit = iota
	SuitHeart
	SuitDiamond
	SuitClub
	SuitSpade
)

type Card struct {
	suit  Suit
	value Value
}

func (r Card) Suit() Suit {
	return r.suit
}

func (r Card) Value() Value {
	return r.value
}

func (r Card) String() string {
	suitMap := map[Suit]string{
		SuitNone:    "none",
		SuitHeart:   "♥",
		SuitDiamond: "♣",
		SuitClub:    "♦",
		SuitSpade:   "♠",
	}
	valueMap := map[Value]string{
		ValueAce:          "A",
		ValueTwo:          "2",
		ValueThree:        "3",
		ValueFour:         "4",
		ValueFive:         "5",
		ValueSix:          "6",
		ValueSeven:        "7",
		ValueEight:        "8",
		ValueNine:         "9",
		ValueTen:          "10",
		ValueJack:         "J",
		ValueQueen:        "Q",
		ValueKing:         "K",
		ValueJoker:        "joker",
		ValueColoredJoker: "JOKER",
	}
	unicodeMap := map[string]string{
		fmt.Sprintf("%d%d", SuitNone, ValueNone): "🂠",

		fmt.Sprintf("%d%d", SuitNone, ValueColoredJoker): "🃏",
		fmt.Sprintf("%d%d", SuitNone, ValueJoker):        "🃟",

		fmt.Sprintf("%d%d", SuitHeart, ValueAce):   "🂱",
		fmt.Sprintf("%d%d", SuitHeart, ValueTwo):   "🂲",
		fmt.Sprintf("%d%d", SuitHeart, ValueThree): "🂳",
		fmt.Sprintf("%d%d", SuitHeart, ValueFour):  "🂴",
		fmt.Sprintf("%d%d", SuitHeart, ValueFive):  "🂵",
		fmt.Sprintf("%d%d", SuitHeart, ValueSix):   "🂶",
		fmt.Sprintf("%d%d", SuitHeart, ValueSeven): "🂷",
		fmt.Sprintf("%d%d", SuitHeart, ValueEight): "🂸",
		fmt.Sprintf("%d%d", SuitHeart, ValueNine):  "🂹",
		fmt.Sprintf("%d%d", SuitHeart, ValueTen):   "🂺",
		fmt.Sprintf("%d%d", SuitHeart, ValueJack):  "🂻",
		fmt.Sprintf("%d%d", SuitHeart, ValueQueen): "🂽",
		fmt.Sprintf("%d%d", SuitHeart, ValueKing):  "🂾",

		fmt.Sprintf("%d%d", SuitDiamond, ValueAce):   "🃑",
		fmt.Sprintf("%d%d", SuitDiamond, ValueTwo):   "🃒",
		fmt.Sprintf("%d%d", SuitDiamond, ValueThree): "🃓",
		fmt.Sprintf("%d%d", SuitDiamond, ValueFour):  "🃔",
		fmt.Sprintf("%d%d", SuitDiamond, ValueFive):  "🃕",
		fmt.Sprintf("%d%d", SuitDiamond, ValueSix):   "🃖",
		fmt.Sprintf("%d%d", SuitDiamond, ValueSeven): "🃗",
		fmt.Sprintf("%d%d", SuitDiamond, ValueEight): "🃘",
		fmt.Sprintf("%d%d", SuitDiamond, ValueNine):  "🃙",
		fmt.Sprintf("%d%d", SuitDiamond, ValueTen):   "🃚",
		fmt.Sprintf("%d%d", SuitDiamond, ValueJack):  "🃛",
		fmt.Sprintf("%d%d", SuitDiamond, ValueQueen): "🃝",
		fmt.Sprintf("%d%d", SuitDiamond, ValueKing):  "🃞",

		fmt.Sprintf("%d%d", SuitClub, ValueAce):   "🂱",
		fmt.Sprintf("%d%d", SuitClub, ValueTwo):   "🂲",
		fmt.Sprintf("%d%d", SuitClub, ValueThree): "🂳",
		fmt.Sprintf("%d%d", SuitClub, ValueFour):  "🂴",
		fmt.Sprintf("%d%d", SuitClub, ValueFive):  "🂵",
		fmt.Sprintf("%d%d", SuitClub, ValueSix):   "🂶",
		fmt.Sprintf("%d%d", SuitClub, ValueSeven): "🂷",
		fmt.Sprintf("%d%d", SuitClub, ValueEight): "🂸",
		fmt.Sprintf("%d%d", SuitClub, ValueNine):  "🂹",
		fmt.Sprintf("%d%d", SuitClub, ValueTen):   "🂺",
		fmt.Sprintf("%d%d", SuitClub, ValueJack):  "🂻",
		fmt.Sprintf("%d%d", SuitClub, ValueQueen): "🂽",
		fmt.Sprintf("%d%d", SuitClub, ValueKing):  "🂾",

		fmt.Sprintf("%d%d", SuitSpade, ValueAce):   "🂡",
		fmt.Sprintf("%d%d", SuitSpade, ValueTwo):   "🂢",
		fmt.Sprintf("%d%d", SuitSpade, ValueThree): "🂣",
		fmt.Sprintf("%d%d", SuitSpade, ValueFour):  "🂤",
		fmt.Sprintf("%d%d", SuitSpade, ValueFive):  "🂥",
		fmt.Sprintf("%d%d", SuitSpade, ValueSix):   "🂦",
		fmt.Sprintf("%d%d", SuitSpade, ValueSeven): "🂧",
		fmt.Sprintf("%d%d", SuitSpade, ValueEight): "🂨",
		fmt.Sprintf("%d%d", SuitSpade, ValueNine):  "🂩",
		fmt.Sprintf("%d%d", SuitSpade, ValueTen):   "🂪",
		fmt.Sprintf("%d%d", SuitSpade, ValueJack):  "🂫",
		fmt.Sprintf("%d%d", SuitSpade, ValueQueen): "🂭",
		fmt.Sprintf("%d%d", SuitSpade, ValueKing):  "🂮",
	}
	display, ok := unicodeMap[fmt.Sprintf("%d%d", r.suit, r.value)]
	if ok {
		return display
	}
	return suitMap[r.suit] + valueMap[r.value]
}

func NewCard(suit Suit, value Value) *Card {
	return &Card{
		suit:  suit,
		value: value,
	}
}
