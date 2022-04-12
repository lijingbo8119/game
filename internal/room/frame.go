package room

import (
	"encoding/json"
	"time"
)

type frame interface {
	json.Marshaler
	Name() string
	Time() time.Time
	beforeUpdate(Room) error
	update(Room) error
}
