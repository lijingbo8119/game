package client

import "game/internal/player"

var currentPlayer *player.Player

func getPlayer() *player.Player {
	return currentPlayer
}

func setPlayer(p *player.Player) {
	currentPlayer = p
}
