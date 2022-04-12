package doudizhu

import (
	"game/internal/poker"
	"reflect"
)

type cardPatternSequenceOfTripletsWithCards struct {
	cardPatternBase
}

func (r cardPatternSequenceOfTripletsWithCards) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternSequenceOfTripletsWithCards) Valid() bool {
	if r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value == poker.ValueColoredJoker || c.Value == poker.ValueJoker
	}) {
		return false
	}

	tripletsCardsCounts := r.Cards().Counts(func(val poker.Value, count int) bool {
		return count == 3 || count == 4
	})
	ranks := DoudizhuValueRanks
	if r.Cards().Exists(func(c *poker.Card) bool {
		for val := range tripletsCardsCounts {
			if c.Value == poker.ValueAce && c.Value == val {
				return true
			}
		}
		return false
	}) && r.Cards().Exists(func(c *poker.Card) bool {
		for val := range tripletsCardsCounts {
			if c.Value == poker.ValueTwo && c.Value == val {
				return true
			}
		}
		return false
	}) {
		ranks = poker.ValueSortRanks
	}
	r.Cards().Sort(ranks)

	if len(tripletsCardsCounts) < 2 {
		return false
	}
	subCards := r.Cards().SubCards(func(c *poker.Card) bool {
		for val := range tripletsCardsCounts {
			if val == c.Value {
				return true
			}
		}
		return false
	})
	subCards.Sort()
	for i := 0; i < subCards.Length()-1; i++ {
		current := subCards[i]
		next := subCards[i+1]
		switch ranks.Rank(next) - ranks.Rank(current) {
		case 0:
		case 1:
		default:
			return false
		}
	}
	return len(tripletsCardsCounts)*(3+1) == r.Cards().Length()
}

func (r cardPatternSequenceOfTripletsWithCards) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternSequenceOfTripletsWithCards) Equal(s poker.CardPattern) bool {
	return false
}

func (r cardPatternSequenceOfTripletsWithCards) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	rCard := r.Cards().Last(func(c1 *poker.Card) bool {
		return r.Cards().Count(func(c2 *poker.Card) bool {
			return c1.Value == c2.Value
		}) >= 3
	})
	sCard := s.Cards().Last(func(c1 *poker.Card) bool {
		return s.Cards().Count(func(c2 *poker.Card) bool {
			return c1.Value == c2.Value
		}) >= 3
	})
	return DoudizhuValueRanks.Rank(rCard) > DoudizhuValueRanks.Rank(sCard)
}

func (r cardPatternSequenceOfTripletsWithCards) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternSequenceOfTripletsWithCards) String() string {
	return ""
}

func (r cardPatternSequenceOfTripletsWithCards) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternSequenceOfTripletsWithCards{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternSequenceOfTripletsWithCards(cards poker.Cards) poker.CardPattern {
	return cardPatternSequenceOfTripletsWithCards{}.Factory(cards)
}
