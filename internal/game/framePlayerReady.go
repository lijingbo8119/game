package game

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerReady struct {
	Time   time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerReady) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerReady) time() time.Time {
	return r.Time
}

func (r framePlayerReady) beforeUpdate(g Game) error {
	f := g.lastFrame()
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

func (r framePlayerReady) update(g Game) error {
	return nil
}

func (r framePlayerReady) String() string {
	return ""
}

func (r framePlayerReady) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":      r.Name(),
		"time":      r.time(),
		"player_id": r.Player.Id,
	}
	return json.Marshal(m)
}

func NewFramePlayerReady(p *player.Player) frame {
	return framePlayerReady{
		Player: p,
		Time:   time.Now(),
	}
}
