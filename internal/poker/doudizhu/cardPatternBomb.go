package doudizhu

import (
	"game/internal/poker"
	"reflect"
)

type cardPatternBomb struct {
	cardPatternBase
}

func (r cardPatternBomb) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternBomb) Valid() bool {
	r.Cards().Sort(DoudizhuValueRanks)
	return r.Cards().Length() == 4 && r.Cards().First().Value() == r.Cards().Last().Value()
}

func (r cardPatternBomb) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternBomb) Equal(s poker.CardPattern) bool {
	return false
}

func (r cardPatternBomb) Greeter(s poker.CardPattern) bool {
	if !r.Valid() || !s.Valid() {
		return false
	}
	if r.Same(s) {
		return DoudizhuValueRanks.Rank(r.Cards().First()) > DoudizhuValueRanks.Rank(s.Cards().First())
	}
	return true
}

func (r cardPatternBomb) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternBomb) String() string {
	return ""
}

func (r cardPatternBomb) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternBomb{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternBomb(cards poker.Cards) poker.CardPattern {
	return cardPatternBomb{}.Factory(cards)
}
