package room

import (
	"fmt"
	"game/internal/player"
	"game/internal/poker"
	"reflect"

	"github.com/samber/lo"
)

type RoomPokerDoudizhuStatus int

const (
	RoomPokerDoudizhuStatusRoomIdle = iota + 1
	RoomPokerDoudizhuStatusRoom
	RoomPokerDoudizhuStatusCountDown
)

type RoomPokerDoudizhu struct {
	roomBase
	PokerPlayers           []*poker.PokerPlayer
	PokerPlayersDizhuIndex int
	LeftoverCards          poker.Cards
}

func (r RoomPokerDoudizhu) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r RoomPokerDoudizhu) Players() []*player.Player {
	return lo.Map(r.PokerPlayers, func(t *poker.PokerPlayer, i int) *player.Player {
		return t.Player
	})
}

func (r RoomPokerDoudizhu) HasPlayer(p *player.Player) bool {
	_, ok := lo.Find(r.Players(), func(t *player.Player) bool {
		return p == t
	})
	return ok
}

func (r *RoomPokerDoudizhu) Enter(p *player.Player) error {
	if r.HasPlayer(p) {
		return nil
	}
	if len(r.PokerPlayers) == 3 {
		return fmt.Errorf("room is full")
	}
	if room := FindRoomByPlayerId(p.Id); room != nil && room.Name() != (RoomHall{}).Name() {
		return fmt.Errorf("Enter room error")
	}
	r.PokerPlayers = append(r.PokerPlayers, &poker.PokerPlayer{Player: p})
	r.AppendFrame(NewFramePlayerEnterRoom(p))
	return nil
}

func (r *RoomPokerDoudizhu) Leave(p *player.Player) error {
	index := lo.IndexOf(r.Players(), p)
	if index == -1 {
		return nil
	}
	r.PokerPlayers = append(r.PokerPlayers[:index], r.PokerPlayers[index+1:]...)
	r.AppendFrame(NewFramePlayerLeaveRoom(p))
	return nil
}

func (r *RoomPokerDoudizhu) AppendFrame(f frame) error {
	if err := f.beforeUpdate(r); err != nil {
		return err
	}
	if err := f.update(r); err != nil {
		return err
	}
	r.Frames = append(r.Frames, f)
	return nil
}
