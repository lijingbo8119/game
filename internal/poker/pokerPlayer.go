package poker

import "game/internal/player"

type PokerPlayer struct {
	Player *player.Player
	Cards  Cards
}
