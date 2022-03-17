package client

import "game/server"

var events = map[server.DataCmd]server.EventClosure{
	server.CmdSignupFailedResponse:  eventSignupFailed,
	server.CmdSignupSucceedResponse: eventSignupSucceed,
	server.CmdSigninFailedResponse:  eventSigninFailed,
	server.CmdSigninSucceedResponse: eventSigninFailed,
}
