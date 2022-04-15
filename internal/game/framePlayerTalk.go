package game

import (
	"encoding/json"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerTalk struct {
	Time    time.Time      `json:"time"`
	Player  *player.Player `json:"player"`
	Content string         `json:"content"`
}

func (r framePlayerTalk) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerTalk) time() time.Time {
	return r.Time
}

func (r framePlayerTalk) beforeUpdate(g Game) error {
	return nil
}

func (r framePlayerTalk) update(g Game) error {
	return nil
}

func (r framePlayerTalk) String() string {
	return ""
}

func (r framePlayerTalk) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":    r.Name(),
		"time":    r.time(),
		"player":  r.Player,
		"Content": r.Content,
	}
	return json.Marshal(m)
}

func NewFramePlayerTalk(p *player.Player, content string) frame {
	return framePlayerTalk{
		Player:  p,
		Time:    time.Now(),
		Content: content,
	}
}
