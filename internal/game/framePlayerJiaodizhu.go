package game

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"game/internal/poker"
	"reflect"
	"time"

	"github.com/samber/lo"
)

type framePlayerJiaodizhu struct {
	Time   time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerJiaodizhu) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerJiaodizhu) time() time.Time {
	return r.Time
}

func (r framePlayerJiaodizhu) beforeUpdate(g Game) error {
	gameDoudizhu, ok := g.(*GamePokerDoudizhu)
	if !ok {
		return fmt.Errorf("framePlayerJiaodizhu: beforeUpdate: Game type error")
	}
	f := gameDoudizhu.lastFrame()
	switch f.(type) {
	case framePlayerJiaodizhu:
		break
	case framePlayerRoundStart:
		break
	case framePlayerRoundEnd:
		break
	default:
		return fmt.Errorf("framePlayerJiaodizhu: beforeUpdate: frame type error")
	}
	return nil
}

func (r framePlayerJiaodizhu) update(g Game) error {
	gameDoudizhu, ok := g.(*GamePokerDoudizhu)
	if !ok {
		return fmt.Errorf("game type error")
	}
	pokerPlayer, ok := lo.Find(gameDoudizhu.PokerPlayers, func(t *poker.PokerPlayer) bool {
		return t.Player == r.Player
	})
	if !ok {
		return fmt.Errorf("framePlayerJiaodizhu find pokerPlayer error")
	}
	gameDoudizhu.PokerPlayersDizhuIndex = lo.IndexOf(gameDoudizhu.PokerPlayers, pokerPlayer)
	return nil
}

func (r framePlayerJiaodizhu) String() string {
	return ""
}

func (r framePlayerJiaodizhu) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":      r.Name(),
		"time":      r.time(),
		"player_id": r.Player.Id,
	}
	return json.Marshal(m)
}

func NewFramePlayerJiaodizhu(p *player.Player) frame {
	return framePlayerJiaodizhu{
		Player: p,
		Time:   time.Now(),
	}
}
