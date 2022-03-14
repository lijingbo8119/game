package landlord

import (
	"fmt"
	"game/card"
	"reflect"
)

type cardPatternTripletWithCard struct {
	cardPatternBase
}

func (r cardPatternTripletWithCard) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternTripletWithCard) Valid() bool {
	r.Cards().Sort(LandlordCardValueRanks)
	if r.Cards().Length() != 4 {
		return false
	}
	firstCardValueCount := r.Cards().Count(func(c *card.Card) bool {
		return c.Value() == r.Cards().First().Value()
	})
	lastCardValueCount := r.Cards().Count(func(c *card.Card) bool {
		return c.Value() == r.Cards().Last().Value()
	})
	return (firstCardValueCount == 3 || lastCardValueCount == 3) && firstCardValueCount+lastCardValueCount == 4
}

func (r cardPatternTripletWithCard) Same(s card.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternTripletWithCard) Equal(s card.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	return false
}

func (r cardPatternTripletWithCard) Greeter(s card.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	rCard := r.Cards().First(func(c1 *card.Card) bool {
		return r.Cards().Count(func(c2 *card.Card) bool {
			return c1.Value() == c2.Value()
		}) == 3
	})
	sCard := s.Cards().First(func(c1 *card.Card) bool {
		return s.Cards().Count(func(c2 *card.Card) bool {
			return c1.Value() == c2.Value()
		}) == 3
	})
	fmt.Println(rCard, sCard)
	return LandlordCardValueRanks.Rank(rCard) > LandlordCardValueRanks.Rank(sCard)
}

func (r cardPatternTripletWithCard) Lesser(s card.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternTripletWithCard) String() string {
	return ""
}

func (r cardPatternTripletWithCard) Factory(cards card.Cards) card.CardPattern {
	return cardPatternTripletWithCard{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternTripletWithCard(cards card.Cards) card.CardPattern {
	return cardPatternTripletWithCard{}.Factory(cards)
}
