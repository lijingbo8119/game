package server

import (
	"game/internal/player"

	"github.com/gorilla/websocket"
)

type EventSignupParams struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
	Nickname string `validate:"required"`
}

func eventSignup(c *websocket.Conn, d Data) error {
	var (
		params = EventSignupParams{}
		p      *player.Player
		err    error
	)
	if err = d.ParsePayload(&params); err != nil {
		return err
	}
	if p = player.FindPlayer(func(p *player.Player) bool {
		return p.Username == params.Username
	}); p != nil {
		return Data{NetCmd: NetCmdSignupFailedResponse, Payload: "Username already exists"}.Send(c)
	}
	if p = player.FindPlayer(func(p *player.Player) bool {
		return p.Nickname == params.Nickname
	}); p != nil {
		return Data{NetCmd: NetCmdSignupFailedResponse, Payload: "Nickname already exists"}.Send(c)
	}
	p = &player.Player{
		Username: params.Username,
		Password: params.Password,
		Nickname: params.Nickname,
	}
	player.CreateOrGetPlayer(p)
	return Data{NetCmd: NetCmdSignupSucceedResponse, Payload: p}.Send(c)
}
