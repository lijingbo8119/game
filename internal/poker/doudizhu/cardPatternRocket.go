package doudizhu

import (
	"game/internal/poker"
	"reflect"
)

type cardPatternRocket struct {
	cardPatternBase
}

func (r cardPatternRocket) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternRocket) Valid() bool {
	r.Cards().Sort(DoudizhuValueRanks)
	if r.Cards().Length() != 2 {
		return false
	}
	if r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value != poker.ValueJoker && c.Value != poker.ValueColoredJoker
	}) {
		return false
	}
	return false
}

func (r cardPatternRocket) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternRocket) Equal(s poker.CardPattern) bool {
	return false
}

func (r cardPatternRocket) Greeter(s poker.CardPattern) bool {
	if !r.Valid() || !s.Valid() {
		return false
	}
	return true
}

func (r cardPatternRocket) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternRocket) String() string {
	return ""
}

func (r cardPatternRocket) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternRocket{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternRocket(cards poker.Cards) poker.CardPattern {
	return cardPatternRocket{}.Factory(cards)
}
