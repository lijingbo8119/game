package landlord

import (
	"game/poker"
	"reflect"
)

type cardPatternSequence struct {
	cardPatternBase
}

func (r cardPatternSequence) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternSequence) Valid() bool {
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
	if r.Cards().Length() < 5 {
		return false
	}
	for i := 0; i < r.Cards().Length()-1; i++ {
		current := r.Cards()[i]
		next := r.Cards()[i+1]
		if ranks.Rank(next)-ranks.Rank(current) != 1 {
			return false
		}
	}
	return true
}

func (r cardPatternSequence) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternSequence) Equal(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	return r.Cards().First().Value() == s.Cards().First().Value()
}

func (r cardPatternSequence) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	return LandlordCardValueRanks.Rank(r.Cards().Last()) > LandlordCardValueRanks.Rank(s.Cards().Last())
}

func (r cardPatternSequence) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternSequence) String() string {
	return ""
}

func (r cardPatternSequence) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternSequence{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternSequence(cards poker.Cards) poker.CardPattern {
	return cardPatternSequence{}.Factory(cards)
}
