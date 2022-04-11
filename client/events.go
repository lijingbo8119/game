package client

import "game/server"

var events = map[server.NetCmd]server.EventClosure{
	server.NetCmdSignupFailedResponse:  eventSignupFailed,
	server.NetCmdSignupSucceedResponse: eventSignupSucceed,
	server.NetCmdSigninFailedResponse:  eventSigninFailed,
	server.NetCmdSigninSucceedResponse: eventSigninFailed,
}
