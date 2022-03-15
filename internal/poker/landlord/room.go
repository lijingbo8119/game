package landlord

import (
	"fmt"
	"game/internal"
	"sync"
)

const RoomName = "landlord"

type Room struct {
	mux     sync.RWMutex
	players []*internal.Player
}

func (r *Room) Name() string {
	return RoomName
}

func (r *Room) GetPlayers() []*internal.Player {
	r.mux.RLock()
	defer r.mux.RUnlock()
	return r.players
}

func (r *Room) Enter(p *internal.Player) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if len(r.players) == 3 {
		return fmt.Errorf("Enter: It's full")
	}
	r.players = append(r.players, p)
	return nil
}
