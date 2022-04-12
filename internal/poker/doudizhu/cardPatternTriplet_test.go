package doudizhu_test

import (
	"game/internal/poker"
	"game/internal/poker/doudizhu"
	"testing"
)

func TestCardPatternTripletFactory(t *testing.T) {
	useCases := []struct {
		cards poker.Cards
		valid bool
	}{
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce)},
			false,
		},
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueTwo)},
			false,
		},
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce)},
			false,
		},
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueTwo), poker.NewCard(poker.SuitHeart, poker.ValueTwo)},
			false,
		},
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce)},
			true,
		},
	}
	for index, useCase := range useCases {
		p := doudizhu.FactoryCardPatternTriplet(useCase.cards)
		if p.Valid() != useCase.valid {
			t.Fatalf("TestCardPatternTripletFactory failed: %d", index)
		}
	}
}

func TestCardPatternTripletGreeter(t *testing.T) {
	useCases := []struct {
		greeterCards poker.Cards
		cards        poker.Cards
	}{
		{
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueTwo), poker.NewCard(poker.SuitHeart, poker.ValueTwo), poker.NewCard(poker.SuitHeart, poker.ValueTwo)},
			poker.Cards{poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce), poker.NewCard(poker.SuitHeart, poker.ValueAce)},
		},
	}
	for index, useCase := range useCases {
		p1 := doudizhu.FactoryCardPatternTriplet(useCase.greeterCards)
		p2 := doudizhu.FactoryCardPatternTriplet(useCase.cards)
		if !p1.Valid() || !p2.Valid() {
			t.Fatalf("TestCardPatternTripletGreeter failed: Valid: %d", index)
		}
		if !p1.Greeter(p2) || !p2.Lesser(p1) {
			t.Fatalf("TestCardPatternTripletGreeter failed: Greeter: %d", index)
		}
	}
}
