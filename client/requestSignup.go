package client

import (
	"game/server"
)

func requestSignup(d server.Data) error {
	var (
		params = server.EventSignupParams{}
		err    error
	)
	if err = d.ParseParams(&params); err != nil {
		return err
	}
	return server.Data{Cmd: server.CmdSignupRequest, Params: params}.Send(conn)
}
