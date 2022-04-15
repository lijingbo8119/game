package game_test

import (
	"fmt"
	"game/internal/game"
	"game/internal/player"
	"game/util"
	"testing"

	"github.com/gofrs/uuid"
)

func TestFramePlayerPlayCards(t *testing.T) {
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
	f1 := game.NewFrameGamePokerDoudizhuStart()
	r.AppendFrame(f1)
	fmt.Println(util.JsonMustMarshalString(r))
}
