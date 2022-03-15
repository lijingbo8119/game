package room

import (
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

func (r roomBase) Enter(p *player.Player) bool {
	if r.players.Exists(func(p2 *player.Player) bool {
		return p == p2
	}) {
		return true
	}
	if room := FindRoomByPlayerId(p.Id); room != nil {
		return false
	}
	r.players.Append(p)
	return true
}

func (r roomBase) Leave(p *player.Player) bool {
	if !r.players.Exists(func(p2 *player.Player) bool {
		return p == p2
	}) {
		return false
	}
	r.players.Remove(p)
	return true
}
