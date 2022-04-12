package room

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerUnready struct {
	T      time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerUnready) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerUnready) Time() time.Time {
	return r.T
}

func (r framePlayerUnready) beforeUpdate(ro Room) error {
	f := ro.lastFrame()
	if f == nil {
		return nil
	}
	switch f.(type) {
	case framePlayerReady:
		break
	case framePlayerUnready:
		break
	default:
		return fmt.Errorf("framePlayerReady beforeUpdate: type error")
	}
	return nil
}

func (r framePlayerUnready) update(ro Room) error {
	return nil
}

func (r framePlayerUnready) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":   r.Name(),
		"time":   r.T,
		"player": r.Player,
	}
	return json.Marshal(m)
}

func NewFramePlayerUnready(p *player.Player) frame {
	return framePlayerUnready{
		Player: p,
		T:      time.Now(),
	}
}
