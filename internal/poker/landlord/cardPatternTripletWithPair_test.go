package landlord_test

import (
	"game/internal/poker"
	"game/internal/poker/landlord"
	"testing"
)

func TestCardPatternTripletWithPairFactory(t *testing.T) {
	useCases := []struct {
		cards poker.Cards
		valid bool
	}{
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce)},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
			},
			true,
		},
	}
	for index, useCase := range useCases {
		p := landlord.FactoryCardPatternTripletWithPair(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternTripletWithPairFactory failed: %d", index)
		}
	}
}

func TestCardPatternTripletWithPairGreeter(t *testing.T) {
	useCases := []struct {
		greeterCards poker.Cards
		cards        poker.Cards
	}{
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
			},
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
			},
		},
	}
	for index, useCase := range useCases {
		p1 := landlord.FactoryCardPatternTripletWithPair(useCase.greeterCards)
		p2 := landlord.FactoryCardPatternTripletWithPair(useCase.cards)
		if !p1.Valid() || !p2.Valid() {
			t.Fatalf("TestCardPatternTripletGreeter failed: Valid: %d", index)
		}
		if !p1.Same(p2) {
			t.Fatalf("TestCardPatternTripletGreeter failed: Same: %d", index)
		}
		if !p1.Greeter(p2) || !p2.Lesser(p1) {
			t.Fatalf("TestCardPatternTripletGreeter failed: Greeter: %d", index)
		}
	}
}
