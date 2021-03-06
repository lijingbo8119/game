package client

import (
	"game/server"
	"game/util"
)

func requestSignin(params server.EventSigninParams) error {
	var (
		err error
	)
	if err = util.ValidateStruct(params); err != nil {
		return err
	}
	return server.Data{NetCmd: server.NetCmdSigninRequest, Payload: params}.Send(conn)
}
