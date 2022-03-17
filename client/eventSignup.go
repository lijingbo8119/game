package client

import (
	"fmt"
	"game/server"

	"github.com/gorilla/websocket"
)

func eventSignupFailed(c *websocket.Conn, d server.Data) error {
	var (
		params = server.EventSigninParams{}
		err    error
	)
	if err = d.ParseParams(&params); err != nil {
		return err
	}
	return server.Data{Cmd: server.CmdSigninRequest, Params: params}.Send(conn)
}

func eventSignupSucceed(c *websocket.Conn, d server.Data) error {
	var (
		params = server.EventSigninParams{}
		err    error
	)
	if err = d.ParseParams(&params); err != nil {
		return err
	}
	if v := activeView(); v != nil && v.Name() == (viewSignup{}).Name() {
		goToView(viewHall{}.Name())
	}
	return fmt.Errorf("eventSignupSucceed failed")
}
