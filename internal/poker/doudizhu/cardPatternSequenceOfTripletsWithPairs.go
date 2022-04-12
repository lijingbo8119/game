package doudizhu

import (
	"game/internal/poker"
	"reflect"
)

type cardPatternSequenceOfTripletsWithPairs struct {
	cardPatternBase
}

func (r cardPatternSequenceOfTripletsWithPairs) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternSequenceOfTripletsWithPairs) Valid() bool {
	if r.Cards().Exists(func(c *poker.Card) bool {
		return c.Value() == poker.ValueColoredJoker || c.Value() == poker.ValueJoker
	}) {
		return false
	}

	tripletsCardsCounts := r.Cards().Counts(func(val poker.Value, count int) bool {
		return count == 3
	})
	ranks := DoudizhuValueRanks
	if r.Cards().Exists(func(c *poker.Card) bool {
		for val := range tripletsCardsCounts {
			if c.Value() == poker.ValueAce && c.Value() == val {
				return true
			}
		}
		return false
	}) && r.Cards().Exists(func(c *poker.Card) bool {
		for val := range tripletsCardsCounts {
			if c.Value() == poker.ValueTwo && c.Value() == val {
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
			if val == c.Value() {
				return true
			}
		}
		return false
	})
	subCards.Sort(ranks)
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
	pairCardsCounts := r.Cards().Counts(func(val poker.Value, count int) bool {
		return count == 2
	})
	tripletCardsCounts := r.Cards().Counts(func(val poker.Value, count int) bool {
		return count == 3
	})
	quadrupletCardsCounts := r.Cards().Counts(func(val poker.Value, count int) bool {
		return count == 4
	})
	if len(pairCardsCounts)+(len(quadrupletCardsCounts)*2) != len(tripletCardsCounts) {
		return false
	}
	return len(tripletsCardsCounts)*3+len(quadrupletCardsCounts)*4+len(pairCardsCounts)*2 == r.Cards().Length()
}

func (r cardPatternSequenceOfTripletsWithPairs) Same(s poker.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternSequenceOfTripletsWithPairs) Equal(s poker.CardPattern) bool {
	return false
}

func (r cardPatternSequenceOfTripletsWithPairs) Greeter(s poker.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	rCard := r.Cards().Last(func(c1 *poker.Card) bool {
		return r.Cards().Count(func(c2 *poker.Card) bool {
			return c1.Value() == c2.Value()
		}) == 3
	})
	sCard := s.Cards().Last(func(c1 *poker.Card) bool {
		return s.Cards().Count(func(c2 *poker.Card) bool {
			return c1.Value() == c2.Value()
		}) == 3
	})
	return DoudizhuValueRanks.Rank(rCard) > DoudizhuValueRanks.Rank(sCard)
}

func (r cardPatternSequenceOfTripletsWithPairs) Lesser(s poker.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternSequenceOfTripletsWithPairs) String() string {
	return ""
}

func (r cardPatternSequenceOfTripletsWithPairs) Factory(cards poker.Cards) poker.CardPattern {
	return cardPatternSequenceOfTripletsWithPairs{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternSequenceOfTripletsWithPairs(cards poker.Cards) poker.CardPattern {
	return cardPatternSequenceOfTripletsWithPairs{}.Factory(cards)
}
