package player

import "github.com/gofrs/uuid"

type Player struct {
	Id       uuid.UUID `json:"id"`
	RoomId   uuid.UUID `json:"room_id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Nickname string    `json:"nickname"`
}
