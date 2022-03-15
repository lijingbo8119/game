package landlord_test

import (
	"game/internal/poker"
	"game/internal/poker/landlord"
	"testing"
)

func TestCardPatternBombFactory(t *testing.T) {
	useCases := []struct {
		cards poker.Cards
		valid bool
	}{
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo)},
			false,
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			false,
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			true,
		},
	}
	for index, useCase := range useCases {
		p := landlord.FactoryCardPatternBomb(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternBombFactory failed: %d", index)
		}
	}
}

func TestCardPatternBombGreeter(t *testing.T) {
	useCases := []struct {
		greeterCards poker.Cards
		cards        poker.Cards
	}{
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo)},
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo)},
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
		},
	}
	for index, useCase := range useCases {
		p1 := landlord.FactoryCardPatternBomb(useCase.greeterCards)
		p2 := landlord.PatternFactory(useCase.cards)
		if !p1.Valid() || !p2.Valid() {
			t.Fatalf("TestCardPatternBombGreeter failed: Valid: %d", index)
		}
		if !p1.Greeter(p2) || !p2.Lesser(p1) {
			t.Fatalf("TestCardPatternBombGreeter failed: Greeter: %d", index)
		}
	}
}
