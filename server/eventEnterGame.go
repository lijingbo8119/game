package server

import (
	"game/internal/game"
	"game/internal/player"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type EventEnterGameParams struct {
	GameId   uuid.UUID
	PlayerId uuid.UUID
}

func eventEnterGame(c *websocket.Conn, d Data) error {
	var (
		params = EventEnterGameParams{}
		g      game.Game
		p      *player.Player
		err    error
	)
	if err = d.ParsePayload(&params); err != nil {
		return err
	}
	g = game.FindGame(func(r game.Game) bool {
		return r.Id() == params.GameId
	})
	p = player.FindPlayer(func(p *player.Player) bool {
		return p.Id == params.PlayerId
	})
	if g == nil || p == nil {
		return Data{NetCmd: NetCmdEnterGameFailedResponse, Payload: "data not exist"}.Send(c)
	}
	if err = g.Enter(p); err != nil {
		return Data{NetCmd: NetCmdEnterGameFailedResponse, Payload: err.Error()}.Send(c)
	}
	return Data{NetCmd: NetCmdEnterGameSucceedResponse, Payload: g}.Send(c)
}
