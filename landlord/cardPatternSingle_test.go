package landlord_test

import (
	"game/card"
	"game/landlord"
	"testing"
)

func TestCardPatternSingleFactory(t *testing.T) {
	useCases := []struct {
		cards card.Cards
		valid bool
	}{
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce)},
			true,
		},
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce), card.NewCard(card.CardSuitHeart, card.CardValueAce)},
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
		greeterCards card.Cards
		cards        card.Cards
	}{
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueTwo)},
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueAce)},
		},
		{
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueBigJoker)},
			card.Cards{card.NewCard(card.CardSuitHeart, card.CardValueTwo)},
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
