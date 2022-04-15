package doudizhu

import (
	"fmt"
	"game/internal/player"
	"sync"
)

const GameName = "doudizhu"

type Game struct {
	mux     sync.RWMutex
	players []*player.Player
}

func (r *Game) Name() string {
	return GameName
}

func (r *Game) GetPlayers() []*player.Player {
	r.mux.RLock()
	defer r.mux.RUnlock()
	return r.players
}

func (r *Game) Enter(p *player.Player) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if len(r.players) == 3 {
		return fmt.Errorf("Enter: It's full")
	}
	r.players = append(r.players, p)
	return nil
}
