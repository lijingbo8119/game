package room

import (
	"fmt"
	"game/internal/player"
	"reflect"

	"github.com/samber/lo"
)

type RoomHall struct {
	roomBase
	players []*player.Player
}

func (r RoomHall) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r RoomHall) Players() []*player.Player {
	return r.players
}

func (r RoomHall) HasPlayer(p *player.Player) bool {
	_, ok := lo.Find(r.Players(), func(t *player.Player) bool {
		return p == t
	})
	return ok
}

func (r *RoomHall) Enter(p *player.Player) error {
	if r.HasPlayer(p) {
		return nil
	}
	if room := FindRoomByPlayerId(p.Id); room != nil {
		return fmt.Errorf("Enter room error")
	}
	r.players = append(r.players, p)
	return nil
}

func (r *RoomHall) Leave(p *player.Player) error {
	index := lo.IndexOf(r.players, p)
	if index == -1 {
		return nil
	}
	r.players = append(r.players[:index], r.players[index+1:]...)
	return nil
}

func (r *RoomHall) AppendFrame(f frame) error {
	return nil
}
