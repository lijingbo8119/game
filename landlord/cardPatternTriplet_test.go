package landlord_test

import (
	"game/landlord"
	"game/poker"
	"testing"
)

func TestCardPatternTripletFactory(t *testing.T) {
	useCases := []struct {
		cards poker.Cards
		valid bool
	}{
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			false,
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo)},
			false,
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			false,
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo)},
			false,
		},
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
			true,
		},
	}
	for index, useCase := range useCases {
		p := landlord.FactoryCardPatternTriplet(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternTripletFactory failed: %d", index)
		}
	}
}

func TestCardPatternTripletGreeter(t *testing.T) {
	useCases := []struct {
		greeterCards poker.Cards
		cards        poker.Cards
	}{
		{
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo), poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo)},
			poker.Cards{poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce), poker.NewCard(poker.CardSuitHeart, poker.CardValueAce)},
		},
	}
	for index, useCase := range useCases {
		p1 := landlord.FactoryCardPatternTriplet(useCase.greeterCards)
		p2 := landlord.FactoryCardPatternTriplet(useCase.cards)
		if !p1.Valid() || !p2.Valid() {
			t.Fatalf("TestCardPatternTripletGreeter failed: Valid: %d", index)
		}
		if !p1.Greeter(p2) || !p2.Lesser(p1) {
			t.Fatalf("TestCardPatternTripletGreeter failed: Greeter: %d", index)
		}
	}
}
