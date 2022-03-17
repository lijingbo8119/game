package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func websocketServer(w http.ResponseWriter, r *http.Request) {
	var (
		upgrader    = websocket.Upgrader{} // use default options
		conn        *websocket.Conn
		messageType int
		payload     []byte
		data        *Data
		err         error
	)
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		panic(err)
	}

	conn.SetPingHandler(func(appData string) error {
		return conn.WriteMessage(websocket.PongMessage, []byte("pong"))
	})
	conn.SetCloseHandler(func(code int, text string) error {
		defer conn.Close()
		return conn.WriteMessage(websocket.PongMessage, []byte("bye"))
	})

	defer conn.Close()
	for {
		if messageType, payload, err = conn.ReadMessage(); err != nil {
			panic(err)
		}
		switch messageType {
		case websocket.TextMessage:
			if data, err = (&Data{}).Parse(payload); err != nil {
				panic(err)
			}
			if err = events[data.Cmd](conn, *data); err != nil {
				panic(err)
			}
		}
	}
}
