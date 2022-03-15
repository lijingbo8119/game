package landlord_test

import (
	"fmt"
	"game/internal"
	"game/internal/poker/landlord"
	"testing"
)

func TestRoom(t *testing.T) {
	room := landlord.Room{}
	p := &internal.Player{
		Name: "zhangsan",
	}
	room.Enter(p)
	fmt.Println(p, len(room.GetPlayers()))
}
