package room

import (
	"game/internal/player"
	"sync"

	"github.com/gofrs/uuid"
)

var rooms = []Room{}
var roomsMux = sync.RWMutex{}

func CreateOrGetRoom(r Room) Room {
	var (
		room Room
	)
	if room = FindRoom(func(r2 Room) bool {
		return r2.Id() == r.Id() || r2.Name() == r.Name()
	}); room != nil {
		return room
	}
	roomsMux.Lock()
	defer roomsMux.Unlock()
	rooms = append(rooms, r)
	return r
}

func GetRooms() []Room {
	roomsMux.RLock()
	defer roomsMux.RUnlock()
	return rooms
}

func FindRoom(closure func(r Room) bool) Room {
	roomsMux.RLock()
	defer roomsMux.RUnlock()
	for _, r := range rooms {
		if closure(r) {
			return r
		}
	}
	return nil
}

func FindRoomById(id uuid.UUID) Room {
	return FindRoom(func(r Room) bool {
		return r.Id() == id
	})
}

func FindRoomByPlayer(closure func(p *player.Player) bool) Room {
	return FindRoom(func(r Room) bool {
		for _, p := range r.Players() {
			if closure(p) {
				return true
			}
		}
		return false
	})
}

func FindRoomByPlayerId(pid uuid.UUID) Room {
	return FindRoomByPlayer(func(p *player.Player) bool {
		return pid == p.Id
	})
}
