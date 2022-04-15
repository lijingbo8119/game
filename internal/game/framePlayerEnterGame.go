package game

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerEnterGame struct {
	Time   time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerEnterGame) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerEnterGame) time() time.Time {
	return r.Time
}

func (r framePlayerEnterGame) beforeUpdate(g Game) error {
	if g.HasPlayer(r.Player) {
		return nil
	}
	switch v := g.(type) {
	case *GamePokerDoudizhu:
		if len(v.PokerPlayers) >= 3 {
			return fmt.Errorf("framePlayerEnterGame beforeUpdate: GamePokerDoudizhu len(v.PokerPlayers) >= 3")
		}
		break
	}
	return nil
}

func (r framePlayerEnterGame) update(g Game) error {
	return nil
}

func (r framePlayerEnterGame) String() string {
	return ""
}

func (r framePlayerEnterGame) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":   r.Name(),
		"time":   r.time(),
		"player": r.Player,
	}
	return json.Marshal(m)
}

func NewFramePlayerEnterGame(p *player.Player) frame {
	return framePlayerEnterGame{
		Player: p,
		Time:   time.Now(),
	}
}
