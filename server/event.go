package server

import (
	"github.com/gorilla/websocket"
)

type EventClosure func(*websocket.Conn, Data) error

var events = map[DataCmd]EventClosure{
	CmdSignupRequest: eventSignup,
	CmdSigninRequest: eventSignin,
}
