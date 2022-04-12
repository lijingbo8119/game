package poker_test

import (
	"fmt"
	"game/internal/poker"
	"testing"
)

func TestCardsSort(t *testing.T) {
	cards := poker.Cards{
		poker.NewCard(0, 3),
		poker.NewCard(0, 6),
		poker.NewCard(0, 1),
	}
	cards.Sort()
	if cards[0].Value >= cards[1].Value || cards[1].Value >= cards[2].Value {
		t.Fatal("TestCards failed")
	}
}

func TestCardsSuffle(t *testing.T) {
	cards := poker.Cards{
		poker.NewCard(0, 3),
		poker.NewCard(0, 4),
		poker.NewCard(0, 5),
		poker.NewCard(0, 6),
		poker.NewCard(0, 1),
	}
	cards.Shuffle()
	fmt.Println(cards)
}

func TestCardsAppend(t *testing.T) {
	cards := poker.Cards{
		poker.NewCard(0, 3),
		poker.NewCard(0, 4),
		poker.NewCard(0, 5),
		poker.NewCard(0, 6),
		poker.NewCard(0, 1),
	}
	cards.Append(poker.NewCard(0, 2))
	cards.Append(poker.NewCard(0, 2))
	cards.Append(poker.NewCard(0, 2))
	fmt.Println(cards)
}
