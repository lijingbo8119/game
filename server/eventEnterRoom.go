package server

import (
	"game/internal/player"
	"game/internal/room"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type EventEnterRoomParams struct {
	RoomId   uuid.UUID
	PlayerId uuid.UUID
}

func eventEnterRoom(c *websocket.Conn, d Data) error {
	var (
		params = EventEnterRoomParams{}
		ro     room.Room
		p      *player.Player
		err    error
	)
	if err = d.ParsePayload(&params); err != nil {
		return err
	}
	ro = room.FindRoom(func(r room.Room) bool {
		return r.Id() == params.RoomId
	})
	p = player.FindPlayer(func(p *player.Player) bool {
		return p.Id == params.PlayerId
	})
	if ro == nil || p == nil {
		return Data{NetCmd: NetCmdEnterRoomFailedResponse, Payload: "data not exist"}.Send(c)
	}
	if err = ro.Enter(p); err != nil {
		return Data{NetCmd: NetCmdEnterRoomFailedResponse, Payload: err.Error()}.Send(c)
	}
	return Data{NetCmd: NetCmdEnterRoomSucceedResponse, Payload: ro}.Send(c)
}
