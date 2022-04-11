package room

import (
	"fmt"
	"game/internal/player"
	"game/internal/poker"
)

type RoomPokerLandlordStatus int

const (
	RoomPokerLandlordStatusIdle = iota + 1
	RoomPokerLandlordStatusPlaying
)

type RoomPokerLandlord struct {
	roomBase
	status RoomPokerLandlordStatus
	frames []poker.Frame
}

func (r RoomPokerLandlord) Enter(p *player.Player) error {
	if r.players.Exists(func(p2 *player.Player) bool {
		return p == p2
	}) {
		return nil
	}
	if r.players.Length() == 3 {
		return fmt.Errorf("room is full")
	}
	return r.roomBase.Enter(p)
}
