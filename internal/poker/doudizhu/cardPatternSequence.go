package doudizhu

import (
	"game/internal/poker"
	"reflect"
)

type cardPatternSequence struct {
	cardPatternBase
}

func (r cardPatternSequence) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternSequence) Valid() bool {
	ranks := DoudizhuValueRanks
	if r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value() == poker.ValueAce
	}) && r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value() == poker.ValueTwo
	}) {
		ranks = poker.ValueSortRanks
	}
	r.Cards().Sort(ranks)
	if r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value() == poker.ValueColoredJoker || c.Value() == poker.ValueJoker
	}) {
		return false
	}
	if r.Cards().Length() < 5 {
		return false
	}
	for i := 0; i < r.Cards().Length()-1; i++ {
		current := r.Cards()[i]
		next := r.Cards()[i+1]
		if ranks.Rank(next)-ranks.Rank(current) != 1 {
			return false
		}
	}
	return true
}

func (r cardPatternSequence) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternSequence) Equal(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	return r.Cards().First().Value() == s.Cards().First().Value()
}

func (r cardPatternSequence) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	return DoudizhuValueRanks.Rank(r.Cards().Last()) > DoudizhuValueRanks.Rank(s.Cards().Last())
}

func (r cardPatternSequence) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternSequence) String() string {
	return ""
}

func (r cardPatternSequence) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternSequence{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternSequence(cards poker.Cards) poker.CardPattern {
	return cardPatternSequence{}.Factory(cards)
}
