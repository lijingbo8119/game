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
	ValueSmallJoker
	ValueBigJoker
)

type Values []Value

func (r Values) Rank(v *Card) ValueRank {
	for index, val := range r {
		if v.Value() == val {
			return ValueRank(index)
		}
	}
	panic("Values Rank not found")
}
