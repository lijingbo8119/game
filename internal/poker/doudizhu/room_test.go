package doudizhu_test

import (
	"fmt"
	"game/internal/player"
	"game/internal/poker/doudizhu"
	"testing"
)

func TestRoom(t *testing.T) {
	room := doudizhu.Room{}
	p := &player.Player{
		Username: "zhangsan",
	}
	room.Enter(p)
	fmt.Println(p, len(room.GetPlayers()))
}
