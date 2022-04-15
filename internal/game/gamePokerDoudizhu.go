package game

import (
	"fmt"
	"game/internal/player"
	"game/internal/poker"
	"reflect"

	"github.com/samber/lo"
)

type GamePokerDoudizhuStatus int

const (
	GamePokerDoudizhuStatusGameIdle = iota + 1
	GamePokerDoudizhuStatusGame
	GamePokerDoudizhuStatusCountDown
)

type GamePokerDoudizhu struct {
	gamePokerBase
	PokerPlayersDizhuIndex int
	LeftoverCards          poker.Cards
}

func (r GamePokerDoudizhu) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r *GamePokerDoudizhu) Enter(p *player.Player) error {
	if r.HasPlayer(p) {
		return nil
	}
	if len(r.PokerPlayers) == 3 {
		return fmt.Errorf("game is full")
	}
	if game := FindGameByPlayerId(p.Id); game != nil && game.Name() != (GameHall{}).Name() {
		return fmt.Errorf("Enter game error")
	}
	r.PokerPlayers = append(r.PokerPlayers, &poker.PokerPlayer{Player: p})
	r.AppendFrame(NewFramePlayerEnterGame(p))
	return nil
}

func (r *GamePokerDoudizhu) Leave(p *player.Player) error {
	index := lo.IndexOf(r.Players(), p)
	if index == -1 {
		return nil
	}
	r.PokerPlayers = append(r.PokerPlayers[:index], r.PokerPlayers[index+1:]...)
	r.AppendFrame(NewFramePlayerLeaveGame(p))
	return nil
}

func (r *GamePokerDoudizhu) AppendFrame(f frame) error {
	if err := f.beforeUpdate(r); err != nil {
		return err
	}
	if err := f.update(r); err != nil {
		return err
	}
	r.Frames = append(r.Frames, f)
	return nil
}
