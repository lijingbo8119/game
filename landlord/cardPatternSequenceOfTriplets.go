package landlord

import (
	"game/poker"
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
	if r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value() == poker.CardValueAce
	}) && r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value() == poker.CardValueTwo
	}) {
		ranks = poker.CardValueSortRanks
	}
	r.Cards().Sort(ranks)
	if r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value() == poker.CardValueBigJoker || c.Value() == poker.CardValueSmallJoker
	}) {
		return false
	}
	subCardsCounts := r.Cards().Counts(func(val poker.CardValue, count int) bool {
		return count == 3
	})
	if len(subCardsCounts) < 2 {
		return false
	}
	if len(
		r.Cards().Counts(func(val poker.CardValue, count int) bool {
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

func (r cardPatternSequenceOfTriplets) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternSequenceOfTriplets) Equal(s poker.CardPattern) bool {
	return false
}

func (r cardPatternSequenceOfTriplets) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	return LandlordCardValueRanks.Rank(r.Cards().Last()) > LandlordCardValueRanks.Rank(s.Cards().Last())
}

func (r cardPatternSequenceOfTriplets) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternSequenceOfTriplets) String() string {
	return ""
}

func (r cardPatternSequenceOfTriplets) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternSequenceOfTriplets{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternSequenceOfTriplets(cards poker.Cards) poker.CardPattern {
	return cardPatternSequenceOfTriplets{}.Factory(cards)
}
