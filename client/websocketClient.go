package client

import (
	"game/server"

	"github.com/gorilla/websocket"
)

func websocketClient() {
	var (
		messageType int
		payload     []byte
		data        *server.Data
		err         error
	)
	c := conn
	defer c.Close()
	for {
		if messageType, payload, err = c.ReadMessage(); err != nil {
			panic(err)
		}
		switch messageType {
		case websocket.TextMessage:
			if data, err = (&server.Data{}).Parse(payload); err != nil {
				panic(err)
			}
			if event, ok := events[data.Cmd]; ok {
				if err = event(c, *data); err != nil {
					panic(err)
				}
			}
		}
	}
}
