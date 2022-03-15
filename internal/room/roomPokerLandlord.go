package room

import "game/internal/player"

type RoomPokerLandlord struct {
	roomBase
}

func (r RoomPokerLandlord) Enter(p *player.Player) bool {
	if r.players.Exists(func(p2 *player.Player) bool {
		return p == p2
	}) {
		return true
	}
	if r.players.Length() == 3 {
		return false
	}
	return r.roomBase.Enter(p)
}
