package landlord_test

import (
	"game/landlord"
	"game/poker"
	"testing"
)

func TestCardPatternSingleFactory(t *testing.T) {
	useCases := []struct {
		cards poker.Cards
		valid bool
	}{
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			true,
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			false,
		},
	}
	for index, useCase := range useCases {
		p := landlord.FactoryCardPatternSingle(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternSingleFactory failed: %d", index)
		}
	}
}

func TestCardPatternSingleGreeter(t *testing.T) {
	useCases := []struct {
		greeterCards poker.Cards
		cards        poker.Cards
	}{
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo)},
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueBigJoker)},
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo)},
		},
	}
	for index, useCase := range useCases {
		p1 := landlord.FactoryCardPatternSingle(useCase.greeterCards)
		p2 := landlord.FactoryCardPatternSingle(useCase.cards)
		if !p1.Valid() || !p2.Valid() {
			t.Fatalf("TestCardPatternSingleGreeter failed: Valid: %d", index)
		}
		if !p1.Greeter(p2) || !p2.Lesser(p1) {
			t.Fatalf("TestCardPatternSingleGreeter failed: Greeter: %d", index)
		}
	}
}
