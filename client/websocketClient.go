package client

import (
	"game/server"
	"game/util"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func websocketClient() {
	var (
		messageType int
		rawMessage  []byte
		data        *server.Data
		err         error
	)
	c := conn
	defer c.Close()
	for {
		if messageType, rawMessage, err = c.ReadMessage(); err != nil {
			panic(err)
		}
		util.LogInfo("ReadMessage", zap.Any("messageType", messageType), zap.ByteString("rawMessage", rawMessage))
		switch messageType {
		case websocket.TextMessage:
			if data, err = (&server.Data{}).Parse(rawMessage); err != nil {
				panic(err)
			}
			if event, ok := events[data.NetCmd]; ok {
				util.LogInfo("event", zap.Any("cmd", data.NetCmd), zap.Any("data", data))
				if err = event(c, *data); err != nil {
					panic(err)
				}
			}
		}
	}
}
