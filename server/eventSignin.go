package server

import (
	"game/internal/player"

	"github.com/gorilla/websocket"
)

type EventSigninParams struct {
	Username string
	Password string
}

func eventSignin(c *websocket.Conn, d Data) error {
	var (
		params = EventSigninParams{}
		p      *player.Player
		err    error
	)
	if err = d.ParsePayload(&params); err != nil {
		return err
	}
	if p = player.FindPlayer(func(p *player.Player) bool {
		return p.Username == params.Username
	}); p == nil {
		return Data{NetCmd: NetCmdSigninFailedResponse, Payload: "player exists already"}.Send(c)
	}
	if params.Password != p.Password {
		return Data{NetCmd: NetCmdSigninFailedResponse, Payload: "invalid password"}.Send(c)
	}
	return Data{NetCmd: NetCmdSigninSucceedResponse, Payload: *p}.Send(c)
}
