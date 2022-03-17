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
	if err = d.ParseParams(&params); err != nil {
		return err
	}
	if p = player.FindPlayer(func(p *player.Player) bool {
		return p.Nickname == params.Username
	}); p == nil {
		return Data{Cmd: CmdSigninFailedResponse, Params: "player exists already"}.Send(c)
	}
	if params.Password != p.Password {
		return Data{Cmd: CmdSigninFailedResponse, Params: "invalid password"}.Send(c)
	}
	return Data{Cmd: CmdSigninSucceedResponse, Params: *p}.Send(c)
}
