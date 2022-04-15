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

type framePlayerPlayCards struct {
	Time   time.Time      `json:"time"`
	Player *player.Player `json:"player"`
	Cards  poker.Cards    `json:"cards"`
}

func (r framePlayerPlayCards) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerPlayCards) time() time.Time {
	return r.Time
}

func (r framePlayerPlayCards) beforeUpdate(g Game) error {
	f1 := g.lastFrame(func(f frame) bool {
		return f.Name() == framePlayerRoundStart{}.Name()
	})
	if f1 == nil || f1.(framePlayerRoundStart).Player != r.Player {
		return fmt.Errorf("framePlayerPlayCards beforeUpdate error: last framePlayerRoundStart error: %v", r)
	}
	f2 := g.lastFrame(func(f frame) bool {
		_f, ok := f.(framePlayerRoundEnd)
		if !ok {
			return false
		}
		if _f.Player == r.Player && _f.time().After(f1.time()) {
			return true
		}
		return false
	})
	if f2 != nil {
		return fmt.Errorf("framePlayerPlayCards beforeUpdate error: player round ended: %v", r)
	}
	return nil
}

func (r framePlayerPlayCards) update(g Game) error {
	gameDoudizhu, ok := g.(*GamePokerDoudizhu)
	if !ok {
		return fmt.Errorf("game type error")
	}
	pokerPlayer, ok := lo.Find(gameDoudizhu.PokerPlayers, func(t *poker.PokerPlayer) bool {
		return t.Player == r.Player
	})
	if !ok {
		return fmt.Errorf("find pokerPlayer error")
	}
	pokerPlayer.Cards.Remove(r.Cards...)
	return nil
}

func (r framePlayerPlayCards) String() string {
	return ""
}

func (r framePlayerPlayCards) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":      r.Name(),
		"time":      r.time(),
		"player_id": r.Player.Id,
		"cards":     r.Cards,
	}
	return json.Marshal(m)
}

func NewFrameGamePokerDoudizhuPlayerPlay(p *player.Player, cards poker.Cards) frame {
	return framePlayerPlayCards{
		Player: p,
		Time:   time.Now(),
		Cards:  cards,
	}
}
