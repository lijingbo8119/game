package client

import (
	"game/server"
	"game/util"
)

func requestSignup(params server.EventSignupParams) error {
	var (
		err error
	)
	if err = util.ValidateStruct(params); err != nil {
		return err
	}
	return server.Data{Cmd: server.CmdSignupRequest, Payload: params}.Send(conn)
}
