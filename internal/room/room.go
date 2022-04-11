package room

import (
	"game/internal/player"

	"github.com/gofrs/uuid"
)

type Room interface {
	Id() uuid.UUID
	Name() string
	Players() []*player.Player
	HasPlayer(p *player.Player) bool
	Enter(*player.Player) error
	Leave(*player.Player) error
}
