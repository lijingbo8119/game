package client

import "game/server"

type requestClosure func(server.Data) error

var requests = map[server.DataCmd]requestClosure{
	server.CmdSignupRequest: requestSignup,
	server.CmdSigninRequest: requestSignin,
}
