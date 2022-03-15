package landlord_test

import (
	"game/internal/poker"
	"game/internal/poker/landlord"
	"testing"
)

func TestCardPatternSequenceOfTripletsWithPairsFactory(t *testing.T) {
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
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
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
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueFour),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueFive),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueFour),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueFive),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueJack),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueJack),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueJack),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
			},
			true,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
			},
			true,
		},
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
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
		greeterCards poker.Cards
		cards        poker.Cards
	}{
		{
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
			},
			poker.Cards{
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueAce),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueTwo),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueThree),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueKing),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueQueen),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueJack),
				poker.NewCard(poker.CardSuitHeart, poker.CardValueJack),
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
