package landlord_test

import (
	"game/card"
	"game/landlord"
	"testing"
)

func TestCardPatternTripletFactory(t *testing.T) {
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
			false,
		},
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueTwo), card.NewCard(card.CardSuitHeart, card.CardValueTwo)},
			false,
		},
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce)},
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
		greeterCards card.Cards
		cards        card.Cards
	}{
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueTwo), card.NewCard(card.CardSuitHeart, card.CardValueTwo), card.NewCard(card.CardSuitHeart, card.CardValueTwo)},
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce)},
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
