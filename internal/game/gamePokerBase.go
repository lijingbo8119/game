package game

import (
	"game/internal/player"
	"game/internal/poker"

	"github.com/samber/lo"
)

type gamePokerBase struct {
	gameBase
	PokerPlayers []*poker.PokerPlayer
}

func (r gamePokerBase) Players() []*player.Player {
	return lo.Map(r.PokerPlayers, func(t *poker.PokerPlayer, i int) *player.Player {
		return t.Player
	})
}

func (r gamePokerBase) HasPlayer(p *player.Player) bool {
	_, ok := lo.Find(r.Players(), func(t *player.Player) bool {
		return p == t
	})
	return ok
}
