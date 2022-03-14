package landlord

import (
	"game/card"
	"reflect"
)

type cardPatternPair struct {
	cardPatternBase
}

func (r cardPatternPair) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternPair) Valid() bool {
	if r.Cards().Exists(func(c *card.Card) bool {
		return c.Value() == card.CardValueBigJoker || c.Value() == card.CardValueSmallJoker
	}) {
		return false
	}
	return r.Cards().Length() == 2 && r.Cards().First().Value() == r.Cards().Last().Value()
}

func (r cardPatternPair) Same(s card.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternPair) Equal(s card.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	return r.Cards().First().Value() == s.Cards().First().Value()
}

func (r cardPatternPair) Greeter(s card.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	return LandlordCardValueRanks.Rank(r.Cards().First()) > LandlordCardValueRanks.Rank(s.Cards().First())
}

func (r cardPatternPair) Lesser(s card.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternPair) String() string {
	return ""
}

func (r cardPatternPair) Factory(cards card.Cards) card.CardPattern {
	return cardPatternPair{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternPair(cards card.Cards) card.CardPattern {
	return cardPatternPair{}.Factory(cards)
}
