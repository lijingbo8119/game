package game

import (
	"fmt"
	"game/internal/player"
	"reflect"

	"github.com/samber/lo"
)

type GameHall struct {
	gameBase
	players []*player.Player
}

func (r GameHall) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r GameHall) Players() []*player.Player {
	return r.players
}

func (r GameHall) HasPlayer(p *player.Player) bool {
	_, ok := lo.Find(r.Players(), func(t *player.Player) bool {
		return p == t
	})
	return ok
}

func (r *GameHall) Enter(p *player.Player) error {
	if r.HasPlayer(p) {
		return nil
	}
	if game := FindGameByPlayerId(p.Id); game != nil {
		return fmt.Errorf("Enter game error")
	}
	r.players = append(r.players, p)
	return nil
}

func (r *GameHall) Leave(p *player.Player) error {
	index := lo.IndexOf(r.players, p)
	if index == -1 {
		return nil
	}
	r.players = append(r.players[:index], r.players[index+1:]...)
	return nil
}

func (r *GameHall) AppendFrame(f frame) error {
	return nil
}
