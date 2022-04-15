package game_test

import (
	"game/internal/game"
	"game/internal/player"
	"testing"

	"github.com/gofrs/uuid"
)

func TestGamePokerDoudizhuEnter(t *testing.T) {
	p1 := &player.Player{
		Id:       uuid.Must(uuid.NewV4()),
		Nickname: "张三",
	}
	p2 := &player.Player{
		Id:       uuid.Must(uuid.NewV4()),
		Nickname: "李四",
	}
	p3 := &player.Player{
		Id:       uuid.Must(uuid.NewV4()),
		Nickname: "王五",
	}
	r := game.GamePokerDoudizhu{}
	r.Enter(p1)
	r.Enter(p2)
	r.Enter(p3)
	if len(r.Players()) != 3 {
		t.Fatalf("TestGamePokerDoudizhuEnter failed")
	}
}
