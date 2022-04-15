package server

import (
	"github.com/gorilla/websocket"
)

type EventClosure func(*websocket.Conn, Data) error

var events = map[NetCmd]EventClosure{
	NetCmdSignupRequest:      eventSignup,
	NetCmdSigninRequest:      eventSignin,
	NetCmdGetGameListRequest: eventGetGameList,
	NetCmdGetGameRequest:     eventGetGame,
	NetCmdEnterGameRequest:   eventEnterGame,
}
