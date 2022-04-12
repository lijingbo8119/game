package room

import (
	"encoding/json"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerTalk struct {
	T       time.Time      `json:"time"`
	Player  *player.Player `json:"player"`
	Content string         `json:"content"`
}

func (r framePlayerTalk) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerTalk) Time() time.Time {
	return r.T
}

func (r framePlayerTalk) beforeUpdate(ro Room) error {
	return nil
}

func (r framePlayerTalk) update(ro Room) error {
	return nil
}

func (r framePlayerTalk) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":    r.Name(),
		"time":    r.T,
		"player":  r.Player,
		"Content": r.Content,
	}
	return json.Marshal(m)
}

func NewFramePlayerSay(p *player.Player, content string) frame {
	return framePlayerTalk{
		Player:  p,
		T:       time.Now(),
		Content: content,
	}
}
