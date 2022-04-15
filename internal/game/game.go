package game

import (
	"game/internal/player"
	"sync"

	"github.com/gofrs/uuid"
)

type Game interface {
	Id() uuid.UUID
	Name() string
	Players() []*player.Player
	HasPlayer(p *player.Player) bool
	Enter(*player.Player) error
	Leave(*player.Player) error
	lastFrame(closure ...func(f frame) bool) frame
	AppendFrame(f frame) error
}

var (
	gameStructs = []Game{
		&GameHall{},
		&GamePokerDoudizhu{},
	}
	games    = []Game{}
	gamesMux = sync.RWMutex{}
)

func CreateOrGetGame(r Game) Game {
	var (
		game Game
	)
	if game = FindGame(func(r2 Game) bool {
		return r2.Id() == r.Id() || r2.Name() == r.Name()
	}); game != nil {
		return game
	}
	gamesMux.Lock()
	defer gamesMux.Unlock()
	games = append(games, r)
	return r
}

func GetGames() []Game {
	gamesMux.RLock()
	defer gamesMux.RUnlock()
	return games
}

func FindGame(closure func(r Game) bool) Game {
	gamesMux.RLock()
	defer gamesMux.RUnlock()
	for _, r := range games {
		if closure(r) {
			return r
		}
	}
	return nil
}

func FindGameById(id uuid.UUID) Game {
	return FindGame(func(r Game) bool {
		return r.Id() == id
	})
}

func FindGameByPlayer(closure func(p *player.Player) bool) Game {
	return FindGame(func(r Game) bool {
		for _, p := range r.Players() {
			if closure(p) {
				return true
			}
		}
		return false
	})
}

func FindGameByPlayerId(pid uuid.UUID) Game {
	return FindGameByPlayer(func(p *player.Player) bool {
		return pid == p.Id
	})
}
