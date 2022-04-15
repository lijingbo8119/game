package game

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerRoundEnd struct {
	Time   time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerRoundEnd) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerRoundEnd) time() time.Time {
	return r.Time
}

func (r framePlayerRoundEnd) beforeUpdate(g Game) error {
	fStart := g.lastFrame(func(f frame) bool {
		_, ok := f.(framePlayerRoundStart)
		return ok
	})
	if fStart == nil {
		return fmt.Errorf("framePlayerRoundEnd beforeUpdate: need framePlayerRoundStart")
	}
	_fStart := fStart.(framePlayerRoundStart)
	if _fStart.Player != r.Player {
		return fmt.Errorf("framePlayerRoundEnd beforeUpdate: player unmatched")
	}
	fEnd := g.lastFrame(func(f frame) bool {
		_, ok := f.(framePlayerRoundEnd)
		return ok
	})
	if fEnd != nil && fEnd.time().After(fStart.time()) {
		return fmt.Errorf("framePlayerRoundEnd beforeUpdate: time error: %v %v", fStart.time(), fEnd.time())
	}
	return nil
}

func (r framePlayerRoundEnd) update(g Game) error {
	return nil
}

func (r framePlayerRoundEnd) String() string {
	return ""
}

func (r framePlayerRoundEnd) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":      r.Name(),
		"time":      r.time(),
		"player_id": r.Player.Id,
	}
	return json.Marshal(m)
}

func NewFrameGamePlayerRoundEnd(p *player.Player) frame {
	return framePlayerRoundEnd{
		Player: p,
		Time:   time.Now(),
	}
}
