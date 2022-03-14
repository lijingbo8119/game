package poker

type CardValue uint8

type CardValueRank int

const (
	CardValueAce        CardValue = 1
	CardValueTwo        CardValue = 2
	CardValueThree      CardValue = 3
	CardValueFour       CardValue = 4
	CardValueFive       CardValue = 5
	CardValueSix        CardValue = 6
	CardValueSeven      CardValue = 7
	CardValueEight      CardValue = 8
	CardValueNine       CardValue = 9
	CardValueTen        CardValue = 10
	CardValueJack       CardValue = 11
	CardValueQueen      CardValue = 12
	CardValueKing       CardValue = 13
	CardValueSmallJoker CardValue = 14
	CardValueBigJoker   CardValue = 15
)

type CardValues []CardValue

func (r CardValues) Rank(v *Card) CardValueRank {
	for index, val := range r {
		if v.Value() == val {
			return CardValueRank(index)
		}
	}
	panic("CardValues Rank not found")
}
