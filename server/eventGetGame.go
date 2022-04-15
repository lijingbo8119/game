package server

import (
	"game/internal/game"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type EventGetGameParams struct {
	GameId uuid.UUID
}

func eventGetGame(c *websocket.Conn, d Data) error {
	var (
		params = EventGetGameParams{}
		g      game.Game
		err    error
	)
	if err = d.ParsePayload(&params); err != nil {
		return err
	}
	if g = game.FindGame(func(r game.Game) bool {
		return r.Id() == params.GameId
	}); g == nil {
		return Data{NetCmd: NetCmdGetGameFailedResponse, Payload: "game is not exist"}.Send(c)
	}
	return Data{NetCmd: NetCmdGetGameSucceedResponse, Payload: g}.Send(c)
}
