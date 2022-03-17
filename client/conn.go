package client

import (
	"net/url"

	"github.com/gorilla/websocket"
)

var conn *websocket.Conn

func connInit(addr string) {
	var err error
	u := url.URL{Scheme: "ws", Host: addr, Path: "/server"}
	if conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil); err != nil {
		panic(err)
	}
}
