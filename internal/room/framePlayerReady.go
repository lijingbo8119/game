package room

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerReady struct {
	T      time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerReady) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerReady) Time() time.Time {
	return r.T
}

func (r framePlayerReady) beforeUpdate(ro Room) error {
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

func (r framePlayerReady) update(ro Room) error {
	return nil
}

func (r framePlayerReady) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":   r.Name(),
		"time":   r.T,
		"player": r.Player,
	}
	return json.Marshal(m)
}

func NewFramePlayerReady(p *player.Player) frame {
	return framePlayerReady{
		Player: p,
		T:      time.Now(),
	}
}
