package game

import (
	"encoding/json"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerLeaveGame struct {
	Time   time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerLeaveGame) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerLeaveGame) time() time.Time {
	return r.Time
}

func (r framePlayerLeaveGame) beforeUpdate(g Game) error {
	return nil
}

func (r framePlayerLeaveGame) update(g Game) error {
	return nil
}

func (r framePlayerLeaveGame) String() string {
	return ""
}

func (r framePlayerLeaveGame) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":      r.Name(),
		"time":      r.time(),
		"player_id": r.Player.Id,
	}
	return json.Marshal(m)
}

func NewFramePlayerLeaveGame(p *player.Player) frame {
	return framePlayerLeaveGame{
		Player: p,
		Time:   time.Now(),
	}
}
