package landlord

import (
	"game/card"
	"reflect"
)

type cardPatternRocket struct {
	cardPatternBase
}

func (r cardPatternRocket) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternRocket) Valid() bool {
	r.Cards().Sort(LandlordCardValueRanks)
	if r.Cards().Length() != 2 {
		return false
	}
	if r.Cards().Exists(func(c *card.Card) bool {
		return c.Value() != card.CardValueSmallJoker && c.Value() != card.CardValueBigJoker
	}) {
		return false
	}
	return false
}

func (r cardPatternRocket) Same(s card.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternRocket) Equal(s card.CardPattern) bool {
	return false
}

func (r cardPatternRocket) Greeter(s card.CardPattern) bool {
	if !r.Valid() || !s.Valid() {
		return false
	}
	return true
}

func (r cardPatternRocket) Lesser(s card.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternRocket) String() string {
	return ""
}

func (r cardPatternRocket) Factory(cards card.Cards) card.CardPattern {
	return cardPatternRocket{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternRocket(cards card.Cards) card.CardPattern {
	return cardPatternRocket{}.Factory(cards)
}
