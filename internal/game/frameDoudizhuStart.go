package game

import (
	"encoding/json"
	"fmt"
	"game/internal/poker"
	"reflect"
	"time"

	"github.com/samber/lo"
)

type frameDoudizhuStart struct {
	Time time.Time `json:"time"`
}

func (r frameDoudizhuStart) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r frameDoudizhuStart) time() time.Time {
	return r.Time
}

func (r frameDoudizhuStart) beforeUpdate(g Game) error {
	gameDoudizhu, ok := g.(*GamePokerDoudizhu)
	if !ok {
		return fmt.Errorf("frameDoudizhuStart beforeUpdate: type error")
	}
	f := gameDoudizhu.lastFrame(func(f frame) bool {
		_, ok := f.(frameDoudizhuStart)
		return ok
	})
	if f != nil {
		return fmt.Errorf("frameDoudizhuStart beforeUpdate: has frameDoudizhuStart")
	}
	f = gameDoudizhu.lastFrame(func(f frame) bool {
		_, ok := f.(framePlayerJiaodizhu)
		return ok
	})
	if f == nil {
		return fmt.Errorf("frameDoudizhuStart beforeUpdate: need framePlayerJiaodizhu")
	}
	return nil
}

func (r frameDoudizhuStart) update(g Game) error {
	gameDoudizhu, ok := g.(*GamePokerDoudizhu)
	if !ok {
		return fmt.Errorf("game type error")
	}
	cards := poker.NewDeckCards()
	cards.Shuffle()
	for _, p := range gameDoudizhu.PokerPlayers {
		lo.Times(17, func(i int) bool {
			p.Cards.Append(cards.Pop())
			return true
		})
	}
	gameDoudizhu.LeftoverCards = cards
	return nil
}

func (r frameDoudizhuStart) String() string {
	return ""
}

func (r frameDoudizhuStart) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name": r.Name(),
		"time": r.time(),
	}
	return json.Marshal(m)
}

func NewFrameGamePokerDoudizhuStart() frame {
	return frameDoudizhuStart{
		Time: time.Now(),
	}
}
