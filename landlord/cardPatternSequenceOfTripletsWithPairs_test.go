package landlord_test

import (
	"game/card"
	"game/landlord"
	"testing"
)

func TestCardPatternSequenceOfTripletsWithPairsFactory(t *testing.T) {
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
			false,
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
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueFour),
				card.NewCard(card.CardSuitHeart, card.CardValueFive),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueFour),
				card.NewCard(card.CardSuitHeart, card.CardValueFive),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueTen),
				card.NewCard(card.CardSuitHeart, card.CardValueJack),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
			},
			false,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueJack),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueJack),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
			},
			true,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
			},
			true,
		},
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
			},
			true,
		},
	}
	for index, useCase := range useCases {
		p := landlord.FactoryCardPatternSequenceOfTripletsWithPairs(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternSequenceOfTripletsWithPairsFactory failed: %d: %s", index, useCase.cards)
		}
	}
}

func TestCardPatternSequenceOfTripletsWithPairsGreeter(t *testing.T) {
	useCases := []struct {
		greeterCards card.Cards
		cards        card.Cards
	}{
		{
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
			},
			card.Cards{
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueAce),
				card.NewCard(card.CardSuitHeart, card.CardValueTwo),
				card.NewCard(card.CardSuitHeart, card.CardValueThree),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueKing),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueQueen),
				card.NewCard(card.CardSuitHeart, card.CardValueJack),
				card.NewCard(card.CardSuitHeart, card.CardValueJack),
			},
		},
	}
	for index, useCase := range useCases {
		p1 := landlord.FactoryCardPatternSequenceOfTripletsWithPairs(useCase.greeterCards)
		p2 := landlord.FactoryCardPatternSequenceOfTripletsWithPairs(useCase.cards)
		if !p1.Valid() || !p2.Valid() {
			t.Fatalf("TestCardPatternQuadrupletGreeter failed: Valid: %d: %s: %s", index, useCase.greeterCards, useCase.cards)
		}
		if !p1.Same(p2) {
			t.Fatalf("TestCardPatternQuadrupletGreeter failed: Same: %d", index)
		}
		if !p1.Greeter(p2) || !p2.Lesser(p1) {
			t.Fatalf("TestCardPatternQuadrupletGreeter failed: Greeter: %d: %s: %s", index, useCase.greeterCards, useCase.cards)
		}
	}
}
