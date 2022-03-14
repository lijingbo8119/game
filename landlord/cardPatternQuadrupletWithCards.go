package landlord

import (
	"game/card"
	"reflect"
)

type cardPatternQuadrupletWithCards struct {
	cardPatternBase
}

func (r cardPatternQuadrupletWithCards) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternQuadrupletWithCards) Valid() bool {
	r.Cards().Sort(LandlordCardValueRanks)
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

func (r cardPatternQuadrupletWithCards) Same(s card.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternQuadrupletWithCards) Equal(s card.CardPattern) bool {
	return false
}

func (r cardPatternQuadrupletWithCards) Greeter(s card.CardPattern) bool {
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

func (r cardPatternQuadrupletWithCards) Lesser(s card.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternQuadrupletWithCards) String() string {
	return ""
}

func (r cardPatternQuadrupletWithCards) Factory(cards card.Cards) card.CardPattern {
	return cardPatternQuadrupletWithCards{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternQuadrupletWithCards(cards card.Cards) card.CardPattern {
	return cardPatternQuadrupletWithCards{}.Factory(cards)
}
