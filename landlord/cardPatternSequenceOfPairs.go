package landlord

import (
	"game/poker"
	"reflect"
)

type cardPatternSequenceOfPairs struct {
	cardPatternBase
}

func (r cardPatternSequenceOfPairs) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternSequenceOfPairs) Valid() bool {
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
	if r.Cards().Length() < 6 {
		return false
	}
	counts := r.Cards().Counts()
	for _, count := range counts {
		if count != 2 {
			return false
		}
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

func (r cardPatternSequenceOfPairs) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternSequenceOfPairs) Equal(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	return r.Cards().First().Value() == s.Cards().First().Value()
}

func (r cardPatternSequenceOfPairs) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	return LandlordCardValueRanks.Rank(r.Cards().Last()) > LandlordCardValueRanks.Rank(s.Cards().Last())
}

func (r cardPatternSequenceOfPairs) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternSequenceOfPairs) String() string {
	return ""
}

func (r cardPatternSequenceOfPairs) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternSequenceOfPairs{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternSequenceOfPairs(cards poker.Cards) poker.CardPattern {
	return cardPatternSequenceOfPairs{}.Factory(cards)
}
