package room_test

import (
	"fmt"
	"game/internal/player"
	"game/internal/room"
	"game/util"
	"testing"

	"github.com/gofrs/uuid"
)

func TestFramePlayerJiaodizhu(t *testing.T) {
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
	r := room.RoomPokerDoudizhu{}
	r.Enter(p1)
	r.Enter(p2)
	r.Enter(p3)
	r.AppendFrame(room.NewFrameRoomDoudizhuStart())
	r.AppendFrame(room.NewFramePlayerJiaodizhu(p2))
	r.AppendFrame(room.NewFramePlayerJiaodizhu(p3))
	fmt.Println(util.JsonMustMarshalString(r))
}
