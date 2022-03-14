package landlord

import (
	"game/card"
	"reflect"
)

type cardPatternSequenceOfTripletsWithCards struct {
	cardPatternBase
}

func (r cardPatternSequenceOfTripletsWithCards) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternSequenceOfTripletsWithCards) Valid() bool {
	if r.Cards().Exists(func(c *card.Card) bool {
		return c.Value() == card.CardValueBigJoker || c.Value() == card.CardValueSmallJoker
	}) {
		return false
	}

	tripletsCardsCounts := r.Cards().Counts(func(val card.CardValue, count int) bool {
		return count == 3 || count == 4
	})
	ranks := LandlordCardValueRanks
	if r.Cards().Exists(func(c *card.Card) bool {
		for val := range tripletsCardsCounts {
			if c.Value() == card.CardValueAce && c.Value() == val {
				return true
			}
		}
		return false
	}) && r.Cards().Exists(func(c *card.Card) bool {
		for val := range tripletsCardsCounts {
			if c.Value() == card.CardValueTwo && c.Value() == val {
				return true
			}
		}
		return false
	}) {
		ranks = card.CardValueSortRanks
	}
	r.Cards().Sort(ranks)

	if len(tripletsCardsCounts) < 2 {
		return false
	}
	subCards := r.Cards().SubCards(func(c *card.Card) bool {
		for val := range tripletsCardsCounts {
			if val == c.Value() {
				return true
			}
		}
		return false
	})
	subCards.Sort()
	for i := 0; i < subCards.Length()-1; i++ {
		current := subCards[i]
		next := subCards[i+1]
		switch ranks.Rank(next) - ranks.Rank(current) {
		case 0:
		case 1:
		default:
			return false
		}
	}
	return len(tripletsCardsCounts)*(3+1) == r.Cards().Length()
}

func (r cardPatternSequenceOfTripletsWithCards) Same(s card.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternSequenceOfTripletsWithCards) Equal(s card.CardPattern) bool {
	return false
}

func (r cardPatternSequenceOfTripletsWithCards) Greeter(s card.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	rCard := r.Cards().Last(func(c1 *card.Card) bool {
		return r.Cards().Count(func(c2 *card.Card) bool {
			return c1.Value() == c2.Value()
		}) >= 3
	})
	sCard := s.Cards().Last(func(c1 *card.Card) bool {
		return s.Cards().Count(func(c2 *card.Card) bool {
			return c1.Value() == c2.Value()
		}) >= 3
	})
	return LandlordCardValueRanks.Rank(rCard) > LandlordCardValueRanks.Rank(sCard)
}

func (r cardPatternSequenceOfTripletsWithCards) Lesser(s card.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternSequenceOfTripletsWithCards) String() string {
	return ""
}

func (r cardPatternSequenceOfTripletsWithCards) Factory(cards card.Cards) card.CardPattern {
	return cardPatternSequenceOfTripletsWithCards{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternSequenceOfTripletsWithCards(cards card.Cards) card.CardPattern {
	return cardPatternSequenceOfTripletsWithCards{}.Factory(cards)
}
