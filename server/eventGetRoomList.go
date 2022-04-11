package server

import (
	"game/internal/room"

	"github.com/gorilla/websocket"
)

func eventGetRoomList(c *websocket.Conn, d Data) error {
	return Data{NetCmd: NetCmdGetRoomListSucceedResponse, Payload: room.GetRooms()}.Send(c)
}
