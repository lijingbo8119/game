package landlord

import (
	"game/internal/poker"
	"reflect"
)

type cardPatternPair struct {
	cardPatternBase
}

func (r cardPatternPair) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternPair) Valid() bool {
	if r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value() == poker.ValueBigJoker || c.Value() == poker.ValueSmallJoker
	}) {
		return false
	}
	return r.Cards().Length() == 2 && r.Cards().First().Value() == r.Cards().Last().Value()
}

func (r cardPatternPair) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternPair) Equal(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	return r.Cards().First().Value() == s.Cards().First().Value()
}

func (r cardPatternPair) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	return LandlordValueRanks.Rank(r.Cards().First()) > LandlordValueRanks.Rank(s.Cards().First())
}

func (r cardPatternPair) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternPair) String() string {
	return ""
}

func (r cardPatternPair) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternPair{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternPair(cards poker.Cards) poker.CardPattern {
	return cardPatternPair{}.Factory(cards)
}
