package landlord_test

import (
	"game/card"
	"game/landlord"
	"testing"
)

func TestCardPatternQuadrupletWithCardsFactory(t *testing.T) {
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
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
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
		p := landlord.FactoryCardPatternQuadrupletWithCards(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternQuadrupletWithCardsFactory failed: %d", index)
		}
	}
}

func TestCardPatternQuadrupletWithCardsGreeter(t *testing.T) {
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
			},
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
			},
		},
	}
	for index, useCase := range useCases {
		p1 := landlord.FactoryCardPatternQuadrupletWithCards(useCase.greeterCards)
		p2 := landlord.FactoryCardPatternQuadrupletWithCards(useCase.cards)
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
