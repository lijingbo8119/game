package room

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"reflect"
	"time"
)

type frameRoomPlayerRoundEnd struct {
	T      time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r frameRoomPlayerRoundEnd) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r frameRoomPlayerRoundEnd) Time() time.Time {
	return r.T
}

func (r frameRoomPlayerRoundEnd) beforeUpdate(ro Room) error {
	fStart := ro.lastFrame(func(f frame) bool {
		_, ok := f.(frameRoomPlayerRoundStart)
		return ok
	})
	if fStart == nil {
		return fmt.Errorf("frameRoomPlayerRoundEnd beforeUpdate: need frameRoomPlayerRoundStart")
	}
	_fStart := fStart.(frameRoomPlayerRoundStart)
	if _fStart.Player != r.Player {
		return fmt.Errorf("frameRoomPlayerRoundEnd beforeUpdate: player unmatched")
	}
	fEnd := ro.lastFrame(func(f frame) bool {
		_, ok := f.(frameRoomPlayerRoundEnd)
		return ok
	})
	if fEnd != nil && fEnd.Time().After(fStart.Time()) {
		return fmt.Errorf("frameRoomPlayerRoundEnd beforeUpdate: time error: %v %v", fStart.Time(), fEnd.Time())
	}
	return nil
}

func (r frameRoomPlayerRoundEnd) update(ro Room) error {
	return nil
}

func (r frameRoomPlayerRoundEnd) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":   r.Name(),
		"time":   r.T,
		"player": r.Player,
	}
	return json.Marshal(m)
}

func NewFrameRoomPlayerRoundEnd(p *player.Player) frame {
	return frameRoomPlayerRoundEnd{
		Player: p,
		T:      time.Now(),
	}
}
