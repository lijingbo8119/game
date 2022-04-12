package room

import (
	"encoding/json"
	"fmt"
	"game/internal/player"
	"reflect"
	"time"
)

type frameRoomPlayerRoundStart struct {
	T      time.Time      `json:"time"`
	Player *player.Player `json:"player"`
}

func (r frameRoomPlayerRoundStart) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r frameRoomPlayerRoundStart) Time() time.Time {
	return r.T
}

func (r frameRoomPlayerRoundStart) beforeUpdate(ro Room) error {
	fStart := ro.lastFrame(func(f frame) bool {
		_, ok := f.(frameRoomPlayerRoundStart)
		return ok
	})
	if fStart == nil {
		return fmt.Errorf("frameRoomPlayerRoundStart beforeUpdate: need frameRoomPlayerRoundStart")
	}
	fEnd := ro.lastFrame(func(f frame) bool {
		_, ok := f.(frameRoomPlayerRoundEnd)
		return ok
	})
	if fEnd == nil {
		return fmt.Errorf("frameRoomPlayerRoundStart beforeUpdate: need frameRoomPlayerRoundEnd")
	}
	if fEnd.Time().Before(fStart.Time()) {
		return fmt.Errorf("frameRoomPlayerRoundStart beforeUpdate: time error: %v %v", fStart.Time(), fEnd.Time())
	}
	_fStart := fStart.(frameRoomPlayerRoundStart)
	_fEnd := fEnd.(frameRoomPlayerRoundEnd)
	if _fStart.Player != _fEnd.Player {
		return fmt.Errorf("frameRoomPlayerRoundStart beforeUpdate: last round player unmatched")
	}
	return nil
}

func (r frameRoomPlayerRoundStart) update(ro Room) error {
	return nil
}

func (r frameRoomPlayerRoundStart) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":   r.Name(),
		"time":   r.T,
		"player": r.Player,
	}
	return json.Marshal(m)
}

func NewFrameRoomPlayerRoundStart(p *player.Player) frame {
	return frameRoomPlayerRoundStart{
		Player: p,
		T:      time.Now(),
	}
}
