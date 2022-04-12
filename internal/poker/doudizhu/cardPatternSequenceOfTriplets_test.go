package doudizhu_test

import (
	"game/internal/poker"
	"game/internal/poker/doudizhu"
	"testing"
)

func TestCardPatternSequenceOfTripletsFactory(t *testing.T) {
	useCases := []struct {
		cards poker.Cards
		valid bool
	}{
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce)},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueThree),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueThree),
				poker.NewCard(poker.SuitHeart, poker.ValueFour),
				poker.NewCard(poker.SuitHeart, poker.ValueFive),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueThree),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueFour),
				poker.NewCard(poker.SuitHeart, poker.ValueFive),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueTen),
				poker.NewCard(poker.SuitHeart, poker.ValueJack),
				poker.NewCard(poker.SuitHeart, poker.ValueQueen),
				poker.NewCard(poker.SuitHeart, poker.ValueKing),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueThree),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueThree),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueQueen),
				poker.NewCard(poker.SuitHeart, poker.ValueKing),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueQueen),
				poker.NewCard(poker.SuitHeart, poker.ValueKing),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
			},
			false,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueThree),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueThree),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueThree),
			},
			true,
		},
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueQueen),
				poker.NewCard(poker.SuitHeart, poker.ValueKing),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueQueen),
				poker.NewCard(poker.SuitHeart, poker.ValueKing),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueQueen),
				poker.NewCard(poker.SuitHeart, poker.ValueKing),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
			},
			true,
		},
	}
	for index, useCase := range useCases {
		p := doudizhu.FactoryCardPatternSequenceOfTriplets(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternSequenceOfTripletsFactory failed: %d: %s", index, useCase.cards)
		}
	}
}

func TestCardPatternSequenceOfTripletsGreeter(t *testing.T) {
	useCases := []struct {
		greeterCards poker.Cards
		cards        poker.Cards
	}{
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueQueen),
				poker.NewCard(poker.SuitHeart, poker.ValueKing),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueQueen),
				poker.NewCard(poker.SuitHeart, poker.ValueKing),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueQueen),
				poker.NewCard(poker.SuitHeart, poker.ValueKing),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
			},
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueThree),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueThree),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueThree),
			},
		},
	}
	for index, useCase := range useCases {
		p1 := doudizhu.FactoryCardPatternSequenceOfTriplets(useCase.greeterCards)
		p2 := doudizhu.FactoryCardPatternSequenceOfTriplets(useCase.cards)
		if !p1.Valid() || !p2.Valid() {
			t.Fatalf("TestCardPatternQuadrupletGreeter failed: Valid: %d", index)
		}
		if !p1.Same(p2) {
			t.Fatalf("TestCardPatternQuadrupletGreeter failed: Same: %d", index)
		}
		if !p1.Greeter(p2) || !p2.Lesser(p1) {
			t.Fatalf("TestCardPatternQuadrupletGreeter failed: Greeter: %d: %s: %s", index, useCase.greeterCards, useCase.cards)
		}
	}
}
