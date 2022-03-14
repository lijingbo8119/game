package landlord

import (
	"game/card"
	"reflect"
)

type cardPatternSequence struct {
	cardPatternBase
}

func (r cardPatternSequence) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r cardPatternSequence) Valid() bool {
	ranks := LandlordCardValueRanks
	if r.Cards().Exists(func(c *card.Card) bool {
		return c.Value() == card.CardValueAce
	}) && r.Cards().Exists(func(c *card.Card) bool {
		return c.Value() == card.CardValueTwo
	}) {
		ranks = card.CardValueSortRanks
	}
	r.Cards().Sort(ranks)
	if r.Cards().Exists(func(c *card.Card) bool {
		return c.Value() == card.CardValueBigJoker || c.Value() == card.CardValueSmallJoker
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

func (r cardPatternSequence) Same(s card.CardPattern) bool {
	return r.Name() == s.Name()
}

func (r cardPatternSequence) Equal(s card.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	return r.Cards().First().Value() == s.Cards().First().Value()
}

func (r cardPatternSequence) Greeter(s card.CardPattern) bool {
	if !r.Same(s) || !r.Valid() || !s.Valid() || r.Cards().Length() != s.Cards().Length() {
		return false
	}
	return LandlordCardValueRanks.Rank(r.Cards().Last()) > LandlordCardValueRanks.Rank(s.Cards().Last())
}

func (r cardPatternSequence) Lesser(s card.CardPattern) bool {
	return s.Greeter(r)
}

func (r cardPatternSequence) String() string {
	return ""
}

func (r cardPatternSequence) Factory(cards card.Cards) card.CardPattern {
	return cardPatternSequence{cardPatternBase: cardPatternBase{cards: cards}}
}

func FactoryCardPatternSequence(cards card.Cards) card.CardPattern {
	return cardPatternSequence{}.Factory(cards)
}
