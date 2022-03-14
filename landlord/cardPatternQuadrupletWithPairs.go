package landlord

import (
	"game/card"
	"reflect"
)

type cardPatternQuadrupletWithPairs struct {
	cardPatternBase
}

func (r cardPatternQuadrupletWithPairs) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternQuadrupletWithPairs) Valid() bool {
	r.Cards().Sort(LandlordCardValueRanks)
	if r.Cards().Length() != 8 {
		return false
	}
	counts := r.Cards().Counts()
	count4 := r.Cards().Count(func(c *card.Card) bool {
		for value, count := range counts {
			if count == 4 && c.Value() == value {
				return true
			}
		}
		return false
	})
	if !(count4 == 4 || count4 == 8) {
		return false
	}
	if count4 == 4 {
		count2 := r.Cards().Count(func(c *card.Card) bool {
			for value, count := range counts {
				if count == 2 && c.Value() == value {
					return true
				}
			}
			return false
		})
		if count2 != 4 {
			return false
		}
	}
	sum := 0
	for _, count := range counts {
		sum += count
		if count != 4 && count != 2 {
			return false
		}
	}
	return sum == 8
}

func (r cardPatternQuadrupletWithPairs) Same(s card.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternQuadrupletWithPairs) Equal(s card.CardPattern) bool {
	return false
}

func (r cardPatternQuadrupletWithPairs) Greeter(s card.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() {
		return false
	}
	rCard := r.Cards().First(func(c1 *card.Card) bool {
		return r.Cards().Count(func(c2 *card.Card) bool {
			return c1.Value() == c2.Value()
		}) == 4
	})
	sCard := s.Cards().First(func(c1 *card.Card) bool {
		return s.Cards().Count(func(c2 *card.Card) bool {
			return c1.Value() == c2.Value()
		}) == 4
	})
	return LandlordCardValueRanks.Rank(rCard) > LandlordCardValueRanks.Rank(sCard)
}

func (r cardPatternQuadrupletWithPairs) Lesser(s card.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternQuadrupletWithPairs) String() string {
	return ""
}

func (r cardPatternQuadrupletWithPairs) Factory(cards card.Cards) card.CardPattern {
	return cardPatternQuadrupletWithPairs{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternQuadrupletWithPairs(cards card.Cards) card.CardPattern {
	return cardPatternQuadrupletWithPairs{}.Factory(cards)
}
