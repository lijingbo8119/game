package doudizhu

import (
	"game/internal/poker"
	"reflect"
)

type cardPatternTripletWithCard struct {
	cardPatternBase
}

func (r cardPatternTripletWithCard) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternTripletWithCard) Valid() bool {
	r.Cards().Sort(DoudizhuValueRanks)
	if r.Cards().Length() != 4 {
		return false
	}
	firstValueCount := r.Cards().Count(func(c *poker.Card) bool {
		return c.Value == r.Cards().First().Value
	})
	lastValueCount := r.Cards().Count(func(c *poker.Card) bool {
		return c.Value == r.Cards().Last().Value
	})
	return (firstValueCount == 3 || lastValueCount == 3) && firstValueCount+lastValueCount == 4
}

func (r cardPatternTripletWithCard) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternTripletWithCard) Equal(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	return false
}

func (r cardPatternTripletWithCard) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	rCard := r.Cards().First(func(c1 *poker.Card) bool {
		return r.Cards().Count(func(c2 *poker.Card) bool {
			return c1.Value == c2.Value
		}) == 3
	})
	sCard := s.Cards().First(func(c1 *poker.Card) bool {
		return s.Cards().Count(func(c2 *poker.Card) bool {
			return c1.Value == c2.Value
		}) == 3
	})
	return DoudizhuValueRanks.Rank(rCard) > DoudizhuValueRanks.Rank(sCard)
}

func (r cardPatternTripletWithCard) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternTripletWithCard) String() string {
	return ""
}

func (r cardPatternTripletWithCard) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternTripletWithCard{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternTripletWithCard(cards poker.Cards) poker.CardPattern {
	return cardPatternTripletWithCard{}.Factory(cards)
}
