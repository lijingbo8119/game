package landlord_test

import (
	"game/internal/poker"
	"game/internal/poker/landlord"
	"testing"
)

func TestCardPatternSingleFactory(t *testing.T) {
	useCases := []struct {
		cards poker.Cards
		valid bool
	}{
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce)},
			true,
		},
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce)},
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
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueTwo)},
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce)},
		},
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueBigJoker)},
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueTwo)},
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
