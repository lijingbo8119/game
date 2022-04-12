package poker

import (
	"encoding/json"
	"fmt"
)

var (
	unicodeMap = map[string]string{
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
)

type Card struct {
	Suit  Suit
	Value Value
}

func (r Card) String() string {
	display, ok := unicodeMap[fmt.Sprintf("%d%d", r.Suit, r.Value)]
	if ok {
		return display
	}
	return fmt.Sprintf("%s-%s", suitMap[r.Suit], valueMap[r.Value])
}

func (r *Card) UnmarshalJSON(data []byte) error {
	m := map[string]string{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	suit, ok := m["suit"]
	if !ok {
		return fmt.Errorf("Card UnmarshalJSON error")
	}
	for k, v := range suitMap {
		if v == suit {
			r.Suit = k
			break
		}
	}
	value, ok := m["value"]
	if !ok {
		return fmt.Errorf("Card UnmarshalJSON error")
	}
	for k, v := range valueMap {
		if v == value {
			r.Value = k
			break
		}
	}
	return nil
}

func (r Card) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		"suit":  suitMap[r.Suit],
		"value": valueMap[r.Value],
	}
	return json.Marshal(m)
}

func NewCard(suit Suit, value Value) *Card {
	return &Card{
		Suit:  suit,
		Value: value,
	}
}
