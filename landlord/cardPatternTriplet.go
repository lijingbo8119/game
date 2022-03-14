package landlord

import (
	"game/card"
	"reflect"
)

type cardPatternTriplet struct {
	cardPatternBase
}

func (r cardPatternTriplet) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternTriplet) Valid() bool {
	r.Cards().Sort(LandlordCardValueRanks)
	return r.Cards().Length() == 3 && r.Cards().First().Value() == r.Cards().Last().Value()
}

func (r cardPatternTriplet) Same(s card.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternTriplet) Equal(s card.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	return r.Cards().First().Value() == s.Cards().First().Value()
}

func (r cardPatternTriplet) Greeter(s card.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	return LandlordCardValueRanks.Rank(r.Cards().First()) > LandlordCardValueRanks.Rank(s.Cards().First())
}

func (r cardPatternTriplet) Lesser(s card.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternTriplet) String() string {
	return ""
}

func (r cardPatternTriplet) Factory(cards card.Cards) card.CardPattern {
	return cardPatternTriplet{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternTriplet(cards card.Cards) card.CardPattern {
	return cardPatternTriplet{}.Factory(cards)
}
