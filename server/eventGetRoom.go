package server

import (
	"game/internal/room"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type EventGetRoomParams struct {
	RoomId uuid.UUID
}

func eventGetRoom(c *websocket.Conn, d Data) error {
	var (
		params = EventGetRoomParams{}
		ro     room.Room
		err    error
	)
	if err = d.ParsePayload(&params); err != nil {
		return err
	}
	if ro = room.FindRoom(func(r room.Room) bool {
		return r.Id() == params.RoomId
	}); ro == nil {
		return Data{NetCmd: NetCmdGetRoomFailedResponse, Payload: "room is not exist"}.Send(c)
	}
	return Data{NetCmd: NetCmdGetRoomSucceedResponse, Payload: ro}.Send(c)
}
