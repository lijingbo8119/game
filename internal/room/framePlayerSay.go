package room

import (
	"encoding/json"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerSay struct {
	T       time.Time      `json:"time"`
	Player  *player.Player `json:"player"`
	Content string
}

func (r framePlayerSay) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerSay) Time() time.Time {
	return r.T
}

func (r framePlayerSay) beforeUpdate(ro Room) error {
	return nil
}

func (r framePlayerSay) update(ro Room) error {
	return nil
}

func (r framePlayerSay) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":    r.Name(),
		"time":    r.T,
		"player":  r.Player,
		"Content": r.Content,
	}
	return json.Marshal(m)
}

func NewFramePlayerSay(p *player.Player, content string) frame {
	return framePlayerSay{
		Player:  p,
		T:       time.Now(),
		Content: content,
	}
}
