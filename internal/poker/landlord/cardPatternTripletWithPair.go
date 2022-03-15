package landlord

import (
	"game/internal/poker"
	"reflect"
)

type cardPatternTripletWithPair struct {
	cardPatternBase
}

func (r cardPatternTripletWithPair) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternTripletWithPair) Valid() bool {
	r.Cards().Sort(LandlordCardValueRanks)
	if r.Cards().Length() != 5 {
		return false
	}
	firstCardValueCount := r.Cards().Count(func(c *poker.Card) bool {
		return c.Value() == r.Cards().First().Value()
	})
	lastCardValueCount := r.Cards().Count(func(c *poker.Card) bool {
		return c.Value() == r.Cards().Last().Value()
	})
	return (firstCardValueCount == 3 || lastCardValueCount == 3) && firstCardValueCount+lastCardValueCount == 5
}

func (r cardPatternTripletWithPair) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternTripletWithPair) Equal(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	return false
}

func (r cardPatternTripletWithPair) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	rCard := r.Cards().First(func(c1 *poker.Card) bool {
		return r.Cards().Count(func(c2 *poker.Card) bool {
			return c1.Value() == c2.Value()
		}) == 3
	})
	sCard := s.Cards().First(func(c1 *poker.Card) bool {
		return s.Cards().Count(func(c2 *poker.Card) bool {
			return c1.Value() == c2.Value()
		}) == 3
	})
	return LandlordCardValueRanks.Rank(rCard) > LandlordCardValueRanks.Rank(sCard)
}

func (r cardPatternTripletWithPair) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternTripletWithPair) String() string {
	return ""
}

func (r cardPatternTripletWithPair) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternTripletWithPair{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternTripletWithPair(cards poker.Cards) poker.CardPattern {
	return cardPatternTripletWithPair{}.Factory(cards)
}
