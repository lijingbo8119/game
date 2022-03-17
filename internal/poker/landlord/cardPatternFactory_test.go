package landlord

import (
	"game/internal/poker"
	"testing"
)

func TestPatternFactory(t *testing.T) {
	useCases := []struct {
		cards       poker.Cards
		patternName string
	}{
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce)},
			cardPatternSingle{}.Name(),
		},
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce)},
			cardPatternPair{}.Name(),
		},
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueTwo)},
			"",
		},
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce)},
			cardPatternTriplet{}.Name(),
		},
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueTwo), poker.NewCard(poker.SuitHeart, poker.ValueAce)},
			"",
		},
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce)},
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
