package room

import (
	"encoding/json"
	"fmt"
	"game/internal/poker"
	"reflect"
	"time"

	"github.com/samber/lo"
)

type frameRoomDoudizhuStart struct {
	T time.Time `json:"time"`
}

func (r frameRoomDoudizhuStart) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r frameRoomDoudizhuStart) Time() time.Time {
	return r.T
}

func (r frameRoomDoudizhuStart) beforeUpdate(ro Room) error {
	roomDoudizhu, ok := ro.(*RoomPokerDoudizhu)
	if !ok {
		return fmt.Errorf("frameRoomDoudizhuStart beforeUpdate: type error")
	}
	f := roomDoudizhu.lastFrame(func(f frame) bool {
		_, ok := f.(frameRoomDoudizhuStart)
		return ok
	})
	if f != nil {
		return fmt.Errorf("frameRoomDoudizhuStart beforeUpdate: has frameRoomDoudizhuStart")
	}
	f = roomDoudizhu.lastFrame(func(f frame) bool {
		_, ok := f.(framePlayerJiaodizhu)
		return ok
	})
	if f == nil {
		return fmt.Errorf("frameRoomDoudizhuStart beforeUpdate: need framePlayerJiaodizhu")
	}
	return nil
}

func (r frameRoomDoudizhuStart) update(ro Room) error {
	roomDoudizhu, ok := ro.(*RoomPokerDoudizhu)
	if !ok {
		return fmt.Errorf("room type error")
	}
	cards := poker.NewDeckCards()
	cards.Shuffle()
	for _, p := range roomDoudizhu.PokerPlayers {
		lo.Times(17, func(i int) bool {
			p.Cards.Append(cards.Pop())
			return true
		})
	}
	roomDoudizhu.LeftoverCards = cards
	return nil
}

func (r frameRoomDoudizhuStart) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name": r.Name(),
		"time": r.T,
	}
	return json.Marshal(m)
}

func NewFrameRoomDoudizhuStart() frame {
	return frameRoomDoudizhuStart{
		T: time.Now(),
	}
}
