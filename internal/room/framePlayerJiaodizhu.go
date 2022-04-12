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

type framePlayerJiaodizhu struct {
	T      time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerJiaodizhu) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerJiaodizhu) Time() time.Time {
	return r.T
}

func (r framePlayerJiaodizhu) beforeUpdate(ro Room) error {
	roomDoudizhu, ok := ro.(*RoomPokerDoudizhu)
	if !ok {
		return fmt.Errorf("framePlayerJiaodizhu: beforeUpdate: Room type error")
	}
	f := roomDoudizhu.lastFrame()
	switch f.(type) {
	case framePlayerJiaodizhu:
		break
	case frameRoomPlayerRoundStart:
		break
	case frameRoomPlayerRoundEnd:
		break
	default:
		return fmt.Errorf("framePlayerJiaodizhu: beforeUpdate: frame type error")
	}
	return nil
}

func (r framePlayerJiaodizhu) update(ro Room) error {
	roomDoudizhu, ok := ro.(*RoomPokerDoudizhu)
	if !ok {
		return fmt.Errorf("room type error")
	}
	pokerPlayer, ok := lo.Find(roomDoudizhu.PokerPlayers, func(t *poker.PokerPlayer) bool {
		return t.Player == r.Player
	})
	if !ok {
		return fmt.Errorf("framePlayerJiaodizhu find pokerPlayer error")
	}
	roomDoudizhu.PokerPlayersDizhuIndex = lo.IndexOf(roomDoudizhu.PokerPlayers, pokerPlayer)
	return nil
}

func (r framePlayerJiaodizhu) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":   r.Name(),
		"time":   r.T,
		"player": r.Player,
	}
	return json.Marshal(m)
}

func NewFramePlayerJiaodizhu(p *player.Player) frame {
	return framePlayerJiaodizhu{
		Player: p,
		T:      time.Now(),
	}
}
