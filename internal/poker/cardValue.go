package poker

type Value uint8

type ValueRank int

const (
	ValueNone Value = iota
	ValueAce
	ValueTwo
	ValueThree
	ValueFour
	ValueFive
	ValueSix
	ValueSeven
	ValueEight
	ValueNine
	ValueTen
	ValueJack
	ValueQueen
	ValueKing
	ValueJoker
	ValueColoredJoker
)

var (
	valueMap = map[Value]string{
		ValueNone:         "",
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
		ValueJoker:        "Joker",
		ValueColoredJoker: "ColoredJoker",
	}
)

type Values []Value

func (r Values) Rank(v *Card) ValueRank {
	for index, val := range r {
		if v.Value == val {
			return ValueRank(index)
		}
	}
	panic("Values Rank not found")
}
