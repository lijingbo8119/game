package room

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"game/internal/poker"
	"reflect"
	"time"

	"github.com/samber/lo"
)

type framePlayerPlayCards struct {
	T      time.Time      `json:"time"`
	Player *player.Player `json:"player"`
	Cards  poker.Cards    `json:"cards"`
}

func (r framePlayerPlayCards) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerPlayCards) Time() time.Time {
	return r.T
}

func (r framePlayerPlayCards) beforeUpdate(ro Room) error {
	f1 := ro.lastFrame(func(f frame) bool {
		return f.Name() == frameRoomPlayerRoundStart{}.Name()
	})
	if f1 == nil || f1.(frameRoomPlayerRoundStart).Player != r.Player {
		return fmt.Errorf("framePlayerPlayCards beforeUpdate error: last frameRoomPlayerRoundStart error: %v", r)
	}
	f2 := ro.lastFrame(func(f frame) bool {
		_f, ok := f.(frameRoomPlayerRoundEnd)
		if !ok {
			return false
		}
		if _f.Player == r.Player && _f.Time().After(f1.Time()) {
			return true
		}
		return false
	})
	if f2 != nil {
		return fmt.Errorf("framePlayerPlayCards beforeUpdate error: player round ended: %v", r)
	}
	return nil
}

func (r framePlayerPlayCards) update(ro Room) error {
	roomDoudizhu, ok := ro.(*RoomPokerDoudizhu)
	if !ok {
		return fmt.Errorf("room type error")
	}
	pokerPlayer, ok := lo.Find(roomDoudizhu.PokerPlayers, func(t *poker.PokerPlayer) bool {
		return t.Player == r.Player
	})
	if !ok {
		return fmt.Errorf("find pokerPlayer error")
	}
	pokerPlayer.Cards.Remove(r.Cards...)
	return nil
}

func (r framePlayerPlayCards) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":   r.Name(),
		"time":   r.T,
		"player": r.Player,
		"cards":  r.Cards,
	}
	return json.Marshal(m)
}

func NewFrameRoomDoudizhuPlayerPlay(p *player.Player, cards poker.Cards) frame {
	return framePlayerPlayCards{
		Player: p,
		T:      time.Now(),
		Cards:  cards,
	}
}
