package doudizhu_test

import (
	"fmt"
	"game/internal/player"
	"game/internal/poker/doudizhu"
	"testing"
)

func TestGame(t *testing.T) {
	game := doudizhu.Game{}
	p := &player.Player{
		Username: "zhangsan",
	}
	game.Enter(p)
	fmt.Println(p, len(game.GetPlayers()))
}
