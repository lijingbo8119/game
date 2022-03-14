package landlord_test

import (
	"game/card"
	"game/landlord"
	"testing"
)

func TestCardPatternQuadrupletWithPairsFactory(t *testing.T) {
	useCases := []struct {
		cards card.Cards
		valid bool
	}{
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
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
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
			},
			true,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
			},
			true,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
			},
			false,
		},
	}
	for index, useCase := range useCases {
		p := landlord.FactoryCardPatternQuadrupletWithPairs(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternQuadrupletWithPairsFactory failed: %d", index)
		}
	}
}

func TestCardPatternQuadrupletWithPairsGreeter(t *testing.T) {
	useCases := []struct {
		greeterCards card.Cards
		cards        card.Cards
	}{
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
			},
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
			},
		},
	}
	for index, useCase := range useCases {
		p1 := landlord.FactoryCardPatternQuadrupletWithPairs(useCase.greeterCards)
		p2 := landlord.FactoryCardPatternQuadrupletWithPairs(useCase.cards)
		if !p1.Valid() || !p2.Valid() {
			t.Fatalf("TestCardPatternQuadrupletGreeter failed: Valid: %d", index)
		}
		if !p1.Same(p2) {
			t.Fatalf("TestCardPatternQuadrupletGreeter failed: Same: %d", index)
		}
		if !p1.Greeter(p2) || !p2.Lesser(p1) {
			t.Fatalf("TestCardPatternQuadrupletGreeter failed: Greeter: %d", index)
		}
	}
}
