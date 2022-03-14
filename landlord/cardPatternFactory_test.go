package landlord

import (
	"game/poker"
	"testing"
)

func TestPatternFactory(t *testing.T) {
	useCases := []struct {
		cards       poker.Cards
		patternName string
	}{
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			cardPatternSingle{}.Name(),
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			cardPatternPair{}.Name(),
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo)},
			"",
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			cardPatternTriplet{}.Name(),
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			"",
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
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
