package landlord_test

import (
	"game/card"
	"game/landlord"
	"testing"
)

func TestCardPatternPairFactory(t *testing.T) {
	useCases := []struct {
		cards card.Cards
		valid bool
	}{
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce)},
			false,
		},
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueTwo)},
			false,
		},
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce)},
			true,
		},
	}
	for index, useCase := range useCases {
		p := landlord.FactoryCardPatternPair(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternPairFactory failed: %d", index)
		}
	}
}

func TestCardPatternPairGreeter(t *testing.T) {
	useCases := []struct {
		greeterCards card.Cards
		cards        card.Cards
	}{
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueTwo), card.NewCard(card.CardSuitHeart, card.CardValueTwo)},
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce)},
		},
	}
	for index, useCase := range useCases {
		p1 := landlord.FactoryCardPatternPair(useCase.greeterCards)
		p2 := landlord.FactoryCardPatternPair(useCase.cards)
		if !p1.Valid() || !p2.Valid() {
			t.Fatalf("TestCardPatternPairGreeter failed: Valid: %d", index)
		}
		if !p1.Greeter(p2) || !p2.Lesser(p1) {
			t.Fatalf("TestCardPatternPairGreeter failed: Greeter: %d", index)
		}
	}
}
