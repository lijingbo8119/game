package doudizhu

import (
	"game/internal/poker"
	"reflect"
)

type cardPatternSequenceOfTriplets struct {
	cardPatternBase
}

func (r cardPatternSequenceOfTriplets) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternSequenceOfTriplets) Valid() bool {
	ranks := DoudizhuValueRanks
	if r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value == poker.ValueAce
	}) && r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value == poker.ValueTwo
	}) {
		ranks = poker.ValueSortRanks
	}
	r.Cards().Sort(ranks)
	if r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value == poker.ValueColoredJoker || c.Value == poker.ValueJoker
	}) {
		return false
	}
	subCardsCounts := r.Cards().Counts(func(val poker.Value, count int) bool {
		return count == 3
	})
	if len(subCardsCounts) < 2 {
		return false
	}
	if len(
		r.Cards().Counts(func(val poker.Value, count int) bool {
			return count != 3
		}),
	) > 0 {
		return false
	}
	for i := 0; i < r.Cards().Length()-1; i++ {
		current := r.Cards()[i]
		next := r.Cards()[i+1]
		switch ranks.Rank(next) - ranks.Rank(current) {
		case 0:
		case 1:
		default:
			return false
		}
	}
	return true
}

func (r cardPatternSequenceOfTriplets) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternSequenceOfTriplets) Equal(s poker.CardPattern) bool {
	return false
}

func (r cardPatternSequenceOfTriplets) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	return DoudizhuValueRanks.Rank(r.Cards().Last()) > DoudizhuValueRanks.Rank(s.Cards().Last())
}

func (r cardPatternSequenceOfTriplets) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternSequenceOfTriplets) String() string {
	return ""
}

func (r cardPatternSequenceOfTriplets) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternSequenceOfTriplets{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternSequenceOfTriplets(cards poker.Cards) poker.CardPattern {
	return cardPatternSequenceOfTriplets{}.Factory(cards)
}
