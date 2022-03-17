package landlord_test

import (
	"fmt"
	"game/internal/player"
	"game/internal/poker/landlord"
	"testing"
)

func TestRoom(t *testing.T) {
	room := landlord.Room{}
	p := &player.Player{
		Username: "zhangsan",
	}
	room.Enter(p)
	fmt.Println(p, len(room.GetPlayers()))
}
