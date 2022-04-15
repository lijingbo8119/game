package game

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type frameDoudizhuEnd struct {
	Time time.Time `json:"time"`
}

func (r frameDoudizhuEnd) d() string {
	return reflect.TypeOf(r).Name()
}

func (r frameDoudizhuEnd) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r frameDoudizhuEnd) time() time.Time {
	return r.Time
}

func (r frameDoudizhuEnd) beforeUpdate(g Game) error {
	gameDoudizhu, ok := g.(*GamePokerDoudizhu)
	if !ok {
		return fmt.Errorf("frameDoudizhuEnd beforeUpdate: type error")
	}
	f := gameDoudizhu.lastFrame(func(f frame) bool {
		_, ok := f.(framePlayerJiaodizhu)
		return ok
	})
	if f == nil {
		return fmt.Errorf("frameDoudizhuEnd beforeUpdate: need framePlayerJiaodizhu")
	}
	if gameDoudizhu.PokerPlayers[gameDoudizhu.PokerPlayersDizhuIndex].Cards.Length() == 0 {
		return nil
	}
	for i, p := range gameDoudizhu.PokerPlayers {
		if i == gameDoudizhu.PokerPlayersDizhuIndex {
			continue
		}
		if p.Cards.Length() > 0 {
			return fmt.Errorf("frameDoudizhuEnd beforeUpdate: failed")
		}
	}
	return nil
}

func (r frameDoudizhuEnd) update(g Game) error {
	return nil
}

func (r frameDoudizhuEnd) String() string {
	return ""
}

func (r frameDoudizhuEnd) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name": r.Name(),
		"time": r.time(),
	}
	return json.Marshal(m)
}

func NewFrameGamePokerDoudizhuEnd() frame {
	return frameDoudizhuEnd{
		Time: time.Now(),
	}
}
