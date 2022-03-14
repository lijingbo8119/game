package landlord

import (
	"game/card"
	"testing"
)

func TestPatternFactory(t *testing.T) {
	useCases := []struct {
		cards       card.Cards
		patternName string
	}{
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce)},
			cardPatternSingle{}.Name(),
		},
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce)},
			cardPatternPair{}.Name(),
		},
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueTwo)},
			"",
		},
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce)},
			cardPatternTriplet{}.Name(),
		},
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueTwo), card.NewCard(card.CardSuitHeart, card.CardValueAce)},
			"",
		},
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce)},
			cardPatternBomb{}.Name(),
		},
	}
	for index, useCase := range useCases {
		p := PatternFactory(useCase.cards)
		if useCase.patternName == "" && p == nil {
			continue
		}
		if p.Name() != useCase.patternName {
			t.Fatalf("TestPatternFactory failed: mismatch patternName: %d", index)
		}
	}
}
