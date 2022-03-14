package landlord_test

import (
	"game/landlord"
	"game/poker"
	"testing"
)

func TestCardPatternTripletWithCardFactory(t *testing.T) {
	useCases := []struct {
		cards poker.Cards
		valid bool
	}{
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
			},
			true,
		},
	}
	for index, useCase := range useCases {
		p := landlord.FactoryCardPatternTripletWithCard(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternTripletWithCardFactory failed: %d", index)
		}
	}
}

func TestCardPatternTripletWithCardGreeter(t *testing.T) {
	useCases := []struct {
		greeterCards poker.Cards
		cards        poker.Cards
	}{
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
			},
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
			},
		},
	}
	for index, useCase := range useCases {
		p1 := landlord.FactoryCardPatternTripletWithCard(useCase.greeterCards)
		p2 := landlord.FactoryCardPatternTripletWithCard(useCase.cards)
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
