package room

import (
	"encoding/json"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerLeaveRoom struct {
	T      time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerLeaveRoom) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerLeaveRoom) Time() time.Time {
	return r.T
}

func (r framePlayerLeaveRoom) beforeUpdate(ro Room) error {
	return nil
}

func (r framePlayerLeaveRoom) update(ro Room) error {
	return nil
}

func (r framePlayerLeaveRoom) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":   r.Name(),
		"time":   r.T,
		"player": r.Player,
	}
	return json.Marshal(m)
}

func NewFramePlayerLeaveRoom(p *player.Player) frame {
	return framePlayerLeaveRoom{
		Player: p,
		T:      time.Now(),
	}
}
