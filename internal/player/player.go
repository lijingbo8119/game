package player

import (
	"sync"

	"github.com/gofrs/uuid"
)

type Player struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"-"`
	Nickname string    `json:"nickname"`
}

type PlayerBooleanClosure = func(*Player) bool

var players = []*Player{}
var playersMux = sync.RWMutex{}

func CreateOrGetPlayer(p *Player) *Player {
	var (
		player *Player
	)
	if len(p.Id.String()) == 0 {
		p.Id = uuid.Must(uuid.NewV4())
	}
	if player = FindPlayer(func(p2 *Player) bool {
		return p2.Id == p.Id || p2.Nickname == p.Nickname
	}); player != nil {
		return player
	}
	playersMux.Lock()
	defer playersMux.Unlock()
	players = append(players, p)
	return p
}

func FindPlayer(closure func(p *Player) bool) *Player {
	playersMux.RLock()
	defer playersMux.RUnlock()
	for _, p := range players {
		if closure(p) {
			return p
		}
	}
	return nil
}

func FindPlayerById(id uuid.UUID) *Player {
	return FindPlayer(func(p *Player) bool {
		return p.Id == id
	})
}
