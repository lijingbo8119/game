package player

import (
	"sync"

	"github.com/gofrs/uuid"
)

type PlayerBooleanClosure = func(*Player) bool

type Players []*Player

func (r Players) Length() int {
	return len(r)
}

func (r Players) Exists(closure PlayerBooleanClosure) bool {
	return r.First(closure) != nil
}

func (r Players) First(closure ...PlayerBooleanClosure) *Player {
	_closure := func() PlayerBooleanClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(c *Player) bool { return true }
	}()
	for _, c := range r {
		if _closure(c) {
			return c
		}
	}
	return nil
}

func (r Players) Last(closure ...PlayerBooleanClosure) *Player {
	_closure := func() PlayerBooleanClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(c *Player) bool { return true }
	}()
	for i := r.Length() - 1; i >= 0; i-- {
		if _closure(r[i]) {
			return r[i]
		}
	}
	return nil
}

func (r Players) Count(closure ...PlayerBooleanClosure) int {
	_closure := func() PlayerBooleanClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(c *Player) bool { return true }
	}()
	res := 0
	for i := r.Length() - 1; i >= 0; i-- {
		if _closure(r[i]) {
			res++
		}
	}
	return res
}

func (r *Players) Append(players ...*Player) {
	*r = append(*r, players...)
}

func (r *Players) Remove(player *Player) bool {
	found := false
	temp := Players{}
	for _, p := range *r {
		if p == player {
			found = true
			continue
		}
		temp = append(temp, p)
	}
	*r = temp
	return found
}

var players Players = []*Player{}
var playersMux = sync.RWMutex{}

func CreateOrGetPlayer(p *Player) *Player {
	var (
		player *Player
	)
	if len(p.Id.String()) == 0 {
		p.Id = uuid.Must(uuid.NewV4())
	}
	if player = FindPlayer(func(p2 *Player) bool {
		return p2.Id == p.Id || p2.Name == p.Name
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
