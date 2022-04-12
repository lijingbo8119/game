package client

import (
	"fmt"
	"game/server"

	"github.com/gorilla/websocket"
)

func eventSigninFailed(c *websocket.Conn, d server.Data) error {
	var (
		params string
		err    error
	)
	if err = d.ParsePayload(&params); err != nil {
		return err
	}
	fmt.Println(d)
	return nil
}

func eventSigninSucceed(c *websocket.Conn, d server.Data) error {
	var (
		params string
		err    error
	)
	if err = d.ParsePayload(&params); err != nil {
		return err
	}
	if v := getViewModel(); v != nil && v.Name() == (viewModelSignin{}.Name()) {
		getViewModel(viewModelHall{}.Name())
	}
	return fmt.Errorf("eventSigninSucceed failed")
}
