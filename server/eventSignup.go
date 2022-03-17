package server

import (
	"game/internal/player"

	"github.com/gorilla/websocket"
)

type EventSignupParams struct {
	Username string
	Password string
	Nickname string
}

func eventSignup(c *websocket.Conn, d Data) error {
	var (
		params = EventSignupParams{}
		p      *player.Player
		err    error
	)
	if err = d.ParseParams(&params); err != nil {
		return err
	}
	if p = player.FindPlayer(func(p *player.Player) bool {
		return p.Nickname == params.Username
	}); p != nil {
		return Data{Cmd: CmdSignupFailedResponse, Params: "player exists already"}.Send(c)
	}
	p = &player.Player{
		Username: params.Username,
		Password: params.Password,
		Nickname: params.Nickname,
	}
	player.CreateOrGetPlayer(p)
	return Data{Cmd: CmdSignupSucceedResponse}.Send(c)
}
