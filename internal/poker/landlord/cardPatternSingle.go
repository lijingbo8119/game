package landlord

import (
	"game/internal/poker"
	"reflect"
)

type cardPatternSingle struct {
	cardPatternBase
}

func (r cardPatternSingle) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternSingle) Valid() bool {
	return r.Cards().Length() == 1
}

func (r cardPatternSingle) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternSingle) Equal(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	return r.Cards().First().Value() == s.Cards().First().Value()
}

func (r cardPatternSingle) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	return LandlordCardValueRanks.Rank(r.Cards().First()) > LandlordCardValueRanks.Rank(s.Cards().First())
}

func (r cardPatternSingle) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternSingle) String() string {
	return ""
}

func (r cardPatternSingle) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternSingle{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternSingle(cards poker.Cards) poker.CardPattern {
	return cardPatternSingle{}.Factory(cards)
}
