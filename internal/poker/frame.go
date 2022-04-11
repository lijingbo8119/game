package poker

import (
	"game/internal/player"
	"time"
)

type Frame struct {
	Cmd FrameCmd
	p   *player.Player
	t   *time.Time
	c   Cards
}
