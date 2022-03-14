package landlord

import (
	"game/card"
	"reflect"
)

type cardPatternSequenceOfTriplets struct {
	cardPatternBase
}

func (r cardPatternSequenceOfTriplets) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternSequenceOfTriplets) Valid() bool {
	ranks := LandlordCardValueRanks
	if r.Cards().Exists(func(c *card.Card) bool {
		return c.Value() == card.CardValueAce
	}) && r.Cards().Exists(func(c *card.Card) bool {
		return c.Value() == card.CardValueTwo
	}) {
		ranks = card.CardValueSortRanks
	}
	r.Cards().Sort(ranks)
	if r.Cards().Exists(func(c *card.Card) bool {
		return c.Value() == card.CardValueBigJoker || c.Value() == card.CardValueSmallJoker
	}) {
		return false
	}
	subCardsCounts := r.Cards().Counts(func(val card.CardValue, count int) bool {
		return count == 3
	})
	if len(subCardsCounts) < 2 {
		return false
	}
	if len(
		r.Cards().Counts(func(val card.CardValue, count int) bool {
			return count != 3
		}),
	) > 0 {
		return false
	}
	for i := 0; i < r.Cards().Length()-1; i++ {
		current := r.Cards()[i]
		next := r.Cards()[i+1]
		switch ranks.Rank(next) - ranks.Rank(current) {
		case 0:
		case 1:
		default:
			return false
		}
	}
	return true
}

func (r cardPatternSequenceOfTriplets) Same(s card.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternSequenceOfTriplets) Equal(s card.CardPattern) bool {
	return false
}

func (r cardPatternSequenceOfTriplets) Greeter(s card.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	return LandlordCardValueRanks.Rank(r.Cards().Last()) > LandlordCardValueRanks.Rank(s.Cards().Last())
}

func (r cardPatternSequenceOfTriplets) Lesser(s card.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternSequenceOfTriplets) String() string {
	return ""
}

func (r cardPatternSequenceOfTriplets) Factory(cards card.Cards) card.CardPattern {
	return cardPatternSequenceOfTriplets{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternSequenceOfTriplets(cards card.Cards) card.CardPattern {
	return cardPatternSequenceOfTriplets{}.Factory(cards)
}
