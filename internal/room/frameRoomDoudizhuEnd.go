package room

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type frameRoomDoudizhuEnd struct {
	T time.Time `json:"time"`
}

func (r frameRoomDoudizhuEnd) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r frameRoomDoudizhuEnd) Time() time.Time {
	return r.T
}

func (r frameRoomDoudizhuEnd) beforeUpdate(ro Room) error {
	roomDoudizhu, ok := ro.(*RoomPokerDoudizhu)
	if !ok {
		return fmt.Errorf("frameRoomDoudizhuEnd beforeUpdate: type error")
	}
	f := roomDoudizhu.lastFrame(func(f frame) bool {
		_, ok := f.(framePlayerJiaodizhu)
		return ok
	})
	if f == nil {
		return fmt.Errorf("frameRoomDoudizhuEnd beforeUpdate: need framePlayerJiaodizhu")
	}
	if roomDoudizhu.PokerPlayers[roomDoudizhu.PokerPlayersDizhuIndex].Cards.Length() == 0 {
		return nil
	}
	for i, p := range roomDoudizhu.PokerPlayers {
		if i == roomDoudizhu.PokerPlayersDizhuIndex {
			continue
		}
		if p.Cards.Length() > 0 {
			return fmt.Errorf("frameRoomDoudizhuEnd beforeUpdate: failed")
		}
	}
	return nil
}

func (r frameRoomDoudizhuEnd) update(ro Room) error {
	return nil
}

func (r frameRoomDoudizhuEnd) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name": r.Name(),
		"time": r.T,
	}
	return json.Marshal(m)
}

func NewFrameRoomDoudizhuEnd() frame {
	return frameRoomDoudizhuEnd{
		T: time.Now(),
	}
}
