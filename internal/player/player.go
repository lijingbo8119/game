package player

import "github.com/gofrs/uuid"

type Player struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	RoomId uuid.UUID `json:"room_id"`
}
