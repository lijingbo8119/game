package landlord

import (
	"game/card"
	"reflect"
)

type cardPatternBomb struct {
	cardPatternBase
}

func (r cardPatternBomb) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternBomb) Valid() bool {
	r.Cards().Sort(LandlordCardValueRanks)
	return r.Cards().Length() == 4 && r.Cards().First().Value() == r.Cards().Last().Value()
}

func (r cardPatternBomb) Same(s card.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternBomb) Equal(s card.CardPattern) bool {
	return false
}

func (r cardPatternBomb) Greeter(s card.CardPattern) bool {
	if !r.Valid() || !s.Valid() {
		return false
	}
	if r.Same(s) {
		return LandlordCardValueRanks.Rank(r.Cards().First()) > LandlordCardValueRanks.Rank(s.Cards().First())
	}
	return true
}

func (r cardPatternBomb) Lesser(s card.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternBomb) String() string {
	return ""
}

func (r cardPatternBomb) Factory(cards card.Cards) card.CardPattern {
	return cardPatternBomb{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternBomb(cards card.Cards) card.CardPattern {
	return cardPatternBomb{}.Factory(cards)
}
