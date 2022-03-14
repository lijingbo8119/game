package landlord_test

import (
	"game/card"
	"game/landlord"
	"testing"
)

func TestCardPatternTripletWithCardFactory(t *testing.T) {
	useCases := []struct {
		cards card.Cards
		valid bool
	}{
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce)},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
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
		greeterCards card.Cards
		cards        card.Cards
	}{
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
			},
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
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
