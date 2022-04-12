package doudizhu_test

import (
	"game/internal/poker"
	"game/internal/poker/doudizhu"
	"testing"
)

func TestCardPatternTripletWithCardFactory(t *testing.T) {
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
			true,
		},
	}
	for index, useCase := range useCases {
		p := doudizhu.FactoryCardPatternTripletWithCard(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternTripletWithCardFactory failed: %d", index)
		}
	}
}

func TestCardPatternTripletWithCardGreeter(t *testing.T) {
	useCases := []struct {
		greeterCards poker.Cards
		cards        poker.Cards
	}{
		{
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
			},
			poker.Cards{
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueAce),
				poker.NewCard(poker.SuitHeart, poker.ValueTwo),
			},
		},
	}
	for index, useCase := range useCases {
		p1 := doudizhu.FactoryCardPatternTripletWithCard(useCase.greeterCards)
		p2 := doudizhu.FactoryCardPatternTripletWithCard(useCase.cards)
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
