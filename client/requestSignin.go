package client

import (
	"game/server"
)

func requestSignin(d server.Data) error {
	var (
		params = server.EventSigninParams{}
		err    error
	)
	if err = d.ParsePayload(&params); err != nil {
		return err
	}
	return server.Data{Cmd: server.CmdSigninRequest, Payload: params}.Send(conn)
}
