package room

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerEnterRoom struct {
	T      time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerEnterRoom) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerEnterRoom) Time() time.Time {
	return r.T
}

func (r framePlayerEnterRoom) beforeUpdate(ro Room) error {
	if ro.HasPlayer(r.Player) {
		return nil
	}
	switch v := ro.(type) {
	case *RoomPokerDoudizhu:
		if len(v.PokerPlayers) >= 3 {
			return fmt.Errorf("framePlayerEnterRoom beforeUpdate: RoomPokerDoudizhu len(v.PokerPlayers) >= 3")
		}
		break
	}
	return nil
}

func (r framePlayerEnterRoom) update(ro Room) error {
	return nil
}

func (r framePlayerEnterRoom) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":   r.Name(),
		"time":   r.T,
		"player": r.Player,
	}
	return json.Marshal(m)
}

func NewFramePlayerEnterRoom(p *player.Player) frame {
	return framePlayerEnterRoom{
		Player: p,
		T:      time.Now(),
	}
}
