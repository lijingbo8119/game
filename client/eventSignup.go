package client

import (
	"fmt"
	"game/internal/player"
	"game/server"

	"github.com/gorilla/websocket"
)

func eventSignupFailed(c *websocket.Conn, d server.Data) error {
	signupModel := currentViewModel(viewModelSignup{}.Name()).(*viewModelSignup)
	signupModel.setStatusBarContent(d.Payload.(string))
	return nil
}

func eventSignupSucceed(c *websocket.Conn, d server.Data) error {
	var (
		player = player.Player{}
		err    error
	)
	if err = d.ParsePayload(&player); err != nil {
		return err
	}
	setPlayer(&player)
	if v := currentViewModel(); v != nil && v.Name() == (viewModelSignup{}.Name()) {
		currentViewModel(viewModelSignin{}.Name())
		return nil
	}
	return fmt.Errorf("eventSignupSucceed failed")
}
