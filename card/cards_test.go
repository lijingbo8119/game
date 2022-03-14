package card_test

import (
	"fmt"
	"game/card"
	"testing"
)

func TestCardsSort(t *testing.T) {
	cards := card.Cards{
		card.NewCard(0, 3),
		card.NewCard(0, 6),
		card.NewCard(0, 1),
	}
	cards.Sort()
	if cards[0].Value() >= cards[1].Value() || cards[1].Value() >= cards[2].Value() {
		t.Fatal("TestCards failed")
	}
}

func TestCardsSuffle(t *testing.T) {
	cards := card.Cards{
		card.NewCard(0, 3),
		card.NewCard(0, 4),
		card.NewCard(0, 5),
		card.NewCard(0, 6),
		card.NewCard(0, 1),
	}
	cards.Shuffle()
	fmt.Println(cards)
}
