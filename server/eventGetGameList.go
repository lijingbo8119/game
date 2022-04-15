package server

import (
	"game/internal/game"

	"github.com/gorilla/websocket"
)

func eventGetGameList(c *websocket.Conn, d Data) error {
	return Data{NetCmd: NetCmdGetGameListSucceedResponse, Payload: game.GetGames()}.Send(c)
}
