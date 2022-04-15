package game

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"reflect"
	"time"
)

type framePlayerRoundStart struct {
	Time   time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r framePlayerRoundStart) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r framePlayerRoundStart) time() time.Time {
	return r.Time
}

func (r framePlayerRoundStart) beforeUpdate(g Game) error {
	fStart := g.lastFrame(func(f frame) bool {
		_, ok := f.(framePlayerRoundStart)
		return ok
	})
	if fStart == nil {
		return fmt.Errorf("framePlayerRoundStart beforeUpdate: need framePlayerRoundStart")
	}
	fEnd := g.lastFrame(func(f frame) bool {
		_, ok := f.(framePlayerRoundEnd)
		return ok
	})
	if fEnd == nil {
		return fmt.Errorf("framePlayerRoundStart beforeUpdate: need framePlayerRoundEnd")
	}
	if fEnd.time().Before(fStart.time()) {
		return fmt.Errorf("framePlayerRoundStart beforeUpdate: time error: %v %v", fStart.time(), fEnd.time())
	}
	_fStart := fStart.(framePlayerRoundStart)
	_fEnd := fEnd.(framePlayerRoundEnd)
	if _fStart.Player != _fEnd.Player {
		return fmt.Errorf("framePlayerRoundStart beforeUpdate: last round player unmatched")
	}
	return nil
}

func (r framePlayerRoundStart) update(g Game) error {
	return nil
}

func (r framePlayerRoundStart) String() string {
	return ""
}

func (r framePlayerRoundStart) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":      r.Name(),
		"time":      r.time(),
		"player_id": r.Player.Id,
	}
	return json.Marshal(m)
}

func NewFrameGamePlayerRoundStart(p *player.Player) frame {
	return framePlayerRoundStart{
		Player: p,
		Time:   time.Now(),
	}
}
