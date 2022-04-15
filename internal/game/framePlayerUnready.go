package game

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerUnready struct {
	Time   time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerUnready) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerUnready) time() time.Time {
	return r.Time
}

func (r framePlayerUnready) beforeUpdate(g Game) error {
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

func (r framePlayerUnready) update(g Game) error {
	return nil
}

func (r framePlayerUnready) String() string {
	return ""
}

func (r framePlayerUnready) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":      r.Name(),
		"time":      r.time(),
		"player_id": r.Player.Id,
	}
	return json.Marshal(m)
}

func NewFramePlayerUnready(p *player.Player) frame {
	return framePlayerUnready{
		Player: p,
		Time:   time.Now(),
	}
}
