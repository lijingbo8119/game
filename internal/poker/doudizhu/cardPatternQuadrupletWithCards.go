package doudizhu

import (
	"game/internal/poker"
	"reflect"
)

type cardPatternQuadrupletWithCards struct {
	cardPatternBase
}

func (r cardPatternQuadrupletWithCards) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternQuadrupletWithCards) Valid() bool {
	r.Cards().Sort(DoudizhuValueRanks)
	if r.Cards().Length() != 6 {
		return false
	}
	counts := r.Cards().Counts()
	hasCount4 := func() bool {
		for _, count := range counts {
			if count == 4 {
				return true
			}
		}
		return false
	}()
	if !hasCount4 {
		return false
	}
	sum := 0
	for _, count := range counts {
		sum += count
		if count != 4 && count != 1 && count != 2 {
			return false
		}
	}
	return sum == 6
}

func (r cardPatternQuadrupletWithCards) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternQuadrupletWithCards) Equal(s poker.CardPattern) bool {
	return false
}

func (r cardPatternQuadrupletWithCards) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	rCard := r.Cards().First(func(c1 *poker.Card) bool {
		return r.Cards().Count(func(c2 *poker.Card) bool {
			return c1.Value == c2.Value
		}) == 4
	})
	sCard := s.Cards().First(func(c1 *poker.Card) bool {
		return s.Cards().Count(func(c2 *poker.Card) bool {
			return c1.Value == c2.Value
		}) == 4
	})
	return DoudizhuValueRanks.Rank(rCard) > DoudizhuValueRanks.Rank(sCard)
}

func (r cardPatternQuadrupletWithCards) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternQuadrupletWithCards) String() string {
	return ""
}

func (r cardPatternQuadrupletWithCards) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternQuadrupletWithCards{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternQuadrupletWithCards(cards poker.Cards) poker.CardPattern {
	return cardPatternQuadrupletWithCards{}.Factory(cards)
}
