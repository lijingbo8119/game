package room

import (
	"fmt"
	"game/internal/player"

	"github.com/gofrs/uuid"
)

type roomBase struct {
	id      uuid.UUID
	players player.Players
}

func (r roomBase) Id() uuid.UUID {
	if r.id.IsNil() {
		r.id = uuid.Must(uuid.NewV4())
	}
	return r.id
}

func (r roomBase) Name() string {
	return "Base"
}

func (r roomBase) Players() player.Players {
	return r.players
}

func (r roomBase) HasPlayer(p *player.Player) bool {
	return r.Players().Exists(func(_p *player.Player) bool {
		return p == _p
	})
}

func (r roomBase) Enter(p *player.Player) error {
	if r.HasPlayer(p) {
		return nil
	}
	if room := FindRoomByPlayerId(p.Id); room != nil && room.Name() != (RoomHall{}).Name() {
		return fmt.Errorf("Enter room error")
	}
	r.players.Append(p)
	return nil
}

func (r roomBase) Leave(p *player.Player) error {
	if !r.players.Exists(func(p2 *player.Player) bool {
		return p == p2
	}) {
		return nil
	}
	r.players.Remove(p)
	return nil
}
